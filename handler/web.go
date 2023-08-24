package handler

import (
	"net/http"
	"text/template"
)

type Product struct {
	Endpoint string
	URL      string
}

// serves the homepage
func (h *HandlerHelper) Homepage(w http.ResponseWriter, r *http.Request) {
	// returns all the endpoints as []data.Endpoint
	shortcuts, w := getAndMarshalEndpoints(h.l, h.db, w)

	// parse html template
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "Unable to generate HTML", http.StatusInternalServerError)
		return
	}

	// construct/inject html file using template and shortcuts object
	err = tmpl.Execute(w, shortcuts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
