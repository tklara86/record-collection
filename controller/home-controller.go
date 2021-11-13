package controller

import (
	"html/template"
	"log"
	"net/http"
)

type HomeController interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type homeController struct {}

func NewHomeController() HomeController{
	return &homeController{}
}

func (*homeController) Home(w http.ResponseWriter, r *http.Request) {
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