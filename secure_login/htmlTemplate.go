package main

import (
	"html/template"
	"net/http"
)

func loginTempl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, nil)
}

func registerTempl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("templates/sginup.html")
	t.Execute(w, nil)
}
