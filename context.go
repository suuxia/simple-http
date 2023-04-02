package simple_http

import "net/http"

type Context struct {
	response http.ResponseWriter
	request  *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,
	}
}
