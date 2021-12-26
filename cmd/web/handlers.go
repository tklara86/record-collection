package main

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gorilla/mux"
	"github.com/record-collection/errors"
	"github.com/record-collection/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	errorType = errors.NewServiceErrors()
)

// Home it is home page
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

	app.render(w, r, "home.page.tmpl", &templateData{Records: s})

}

// AdminLogin login page
func (app *application) AdminLogin(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{})
}

func (app *application) PostAdminLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("it works")
}

// ShowRecord show single record
func (app *application) ShowRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	recordId, err := strconv.Atoi(id)
	if err != nil || recordId < 1 {
		errorType.NotFound(w)
		return
	}

	// get record
	s, err := app.records.Get(recordId)
	if err != nil {
		errorType.ServerError(w, err)
		return
	}

	app.render(w, r, "record.page.tmpl", &templateData{
		Record: s,
	})

}

// CreateRecord displays create record page
func (app *application) CreateRecord(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "create.page.tmpl", &templateData{Record: nil})
}

// CreateRecordForm creates a new record
func (app *application) CreateRecordForm(w http.ResponseWriter, r *http.Request) {
	const MaxUploadSize = 1024 * 1024 // 1MB
	var ctx = context.Background()

	err := r.ParseForm()
	if err != nil {
		errorType.ClientError(w, http.StatusBadRequest)
		return
	}

	errorsMap := make(map[string]string)

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		errorType.ClientError(w, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("inputCover")
	if err != nil {
		errorsMap["noImage"] = "Please select an image"
		return
	}

	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		errorsMap["fileFormat"] = "The provided file format is not allowed. Please upload a JPEG or PNG image"
		//http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Creates a new file in the covers directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cloudinary upload
	cloudinaryName := os.Getenv("CLOUDINARY_NAME")
	cloudinaryAPIKey := os.Getenv("CLOUDINARY_API")
	cloudinarySecret := os.Getenv("CLOUDINARY_SECRET")

	cld, _ := cloudinary.NewFromParams(cloudinaryName, cloudinaryAPIKey, cloudinarySecret)

	resp, err := cld.Upload.Upload(ctx, dst.Name(),
		uploader.UploadParams{
			PublicID:       fileHeader.Filename,
			AllowedFormats: []string{"jpg", "jpeg", "png"},
			Folder:         "covers",
			Tags:           []string{"album cover"},
		})

	record := &models.Record{
		Title: r.PostForm.Get("inputTitle"),
		Label: r.PostForm.Get("inputLabel"),
		Year:  r.PostForm.Get("inputYear"),
		Cover: resp.SecureURL,
	}

	if strings.TrimSpace(record.Title) == "" {
		errorsMap["title"] = "This field cannot be blank"
	}

	if utf8.RuneCountInString(record.Title) > 50 {
		errorsMap["title"] = "This field is too long(maximum 50 characters)"
	}

	if strings.TrimSpace(record.Label) == "" {
		errorsMap["label"] = "This field cannot be blank"
	}

	if strings.TrimSpace(record.Year) == "" {
		errorsMap["year"] = "This field cannot be blank"
	}

	if len(errorsMap) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormData:   r.PostForm,
			FormErrors: errorsMap,
		})
		return
	}

	id, err := app.records.Insert(record)
	if err != nil {
		errorType.ServerError(w, err)
		return
	}

	// Remove uploads file
	err = os.RemoveAll("./uploads")
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, fmt.Sprintf("/record/%d", id), http.StatusSeeOther)

}
