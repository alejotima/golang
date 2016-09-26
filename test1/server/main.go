package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
)

func HandlerTest(w http.ResponseWriter, r*http.Request)  {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r:=mux.NewRouter()
  /*
  * Estas lineas imprimen los archivos de una carpeta.
    files, _ := ioutil.ReadDir("../client/dist/prod/")
    for _, f := range files {
      fmt.Println(f.Name())
    }
  */

  r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../client/dist/prod/"))))
	r.HandleFunc("/test",HandlerTest)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000",loggedRouter)
}
