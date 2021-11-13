package main

import (
	"github.com/record-collection/controller"
	router "github.com/record-collection/http"
	"github.com/record-collection/service"
	"os"
)

var (
	//recordRepositort = repository.RecordRepository()
	recordService    = service.NewRecordService
	recordController = controller.NewRecordController(recordService)
	homeController   = controller.NewHomeController()
	httpRouter       = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/", homeController.Home)
	httpRouter.GET("/record", recordController.ShowRecord)
	httpRouter.POST("/record/create", recordController.CreateRecord)
	httpRouter.SERVE(os.Getenv("PORT"))
}
