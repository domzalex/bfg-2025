package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("templates/services/*.html"))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	if path == "" {
		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/home.html",
		))

		err := tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			log.Println("template error:", err)
			http.Error(w, "Page not found", http.StatusNotFound)
		}
		return
	}

	if path == "contact" {
		tmpl := template.Must(template.ParseFiles(
			"templates/contact.html",
			"templates/nav.html",
			"templates/footer.html",
		))

		err := tmpl.ExecuteTemplate(w, "contact", nil)
		if err != nil {
			log.Println("template error:", err)
			http.Error(w, "Page not found", http.StatusNotFound)
		}
		return
	}

	if strings.HasPrefix(path, "services/") {
		service := r.URL.Path[len("/services/"):] // "foo", "bar", etc.

		tmpl := template.Must(template.ParseFiles(
			"templates/service.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/contact-component.html",
			"templates/services/"+service+".html",
		))

		err := tmpl.ExecuteTemplate(w, "service", nil)
		if err != nil {
			log.Println("template error:", err)
			http.Error(w, "Service not found", http.StatusNotFound)
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe(":3000", nil)
}
