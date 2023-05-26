package gee

import (
	"log"
	"net/http"
)

// router contains the handlerFunc
type router struct {
	handlers map[string]HandlerFunc
}

// newRouter is the method to construct a new router
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute add handler func to the handlers and prints the message
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle handle all the func, if not exists, return 404
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
