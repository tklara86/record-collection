package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
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
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}


func ShowRecord(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1  {
		http.NotFound(w,r)
		return
	}

	_, err = fmt.Fprintf(w,"display record with id of %d", id)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {

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
