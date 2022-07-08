package utils

import (
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
)

// CopyFile is essentially CopyFileWithParams with overwrite set to true and perm set to 0700
func CopyFile(fs billy.Filesystem, src string, dst string) error {
	return CopyFileWithParams(fs, src, dst, true, 0700)
}
func CopyFileWithParams(bfs billy.Filesystem, src string, dst string, overwrite bool, perm os.FileMode) error {
	// verify that src exsists
	{
		fileInfo, err := bfs.Lstat(src)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return ErrResourceIsNotAFile
		}
	}
	// verify that dst not exsists (based on overwrite flag)
	{
		fileInfo, err := bfs.Lstat(dst)
		if err != nil {
			// file does not exist, so we are good to go
			if os.IsNotExist(err) {
				// as file does not exist, we can start the copy process.
				goto StartCopy
			}
			// if its any other error, return
			return err
		}
		// if we are here, means src exsists
		if !overwrite {
			// if overwrite is set to false, we just return with error
			return os.ErrExist
		}

		// make sure that src is not a dir
		if fileInfo.IsDir() {
			return ErrResourceIsNotAFile
		}
	}
StartCopy:
	bytes, err := util.ReadFile(bfs, src)
	if err != nil {
		return err
	}

	err = util.WriteFile(bfs, dst, bytes, perm)
	if err != nil {
		return err
	}

	return nil
}
