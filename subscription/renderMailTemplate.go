package subscription

import (
	"bytes"
	"log"
	"text/template"
)

func renderMailTemplate(s Subscription, templateName string) []byte {

	buf := new(bytes.Buffer)

	err := template.Must(
		template.New(templateName).ParseGlob("templates/mail/*.tmpl")).
		ExecuteTemplate(buf, templateName, s)
	if err != nil {
		log.Print(err)
	}

	return []byte(buf.String())
}
