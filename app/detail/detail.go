package detail

import (
	"github.com/GeertJohan/go.rice"
	"log"
	"os"
	"html/template"
	"net/http"
)

var l *log.Logger

func init()  {
	l = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
}

func GetDetailPage(w http.ResponseWriter, r *http.Request) {
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
	templateString, err := templateBox.String("detail/detail.tpl")
	if err != nil {
		l.Println(err)
	}

	// parse and execute the template
	t, err := template.New("detail").Parse(baseTemplateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t, err = t.Parse(templateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}
