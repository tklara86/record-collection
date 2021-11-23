package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/record-collection/config"
	"github.com/record-collection/controller"
	router "github.com/record-collection/http"
	"github.com/record-collection/repository"
	"github.com/record-collection/service"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	repo 			 = repository.NewMySQLRepository()
	recordService    = service.NewRecordService
	recordController = controller.NewRecordController(recordService)
	homeController   = controller.NewHomeController()
	httpRouter       = router.NewMuxRouter()
	logError  		*config.Application
)

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

	httpRouter.GET("/", homeController.Home)
	httpRouter.GET("/record", recordController.ShowRecord)
	httpRouter.POST("/record/create", recordController.CreateRecord)
	httpRouter.SERVE(os.Getenv("PORT"))
}


func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}