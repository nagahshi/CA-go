package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) Serve(port string) {
	fmt.Printf("CHI HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}

func (*chiRouter) Get(uri string, fn func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, fn)
}

func (*chiRouter) Post(uri string, fn func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, fn)
}
