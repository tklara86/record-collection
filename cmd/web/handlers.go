package main

import (
	//"errors"
	"fmt"
	"github.com/record-collection/errors"
	"github.com/record-collection/models"
	"net/http"
	"strconv"
)

var (
	errorType = errors.NewServiceErrors()
	links = []link{
		{
			LinkTitle: "Home",
			LinkPath:  "/",
		},
		{
			LinkTitle: "Add record",
			LinkPath:  "/record/create",
		},
	}
)



func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorType.NotFound(w)
		return
	}

	// get all records
	s, err := app.records.GetAll()
	if err != nil {
		errorType.ServerError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Records: s, Links: links})

}



func (app *application) ShowRecord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		errorType.NotFound(w)
		return
	}

	// get record
	s, err := app.records.Get(id)
	if err != nil {
		errorType.ServerError(w, err)
		return
	}


	app.render(w, r, "record.page.tmpl", &templateData{
		Links: links,
		Record:  s,
	})

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




