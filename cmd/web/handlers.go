package main

import (
	"fmt"
	"github.com/record-collection/errors"
	"github.com/record-collection/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	errorType = errors.NewServiceErrors()
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorType.NotFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/nav.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		errorType.ServerError(w, err)
		return
	}


	type Link struct {
		LinkTitle string
		LinkPath string
	}

	type Links struct {
		Link []Link
	}

	data := Links{
		Link: []Link{
			{
				LinkTitle: "Home",
				LinkPath:  "/",
			},
			{
				LinkTitle: "Add Record",
				LinkPath:  "/record/create",
			},
		},
	}

	err = ts.Execute(w, data)
	if err != nil {
		errorType.ServerError(w, err)
	}
}



func (app *application) ShowRecord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		errorType.NotFound(w)
		return
	}

	_, err = fmt.Fprintf(w, "display record with id of %d", id)
	if err != nil {
		log.Fatal(err)
	}
}


func (app *application) CreateRecord(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			return
		}

		record := &models.Record{
			Title:     	"New Album",
			Label: 		"DG",
			Year: 		"2016",

		}

		id, err := app.records.Insert(record)
		if err != nil {
			errorType.ServerError(w, err)
		}

		http.Redirect(w, r, fmt.Sprintf("/record?id=%d", id), http.StatusSeeOther)
}




