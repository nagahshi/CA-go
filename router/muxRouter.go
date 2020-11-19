package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Serve(port string) {
	fmt.Printf("MUX HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (*muxRouter) Get(uri string, fn func(res http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, fn).Methods("GET")
}

func (*muxRouter) Post(uri string, fn func(res http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, fn).Methods("POST")
}
