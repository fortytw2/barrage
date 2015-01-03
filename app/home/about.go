package home

import (
	"github.com/fortytw2/barrage/lib"
	"net/http"
)

func GetAboutPage(w http.ResponseWriter, r *http.Request) {
	
	t, err := lib.RenderTemplate("home/about.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}
