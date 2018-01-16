#### block.go



##### Block

###### Delete

```go
// Delete matching series between mint and maxt in the block.
// 前面说到, Delete 的时候会暂时先标记为 Tombstone, 这里即实现部分
func (pb *Block) Delete(mint, maxt int64, ms ...labels.Matcher) error {
	// ...

	err = pb.tombstones.Iter(func(id uint64, ivs Intervals) error {
		for _, iv := range ivs {
			stones.add(id, iv)
			pb.meta.Stats.NumTombstones++
		}
		return nil
	})
	if err != nil {
		return err
	}
	pb.tombstones = stones

	if err := writeTombstoneFile(pb.dir, pb.tombstones); err != nil {
		return err
	}
	return writeMetaFile(pb.dir, &pb.meta)
}
```



###### CleanTombstones

```
// CleanTombstones will rewrite the block if there any tombstones to remove them
// and returns if there was a re-write.
func (pb *Block) CleanTombstones(dest string, c Compactor) (bool, error) {
	numStones := 0

	pb.tombstones.Iter(func(id uint64, ivs Intervals) error {
		for _ = range ivs {
			numStones++
		}

		return nil
	})

	if numStones == 0 {
		return false, nil
	}

	if _, err := c.Write(dest, pb, pb.meta.MinTime, pb.meta.MaxTime); err != nil {
		return false, err
	}

	return true, nil
}
```



###### Snapshot

疑问, 这里仅对目标文件夹及其内部文件做了 hardlink, 怎么确保内容不变?



