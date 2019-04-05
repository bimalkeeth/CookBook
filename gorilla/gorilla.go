package main

import (
	"CookBook/gorilla/middleware"
	"CookBook/gorilla/newhandlers"
	"fmt"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/profile/{username}", newhandlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/triggerpanic", newhandlers.TriggerPanicHandler).Methods("GET")

	fmt.Println(gh.LoggingHandler(os.Stdout, r))
	http.Handle("/", middleware.PanicRecoveryHandler(gh.LoggingHandler(os.Stdout, r)))
	_ = http.ListenAndServe(WEBSERVERPORT, nil)

}
