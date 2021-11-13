package controller

import (
	"fmt"
	"github.com/record-collection/service"
	"log"
	"net/http"
	"strconv"
)

type recordController struct{}

var (
	recordService service.RecordService
)

type RecordController interface {
	ShowRecord(w http.ResponseWriter, r *http.Request)
	CreateRecord(w http.ResponseWriter, r *http.Request)
}

func NewRecordController(r service.RecordService) RecordController {
	recordService = r
	return &recordController{}
}

func (*recordController) ShowRecord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = fmt.Fprintf(w, "display record with id of %d", id)
	if err != nil {
		log.Fatal(err)
	}
}

func (*recordController) CreateRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)

		return
	}
	_, err := w.Write([]byte("Create record"))
	if err != nil {
		log.Fatal(err)
	}
}
