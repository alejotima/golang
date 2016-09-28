package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/alejotima/golang/test2/server/api"
	"github.com/gorilla/handlers"
	"os"
)

func main()  {
	r:=mux.NewRouter()
	r.HandleFunc("/api/features", api.GetAllFeatures).Methods("GET")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../client/dist/prod/"))))
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggedRouter)
}