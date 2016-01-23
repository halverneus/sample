package file

import (
	"io"
	"os"
	"path"

	"github.com/halverneus/sample/settings"
)

// Upload file into storage.
func Upload(name string, reader io.Reader) (err error) {
	fullpath := settings.Get.Storage.Folder + "/" + name
	dir := path.Dir(fullpath)

	// Make all necessary directories.
	if err = os.MkdirAll(dir, 0755); nil != err {
		return
	}

	// Open file for writing.
	var file *os.File
	if file, err = os.OpenFile(fullpath, os.O_RDWR|os.O_CREATE, 0755); nil != err {
		return
	}
	defer file.Close()

	// Write to file.
	_, err = io.Copy(file, reader)
	return
}
