package web

import (
	"encoding/json"
)

// Reply to the client with a provided message.
type Reply struct {
	ctx    *Context
	status int
	obj    interface{}
}

// Status code to be returned for the request.
func (reply *Reply) Status(status int) *Reply {
	reply.status = status
	return reply
}

// With the supplied message, reply to the client.
func (reply *Reply) With(obj interface{}) *Reply {
	reply.obj = obj
	return reply
}

// Do performs the actual reply.
func (reply *Reply) Do() error {
	raw, err := json.Marshal(reply.obj)
	if nil != err {
		return err
	}

	reply.ctx.W.WriteHeader(reply.status)
	_, err = reply.ctx.W.Write(raw)
	return err
}
