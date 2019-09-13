package handlers

import (
	"net/http"
	"text/template"
)

type page struct {
	Title string
	Msg   string
}

// home is a simple HTTP handler function which writes a response.
func home(w http.ResponseWriter, _ *http.Request) {
	// fmt.Fprint(w, "Hello! Your request was processed.")
	w.Header().Set("Content-type", "text/html")

	t, _ := template.ParseFiles("template/home.html")
	t.Execute(w, &page{Title: "Тестовое задание", Msg: "Тестовое задание"})
}
