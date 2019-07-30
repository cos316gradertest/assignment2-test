// This router implements only the most basic, naive path matching, with no capturing

package http_router

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpRouter struct {
	// this can include whatever fields you want
	lookup map[string]http.HandlerFunc
}

// Creates a new HttpRouter
func NewRouter() *HttpRouter {
	router := new(HttpRouter)
	router.lookup = make(map[string]http.HandlerFunc)
	return router
}

func makeKey(method string, pattern string) string {
	method = strings.ToLower(method)
	path := strings.Trim(pattern, "/")
	return fmt.Sprintf("%s/%s", method, path)
}

// Adds a new route to the HttpRouter
func (self *HttpRouter) AddRoute(method string, pattern string, handler http.HandlerFunc) {
	self.lookup[makeKey(method, pattern)] = handler
}

// Conforms to the `http.HandlerHttp` interface
func (self *HttpRouter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	handler, ok := self.lookup[makeKey(request.Method, request.URL.Path)]
	if !ok {
		http.NotFound(response, request)
		return
	}
	handler.ServeHTTP(response, request)
}
