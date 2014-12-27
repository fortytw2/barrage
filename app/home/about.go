package home

import (
	"github.com/GeertJohan/go.rice"
	"html/template"
	"net/http"
)

func GetAboutPage(w http.ResponseWriter, r *http.Request) {
	// find a rice.Box
	templateBox, err := rice.FindBox("../../templates")
	if err != nil {
		l.Println(err)
	}
	// load the base templatestring
	baseTemplateString, err := templateBox.String("base.tpl")
	if err != nil {
		l.Println(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("home/about.tpl")
	if err != nil {
		l.Println(err)
	}

	// parse and execute the template
	t, err := template.New("about").Parse(baseTemplateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t, err = t.Parse(templateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}
