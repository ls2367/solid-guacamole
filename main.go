package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	http.HandleFunc("/", handlerindex)
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}

func handlerindex(w http.ResponseWriter, r *http.Request) {
	path := "project-reports.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	t.Execute(w, nil)
}
