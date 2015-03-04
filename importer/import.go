package importer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/fortytw2/barrage/models"
	"github.com/jmoiron/sqlx"
	"github.com/kr/fs"
)

func Import(rootdir string, db *sqlx.DB) {

	walker := fs.Walk(rootdir)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		rel, err := filepath.Rel(rootdir, walker.Path())
		if err != nil {
			log.Println("importer: ", err)
		}
		// skip over the . file
		if rel == "." {
			continue
		}
		// so we find a tv series...
		if filepath.Base(walker.Path()) == "series.toml" {

			var series models.Series
			if _, err := toml.DecodeFile(walker.Path(), &series); err != nil {
				fmt.Println(err)
			}

			series.PosterURL = "/video/" + series.Title + "/poster.png"

			if series.Seasons == 0 {
				series.Seasons = 1
			}

			if err := series.Save(db); err != nil {
				log.Println("importer: ", err)
			} else {
				log.Println("importer: inserted series ", series.Title)
			}

			dbseries, err := models.GetSeriesByTitle(series.Title, db)
			if err != nil {
				log.Println(err)
			}

			var episodes EpisodeList
			if _, err := toml.DecodeFile(walker.Path(), &episodes); err != nil {
				log.Println(err)
			}

			for i, e := range episodes.Episodes {
				e.SeriesID = dbseries.ID
				e.Number = int64(i) + 1
				if err = e.Save(db); err != nil {
					log.Println(err)
				} else {
					log.Println("importer: inserted episode ", e.Title)
				}
			}
		}
	}
}

type EpisodeList struct {
	Episodes []models.Episode
}
