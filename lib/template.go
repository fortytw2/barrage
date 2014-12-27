package lib

import (
	"github.com/GeertJohan/go.rice"
	"html/template"
)

func RenderTemplate(name string) (*template.Template, error) {
	// find a rice.Box
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		return nil, err
	}
	// load the base templatestring
	baseTemplateString, err := templateBox.String("base.tpl")
	if err != nil {
		return nil, err
	}
	// get file contents as string
	templateString, err := templateBox.String(name)
	if err != nil {
		return nil, err
	}

	// parse and execute the template
	t, err := template.New(name).Parse(baseTemplateString)
	if err != nil {
		return nil, err
	}
	t, err = t.Parse(templateString)
	if err != nil {
		return nil, err
	}

	return t, nil
}
