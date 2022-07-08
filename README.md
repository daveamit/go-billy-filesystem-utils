# Utils for [billy.Filesystem](github.com/go-git/go-billy) based filesystem 
This package contains following most commanly used utility functions

---


## Installation

```go
import utils "github.com/daveamit/go-billy-filesystem-utils" 
```

---

## CopyFile
CopyFile is just a sugared wrapper over CopyFileWithParams with overwrite set to true and perms set to 0700

# Sample CopyFile usage
```go
func main()  {
	bfs := osfs.New("path/to/root/dir")

    err := utils.CopyFile(bfs, "src-file", "dst-file")
    if err != nil {
        panic(err)
    }
}
```

---

## CopyFileWithParams
CopyFileWithParams takes src, dst, overwrite flag which indicates if the dst is to be overwritten or not in case it already exists and perms. As the name suggests, the function calls src to dst, it creates dst with given perms

# Sample CopyFile usage
```go
func main()  {
	bfs := osfs.New("path/to/root/dir")

    err := utils.CopyFileWithParams(bfs, "src-file", "dst-file", true, 0700)
    if err != nil {
        panic(err)
    }
}
```
---
## more functions to come 
 * CopyDir
 * CopyDirWithParams
 * MoveFile
 * MoveFileWithParams
 * MoveDir
 * MoveDirWithParams
 * (suggest more ...)

---

## License
This package is distributed under the Apache License Version 2.0, see [LICENSE](LICENSE) for details.