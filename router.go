package main

import (
	"github.com/gorilla/mux"
	"ocrClient/controller"
)

func InitRouter() (*mux.Router)  {

	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetOcrUi).Methods("GET")
	router.HandleFunc("/ocr/api", controller.GetOcrApi).Methods("GET")
//	router.HandleFunc("/", controller.MainUi).Methods("GET")
//	router.NotFoundHandler = http.HandlerFunc(controller.MainUi)

	return router
}
