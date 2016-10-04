package frontend

import (
	"github.com/ericmieth/c3-registration/format"

	"html/template"
	"net/http"
)

func renderView(w http.ResponseWriter, data interface{}) {

	funcMap := template.FuncMap{
		"formatDate":  format.Date,
		"formatError": format.Error,
	}

	err := template.Must(
		template.New("index").
			Funcs(funcMap).
			ParseGlob("templates/frontend/*.tmpl")).Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
