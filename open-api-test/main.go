package main

import (
  "log"
	"net/http"
	"github.com/ohsu-comp-bio/funnel/tes/openapi"
)

func main() {
	log.Printf("Server started")

	TaskServiceApiService := openapi.NewTaskServiceApiService()
	TaskServiceApiController := openapi.NewTaskServiceApiController(TaskServiceApiService)

	router := openapi.NewRouter(TaskServiceApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
