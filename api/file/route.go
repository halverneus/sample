package file

import (
	"net/http"

	"github.com/halverneus/sample/common/web"
)

// Route request based on method.
func Route(ctx *web.Context) {
	switch ctx.R.Method {
	case "GET":
		get(ctx)

	case "PUT":
		put(ctx)

	case "DELETE":
		del(ctx)

	default:
		ctx.Reply().Status(http.StatusNotImplemented).Do()
	}
}
