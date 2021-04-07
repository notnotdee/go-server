package about

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Body string
}

func About(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "About", Body: "body"}
	t, err := template.ParseFiles("pages/about/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	t.Execute(w, p)
}