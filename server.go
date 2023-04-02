package simple_http

import "net/http"

type NextFunc func()
type HandlerFunc func(ctx *Context, next NextFunc)

type Server struct {
	middleware []HandlerFunc
}

func (s *Server) NewServer() *Server {
	return &Server{
		middleware: make([]HandlerFunc, 0, 10),
	}
}

func (s *Server) Use(handler HandlerFunc) {
	s.middleware = append(s.middleware, handler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(s.middleware) == 0 {
		http.NotFound(w, r)
	} else {
		fn := compose(s.middleware)
		ctx := NewContext(w, r)

		fn(ctx, nil)
	}
}

func (s *Server) Listen(port string) {
	http.ListenAndServe(port, s)
}

func compose(middleware []HandlerFunc) HandlerFunc {

	return func(ctx *Context, next NextFunc) {

	}
}
