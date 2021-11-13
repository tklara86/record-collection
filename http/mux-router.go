package router

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/record-collection/config"
	"log"
	"net/http"
	"os"
	"strings"
)

type muxRouter struct {}

var (
	muxDispatcher = mux.NewRouter()
	app = &config.Application{}

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
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.LstdFlags)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.LstdFlags)

	addr := flag.String("addr", fmt.Sprintf(":%v", port), "HTTP network address")

	fileServer := http.FileServer(http.Dir("dist"))
	muxDispatcher.PathPrefix("/").Handler(http.StripPrefix("/dist", neuter(fileServer)))

	srv := &http.Server{
		Addr:           *addr,
		ErrorLog:		app.ErrorLog,
		Handler:        muxDispatcher,
	}

	app.InfoLog.Printf("Staring server on port: %s", *addr)
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
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