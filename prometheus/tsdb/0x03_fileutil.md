#### fileutil

```
fileutil
├── dir_unix.go
├── dir_windows.go
├── fileutil.go
├── mmap.go
├── mmap_unix.go
├── mmap_windows.go
├── preallocate.go
├── preallocate_darwin.go
├── preallocate_linux.go
├── preallocate_other.go
├── sync.go
├── sync_darwin.go
└── sync_linux.go
```



[fileutil](https://godoc.org/github.com/prometheus/tsdb/fileutil) 提供了一些操作文件/目录的函数, 处理了不同平台 (主要是 win) 的兼容性问题.

除非遇到类似的场景需要相关的处理思路.

直接阅读 godoc 即可.

