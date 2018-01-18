#### chunks

```
chunks
└── chunks.go
```



[chunks](https://godoc.org/github.com/prometheus/tsdb/chunks#Meta) 是 chunk 数据持久化的实现



##### Meta

Chunk 的元数据



##### Writer

是一个基于文件目录的 [ChunkWriter](https://github.com/prometheus/tsdb/blob/d45b595a1daefad23c09a2d994bf956f8b5f15a9/block.go#L99-L110) 实现.

执行一组 Chunk 的写入, 并按体积进行分片, 每一片称为一个 `sequenceFile`.

1. 结构体定义

   ```go
   // Writer implements the ChunkWriter interface for the standard
   // serialization format.
   type Writer struct {
   	// 当前目录
   	dirFile *os.File

   	// 所有用于写入的数据文件, 只有最后一个是当前有效的
   	files   []*os.File
   	wbuf    *bufio.Writer
   	
   	// 当前分片文件已写入的字节数
   	n       int64

   	// 复用的 crc32, 用于每一个写入的 Chunk 的校验
   	crc32   hash.Hash

   	// 分片的尺寸, 目前是 512 << 20
   	segmentSize int64
   }
   ```

2. Writer.finalizeTail & Writer.cut

   ```go
   // 安全地关闭当前用于写入的文件
   func (w *Writer) finalizeTail() error {
   	// ...
   	
   	// As the file was pre-allocated, we truncate any superfluous zero bytes.
   	// 由于每个 seq file 都会预先分配空间, 因此需要按照实际使用量进行一次 Truncate
   	off, err := tf.Seek(0, os.SEEK_CUR)
   	if err != nil {
   		return err
   	}
   	if err := tf.Truncate(off); err != nil {
   		return err
   	}

   	return tf.Close()
   }
   ```

   ``` go
   func (w *Writer) cut() error {
   	// Sync current tail to disk and close.
   	if err := w.finalizeTail(); err != nil {
   		return err
   	}

     	// 打开一个新文件用于数据写入
   	p, _, err := nextSequenceFile(w.dirFile.Name())
   	if err != nil {
   		return err
   	}
   	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0666)
   	if err != nil {
   		return err
   	}
     
     	// 为文件预分配 segmentSize 大小
   	if err = fileutil.Preallocate(f, w.segmentSize, true); err != nil {
   		return err
   	}
   	if err = w.dirFile.Sync(); err != nil {
   		return err
   	}

   	// Write header metadata for new file.

   	metab := make([]byte, 8)
   	binary.BigEndian.PutUint32(metab[:4], MagicChunks)
   	metab[4] = chunksFormatV1

   	if _, err := f.Write(metab); err != nil {
   		return err
   	}

     	// 重置或初始化一个 bufio.Writer
   	w.files = append(w.files, f)
   	if w.wbuf != nil {
   		w.wbuf.Reset(f)
   	} else {
   		w.wbuf = bufio.NewWriterSize(f, 8*1024*1024)
   	}
     
     	// 重置已写入的字节数, 8 是 文件头 MagicChunks 占用的大小
     	// 即前面 `Write header metadata for new file.` 的 b 部分
   	w.n = 8

   	return nil
   }
   ```

3. Writer.WriteChunks

   ```go
   func (w *Writer) WriteChunks(chks ...Meta) error {
   	// Calculate maximum space we need and cut a new segment in case
   	// we don't fit into the current one.
     	// 计算所有 chunks 可能占用的空间
   	maxLen := int64(binary.MaxVarintLen32) // The number of chunks.
   	for _, c := range chks {
   		maxLen += binary.MaxVarintLen32 + 1 // The number of bytes in the chunk and its encoding.
   		maxLen += int64(len(c.Chunk.Bytes()))
   	}
   	newsz := w.n + maxLen

     // 根据这里执行 w.cut() 的判断条件, 实际上是可能出现单个文件超过 w.segmentSize 的
   	if w.wbuf == nil || w.n > w.segmentSize || newsz > w.segmentSize && maxLen <= w.segmentSize {
   		if err := w.cut(); err != nil {
   			return err
   		}
   	}

   	var (
   		// 初始化 b 作为写入长度, 写入 chunk 编码方式, 计算 hash 时等的 buffer
   		b   = [binary.MaxVarintLen32]byte{}
   		seq = uint64(w.seq()) << 32
   	)
   	for i := range chks {
   		chk := &chks[i]

   		// 用于定位 Chunk 
   		chk.Ref = seq | uint64(w.n)

   		// ...

   		// 校验数据
   		w.crc32.Reset()
   		if err := chk.writeHash(w.crc32); err != nil {
   			return err
   		}
   		if err := w.write(w.crc32.Sum(b[:0])); err != nil {
   			return err
   		}
   	}

   	return nil
   }
   ```



##### ByteSlice

用于 Reader 中逐段读取数据



##### Reader

用于读取数据块

1. NewDirReader

   ```go
   // NewDirReader returns a new Reader against sequentially numbered files in the
   // given directory.
   func NewDirReader(dir string, pool chunkenc.Pool) (*Reader, error) {
    	// 根据命名规则得到所有数据文件
   	files, err := sequenceFiles(dir)
   	if err != nil {
   		return nil, err
   	}
     
     	// 初始化一个 chunkenc.Pool
   	if pool == nil {
   		pool = chunkenc.NewPool()
   	}

   	var bs []ByteSlice
   	var cs []io.Closer

   	for _, fn := range files {
   		f, err := fileutil.OpenMmapFile(fn)
   		if err != nil {
   			return nil, errors.Wrapf(err, "mmap files")
   		}
   		cs = append(cs, f)
   		bs = append(bs, realByteSlice(f.Bytes()))
   	}
   	return newReader(bs, cs, pool)
   }
   ```

2. Reader.Chunk

   ```
   // 根据定位读取指定 Chunk
   func (s *Reader) Chunk(ref uint64) (chunkenc.Chunk, error) {
   	// 分别计算文件所在的位置, 和 Chunk 数据的起始位置
   	var (
   		seq = int(ref >> 32)
   		off = int((ref << 32) >> 32)
   	)
   	
   	// ...

   	// 将数据封装成一个 chunkenc.Chunk
   	return s.pool.Get(chunkenc.Encoding(r[0]), r[1:1+l])
   }
   ```

   ​