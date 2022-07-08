package utils_test

import (
	"os"
	"path"
	"testing"

	utils "github.com/daveamit/go-billy-filesystem-utils"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-billy/v5/util"
)

func TestCopyWithParamsNoOverwrite(t *testing.T) {
	bfs := osfs.New(path.Join("test-data", "copy", "file"))

	err := utils.CopyFileWithParams(bfs, "a.txt", "a-copy.txt", false, 0700)
	if err != nil && err != os.ErrExist {
		t.Error(err)
		t.FailNow()
	}

	bytes, err := util.ReadFile(bfs, "a-copy.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if string(bytes) != "a-copy" {
		t.Error("file content changed even with overwrite set to no")
		t.FailNow()
	}
}

func TestCopyWithParamsOverwrite(t *testing.T) {
	bfs := osfs.New(path.Join("test-data", "copy", "file"))

	err := utils.CopyFileWithParams(bfs, "a.txt", "a-copy.txt", true, 0700)
	if err != nil && err != os.ErrExist {
		t.Error(err)
		t.FailNow()
	}

	bytes, err := util.ReadFile(bfs, "a-copy.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if string(bytes) != "a" {
		t.Error("file content expected to be a but is set to: " + string(bytes))
		t.FailNow()
	}

	// cleanup
	{
		err = util.WriteFile(bfs, "a-copy.txt", []byte("a-copy"), 0700)
		if err != nil {
			t.Error("Failed to cleanup: " + err.Error())
			t.FailNow()
		}
	}
}

func TestCopyFileWithParams(t *testing.T) {
	bfs := osfs.New(path.Join("test-data", "copy", "file"))

	err := utils.CopyFileWithParams(bfs, "a.txt", "a-copy-1.txt", false, 0700)
	if err != nil && err != os.ErrExist {
		t.Error(err)
		t.FailNow()
	}

	bytes, err := util.ReadFile(bfs, "a-copy-1.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if string(bytes) != "a" {
		t.Error("file content expected to be a but is set to: " + string(bytes))
		t.FailNow()
	}

	// cleanup
	{
		err = util.RemoveAll(bfs, "a-copy-1.txt")
		if err != nil {
			t.Error("Failed to cleanup: " + err.Error())
			t.FailNow()
		}
	}
}

func TestCopyFile(t *testing.T) {
	bfs := osfs.New(path.Join("test-data", "copy", "file"))

	err := utils.CopyFile(bfs, "a.txt", "a-copy-1.txt")
	if err != nil && err != os.ErrExist {
		t.Error(err)
		t.FailNow()
	}

	bytes, err := util.ReadFile(bfs, "a-copy-1.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if string(bytes) != "a" {
		t.Error("file content expected to be a but is set to: " + string(bytes))
		t.FailNow()
	}

	// cleanup
	{
		err = util.RemoveAll(bfs, "a-copy-1.txt")
		if err != nil {
			t.Error("Failed to cleanup: " + err.Error())
			t.FailNow()
		}
	}
}

func TestCopyFileIsADir(t *testing.T) {
	bfs := osfs.New(path.Join("test-data", "copy", "file"))

	err := utils.CopyFile(bfs, "d", "d.txt")
	if err != nil && err != utils.ErrResourceIsNotAFile {
		t.Error(err)
		t.FailNow()
	}

	err = utils.CopyFile(bfs, "a.txt", "d")
	if err != nil && err != utils.ErrResourceIsNotAFile {
		t.Error(err)
		t.FailNow()
	}
}

func TestCopyFileSrcDoesNotExist(t *testing.T) {
	bfs := osfs.New(path.Join("test-data", "copy", "file"))

	err := utils.CopyFile(bfs, "b.txt", "b-copy.txt")
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
		t.FailNow()
	}

	err = utils.CopyFile(bfs, "a.txt", "d")
	if err != nil && err != utils.ErrResourceIsNotAFile {
		t.Error(err)
		t.FailNow()
	}
}
