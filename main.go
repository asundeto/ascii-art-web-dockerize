package main

import (
	ascart "ascart/functions"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Page)
	mux.HandleFunc("/ascii-art", ascart.Control)
	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	log.Println("Server is running on port http://localhost:8081/")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func Page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ascart.Error(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ascart.Error(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ascart.Error(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
