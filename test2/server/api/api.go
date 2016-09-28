package api

import (
	"net/http"
	"fmt"
	"github.com/alejotima/golang/test2/server/db"
	"log"
	"encoding/json"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

// GetAllItems returns a list of all database items to the response.
func GetAllFeatures(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed get all features: ", err)
		return
	}

	respBody, err := json.MarshalIndent(rs, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

