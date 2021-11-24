package controller

import (
	"github.com/record-collection/models/mysql"
	"net/http"
)

type recordController struct{
	app *mysql.RecordModel
}

var (
	//recordService service.RecordService
	//logError  		errors.Errors

)

type RecordController interface {
	ShowRecord(w http.ResponseWriter, r *http.Request)
	CreateRecord(controller *recordController) http.HandlerFunc

}
//
//func NewRecordController() RecordController {
//	//recordService = r
//	return &recordController{}
//}
//
//func (rc *recordController) ShowRecord(w http.ResponseWriter, r *http.Request) {
//	id, err := strconv.Atoi(r.URL.Query().Get("id"))
//	if err != nil || id < 1 {
//		errorType.NotFound(w)
//		return
//	}
//
//	_, err = fmt.Fprintf(w, "display record with id of %d", id)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//




//func (rc *recordController) CreateRecord(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method != http.MethodPost {
//		w.Header().Set("Allow", http.MethodPost)
//		return
//	}
//	//app := &Application{Records: &mysql.RecordModel{DB: db}}
//
//	record := &models.Record{
//		Title:     	"New Album",
//		Label: 		"DG",
//		Year: 		"2016",
//
//	}
//
//	id, err := rc.app.Insert(record)
//	if err != nil {
//		//logError.ServerError(w, err)
//	}
//
//	http.Redirect(w, r, fmt.Sprintf("/record?id=%d", id), http.StatusSeeOther)
//}

