package subscription

import (
	"bytes"
	"log"
	"text/template"
)

func renderMailTemplate(s Subscription, templateName string) []byte {

	buf := new(bytes.Buffer)

	pathPrefix := "templates/mail/"

	err := template.Must(
		ParseFiles(pathPrefix+templateName+".tmpl")).
		Execute(buf, s)
	if err != nil {
		log.Print(err)
	}

	return []byte(buf.String())
}
