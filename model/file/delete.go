package file

import (
	"io"
	"os"
	"path"

	"github.com/halverneus/sample/settings"
)

// Delete file from storage.
func Delete(name string) (err error) {
	fullpath := settings.Get.Storage.Folder + "/" + name

	// Remove file.
	if err = os.Remove(fullpath); nil != err {
		return
	}

	// Remove empty parent directories.
	dir := path.Dir(fullpath)
	for dirIsEmpty(dir) {
		if err = os.Remove(dir); nil != err {
			return nil // We don't really care about this error.
		}
		dir = path.Clean(dir + "/..")
	}

	return
}

// dirIsEmpty returns true when a provided directory is empty.
func dirIsEmpty(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdir(1)
	return io.EOF == err
}
