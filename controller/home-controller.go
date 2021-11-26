package controller

//import (
//	"github.com/record-collection/errors"
//	"html/template"
//	"net/http"
//)
//
//var (
//	errorType = errors.NewServiceErrors()
//)
//
//type HomeController interface {
//	Home(w http.ResponseWriter, r *http.Request)
//}
//
//type homeController struct {}
//
//func NewHomeController() HomeController{
//	return &homeController{}
//}
//
//func (*homeController) Home(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		errorType.NotFound(w)
//		return
//	}
//
//	files := []string{
//		"./ui/html/home.page.tmpl",
//		"./ui/html/base.layout.tmpl",
//		"./ui/html/footer.partial.tmpl",
//		"./ui/html/nav.partial.tmpl",
//	}
//
//	ts, err := template.ParseFiles(files...)
//
//	if err != nil {
//		errorType.ServerError(w, err)
//		return
//	}
//
//
//	type Link struct {
//		LinkTitle string
//		LinkPath string
//	}
//
//	type Links struct {
//		Link []Link
//	}
//
//	data := Links{
//		Link: []Link{
//			{
//				LinkTitle: "Home",
//				LinkPath:  "/",
//			},
//			{
//				LinkTitle: "Add Record",
//				LinkPath:  "/record/create",
//			},
//		},
//	}
//
//	err = ts.Execute(w, data)
//	if err != nil {
//		errorType.ServerError(w, err)
//	}
//}