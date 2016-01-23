package file

import (
	"io"
	"os"

	"github.com/halverneus/sample/settings"
)

// Download file from storage.
func Download(name string, writer io.Writer) (err error) {
	fullpath := settings.Get.Storage.Folder + "/" + name

	// Open file for reading.
	var file *os.File
	if file, err = os.OpenFile(fullpath, os.O_RDONLY, 0755); nil != err {
		return
	}
	defer file.Close()

	// Send contents to client.
	_, err = io.Copy(writer, file)
	return
}
