package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fortytw2/barrage/models"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

// be sure to properly pass the database connection to the function. Not 100%
// sure whether or not this is the right way to do things
func GetSeries(db *sqlx.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		js, err := json.Marshal(models.GetAllSeries(db))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}
}

type JSONSeries struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterURL   string `json:"poster"`
	Seasons     int    `json:"seasons"`
	Episodes    []models.Episode
}

func GetSeriesDetail(db *sqlx.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s, err := models.GetSeriesByID(id, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		episodes, err := s.GetEpisodes(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonseries := &JSONSeries{
			ID:          s.ID,
			Title:       s.Title,
			Description: s.Description,
			PosterURL:   s.PosterURL,
			Seasons:     s.Seasons,
			Episodes:    episodes,
		}

		js, err := json.Marshal(jsonseries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}
}
