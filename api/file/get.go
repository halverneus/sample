package file

import (
	"net/http"
	"strings"

	"github.com/halverneus/sample/common/web"
	"github.com/halverneus/sample/model/file"
)

// get file from storage.
func get(ctx *web.Context) {
	path := strings.Replace(ctx.R.URL.Path, "/api/file/", "", 1)

	if 0 == len(path) {
		msg := "Format should be /api/file/path/to/my.file"
		ctx.Reply().Status(http.StatusBadRequest).With(msg).Do()
		return
	}

	if err := file.Download(path, ctx.W); nil != err {
		ctx.Reply().Status(http.StatusNotFound).With(err.Error()).Do()
		return
	}

	// Explicit response not required.
}
