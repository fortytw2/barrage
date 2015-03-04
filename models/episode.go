package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Episode struct {
	ID          int64     `json:"-"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"date"`
	Season      int64     `json:"season"`
	Number      int64     `json:"id"`
	Files       []string  `json:"-"`
	SeriesID    int64     `json:"-"`
}

func (e *Episode) Save(db *sqlx.DB) error {
	_, err := db.NamedExec("INSERT INTO episodes (title, releasedate, season, number, seriesid) VALUES (:title, :releasedate, :season, :number, :seriesid)", e)
	if err != nil {
		return err
	}

	return nil
}

func (s *Series) GetEpisodes(db *sqlx.DB) ([]Episode, error) {
	var episodes []Episode

	rows, err := db.Queryx("SELECT * FROM episodes WHERE seriesid=$1", s.ID)
	for rows.Next() {
		var e Episode
		err = rows.StructScan(&e)
		if err != nil {
			return nil, err
		}

		episodes = append(episodes, e)
	}

	return episodes, nil
}
