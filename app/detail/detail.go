package detail

import (
	"github.com/fortytw2/barrage/lib"
	"log"
	"os"
	"net/http"
)

var l *log.Logger

func init()  {
	l = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
}

func GetDetailPage(w http.ResponseWriter, r *http.Request) {
	t, err := lib.RenderTemplate("detail/detail.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}
