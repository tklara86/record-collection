package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/record-collection/config"
	router "github.com/record-collection/http"
	"github.com/record-collection/models/mysql"
	"github.com/record-collection/repository"
	"html/template"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	repo 			 = repository.NewMySQLRepository()
	httpRouter       = router.NewMuxRouter()
	logError  		 *config.Application
)

type application struct {
	records *mysql.RecordModel
	templateCache map[string]*template.Template
}


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	db, err := repo.OpenDB(dsn)

	if err != nil {
		logError.ErrorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		logError.ErrorLog.Fatal(err)
	}

	app := &application{
		records: &mysql.RecordModel{DB: db},
		templateCache: templateCache,
	}


	httpRouter.GET("/", app.Home)
	httpRouter.GET("/record", app.ShowRecord)
	httpRouter.POST("/record/create", app.CreateRecord)
	httpRouter.SERVE(os.Getenv("PORT"))


}

