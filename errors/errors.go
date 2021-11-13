package errors

import "net/http"

type Errors interface {
	ServerError(w http.ResponseWriter, err error)
	ClientError(w http.ResponseWriter, status int)
	NotFound(w http.ResponseWriter)
}