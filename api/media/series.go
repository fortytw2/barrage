package media

import (
	"encoding/json"
	"github.com/fortytw2/barrage/cache"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func GetSeriesData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	js, err := json.Marshal(cache.SeriesDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// this needs to be fixed
func GetSingleSeries(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id < int64(len(cache.SeriesDB)) {
		js, err := json.Marshal(cache.SeriesDB[id])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
