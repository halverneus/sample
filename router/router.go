package router

import (
	"net/http"

	"github.com/halverneus/sample/api/file"
	"github.com/halverneus/sample/common/log"
	"github.com/halverneus/sample/common/web"
	"github.com/halverneus/sample/model/user"
)

// Serve the API and bind to addr.
func Serve(addr string) error {

	http.HandleFunc("/api/file/", auth(file.Route))

	return http.ListenAndServe(addr, nil)
}

// auth (orize) the user to use the API service.
func auth(
	handler func(*web.Context),
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := web.NewContext(w, r)

		// Retrieve username and password from headers.
		username, password, ok := r.BasicAuth()
		if !ok {
			if 0 == len(username) {
				username = "UNKNOWN"
			}
			log.For("/ROUTER(AUTH)", username).Warning().Print("Unknown user and/or password")
			ctx.Reply().Status(http.StatusUnauthorized).Do()
			return
		}

		// Perform authorization check.
		if id, err := user.Authorize(username, password); nil != err {
			// FAILED
			log.For("/ROUTER(AUTH)", username).Warning().Print("Authentication failed")
			ctx.Reply().Status(http.StatusUnauthorized).With(err.Error()).Do()
		} else {
			// SUCCESS
			ctx.User = username
			ctx.UserID = id
			handler(ctx)
		}
	}
}
