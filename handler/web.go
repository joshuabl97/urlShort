package handler

import (
	"net/http"
	"text/template"
)

type Product struct {
	Endpoint string
	URL      string
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"/test", "https://www.google.com"},
		{"/example", "https://www.example.com"},
		// Add more products as needed
	}

	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "sometin", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
