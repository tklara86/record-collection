package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type muxRouter struct {}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (m *muxRouter) SERVE(port string) {

	fileServer := http.FileServer(http.Dir("dist"))
	muxDispatcher.PathPrefix("/").Handler(http.StripPrefix("/dist", neuter(fileServer)))

	fmt.Printf("Listening on port: %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), muxDispatcher)
	if err != nil {
		log.Fatal(err)
	}
}

func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}