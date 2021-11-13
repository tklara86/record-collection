package errors

import (
	"fmt"
	"github.com/record-collection/config"
	"log"
	"net/http"
	"runtime/debug"
)

type serviceError struct {}

var (
	app = &config.Application{}
)

func NewServiceErrors() Errors {
	return &serviceError{}
}

func (s *serviceError) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	err = app.ErrorLog.Output(2, trace)
	if err != nil {
		log.Fatal(err)
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (s *serviceError) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (s *serviceError) NotFound(w http.ResponseWriter) {
	s.ClientError(w, http.StatusNotFound)
}
