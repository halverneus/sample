package file

import (
	"net/http"

	"github.com/halverneus/sample/common/log"
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
		log.For("/API/FILE(ROUTE)", ctx.User).Warning().Printf("Unimplemented method: %s", ctx.R.Method)
		ctx.Reply().Status(http.StatusNotImplemented).Do()
	}
}
