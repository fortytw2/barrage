package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Series struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterURL   string `json:"poster"`
	Seasons     int    `json:"seasons"`
}

func GetAllSeries(db *sqlx.DB) []Series {
	var seriesDB []Series

	rows, err := db.Queryx("SELECT * FROM series")
	for rows.Next() {
		var s Series
		err = rows.StructScan(&s)
		if err != nil {
			log.Println(err)
		}

		seriesDB = append(seriesDB, s)
	}

	return seriesDB
}

func GetSeriesByTitle(title string, db *sqlx.DB) (*Series, error) {
	var series Series
	err := db.Get(&series, "SELECT * FROM series WHERE title=$1", title)
	if err != nil {
		return nil, err
	}

	return &series, nil
}

func GetSeriesByID(id int64, db *sqlx.DB) (*Series, error) {
	var series Series
	err := db.Get(&series, "SELECT * FROM series WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &series, nil
}

func (s *Series) Save(db *sqlx.DB) error {
	_, err := db.NamedExec("INSERT INTO series (title, description, posterurl, seasons) VALUES (:title, :description, :posterurl, :seasons)", s)
	if err != nil {
		return err
	}

	return nil
}

//
// func (e *Episode) GetSources() []Source {
// 	return nil
// }
