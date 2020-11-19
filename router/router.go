package router

import (
	"net/http"
)

type Router interface {
	Serve(port string)
	Get(uri string, fn func(res http.ResponseWriter, req *http.Request))
	Post(uri string, fn func(res http.ResponseWriter, req *http.Request))
	// Put(uri string, fn func(res http.ResponseWriter, req *http.Request))
	// Patch(uri string, fn func(res http.ResponseWriter, req *http.Request))
	// Delete(uri string, fn func(res http.ResponseWriter, req *http.Request))
}
