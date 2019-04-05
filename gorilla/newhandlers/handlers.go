package newhandlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	_, err := w.Write([]byte(username))
	if err != nil {
		log.Fatal("error in writing to response")
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome"))
	if err != nil {
		log.Fatal("error in writing to response")
	}
}

func TriggerPanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("triggering panic")
}
