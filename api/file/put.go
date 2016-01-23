package file

import (
	"net/http"
	"strings"

	"github.com/halverneus/sample/common/log"
	"github.com/halverneus/sample/common/web"
	"github.com/halverneus/sample/model/file"
)

// put file into storage.
func put(ctx *web.Context) {
	lg := log.For("/API/FILE(PUT)", ctx.User)
	path := strings.Replace(ctx.R.URL.Path, "/api/file/", "", 1)
	lg.Debug().Print("Attempting to put ", path)

	if 0 == len(path) {
		msg := "Format should be /api/file/path/to/my.file"
		lg.Warning().Print(msg)
		ctx.Reply().Status(http.StatusBadRequest).With(msg).Do()
		return
	}

	if err := file.Upload(path, ctx.R.Body); nil != err {
		lg.Warning().Print(err)
		ctx.Reply().Status(http.StatusInternalServerError).With(err.Error()).Do()
		return
	}

	ctx.Reply().Status(http.StatusOK).Do()
}
