package media

import (
	"encoding/json"
	"github.com/fortytw2/barrage/cache"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetMovieData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	js, err := json.Marshal(cache.MovieDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}