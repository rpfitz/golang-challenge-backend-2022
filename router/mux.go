package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var muxInstance = mux.NewRouter()

type Router interface {
	GET(url string, f func(http.ResponseWriter, *http.Request))
	POST(url string, f func(http.ResponseWriter, *http.Request))
	SERVE(port string)
}

type muxRouter struct {
}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(url string, f func(http.ResponseWriter, *http.Request)) {
	muxInstance.HandleFunc(url, f).Methods("GET", http.MethodOptions)
}

func (*muxRouter) POST(url string, f func(http.ResponseWriter, *http.Request)) {
	muxInstance.HandleFunc(url, f).Methods("POST", http.MethodOptions)
}

func (*muxRouter) SERVE(port string) {
	log.Println("Server is running on PORT:", port)
	http.ListenAndServe(":"+port, muxInstance)
}
