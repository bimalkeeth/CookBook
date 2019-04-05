package main

import (
	"html/template"
	"log"
	"net/http"
)

type Gopher struct {
	Name string
}

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Hopher"
	}
	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./templates/sample.html", gopher)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing template")
	}
	err = t.Execute(w, templateData)
	if err != nil {
		log.Fatal("Error in template data")
	}
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	_ = http.ListenAndServe(":8080", nil)
}
