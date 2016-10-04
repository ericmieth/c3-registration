package backend

import (
	"github.com/ericmieth/c3-registration/format"

	"net/http"
	"text/template"
)

func renderView(w http.ResponseWriter, name string, data interface{}) {

	funcMap := template.FuncMap{
		"formatDate": format.Date,
	}

	err := template.Must(
		template.New(name).
			Funcs(funcMap).
			ParseGlob("templates/backend/*.tmpl")).ExecuteTemplate(w, name, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
