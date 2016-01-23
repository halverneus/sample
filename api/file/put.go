package file

import (
	"net/http"
	"strings"

	"github.com/halverneus/sample/common/web"
	"github.com/halverneus/sample/model/file"
)

// put file into storage.
func put(ctx *web.Context) {
	path := strings.Replace(ctx.R.URL.Path, "/api/file/", "", 1)

	if 0 == len(path) {
		msg := "Format should be /api/file/path/to/my.file"
		ctx.Reply().Status(http.StatusBadRequest).With(msg).Do()
		return
	}

	if err := file.Upload(path, ctx.R.Body); nil != err {
		ctx.Reply().Status(http.StatusInternalServerError).With(err.Error()).Do()
		return
	}

	ctx.Reply().Status(http.StatusOK).Do()
}
