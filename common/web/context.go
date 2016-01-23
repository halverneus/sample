package web

import (
	"encoding/json"
	"net/http"
)

// Context is used to encapsulate an HTTP session and to simplify interactions
// with the client.
type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	User   string
	UserID string
}

// NewContext creates a new HTTP connection context.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

// FromJSON encoded request, populate the passed interface 'v'.
func (ctx *Context) FromJSON(v interface{}) error {
	decoder := json.NewDecoder(ctx.R.Body)
	return decoder.Decode(v)
}

// Reply the the client. This doubles as the sole constructor for a Reply.
func (ctx *Context) Reply() *Reply {
	return &Reply{
		ctx:    ctx,
		status: http.StatusOK,
	}
}
