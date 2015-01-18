package cache

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fortytw2/barrage/config"
	"github.com/kr/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Series struct {
	Title       string
	Description string
	RootURI     string
	Episodes    []Episode
}

type Episode struct {
	Title       string
	ReleaseDate time.Time `toml:"date"`
	Description string
	Season      int    `toml:"season"`
	File        string `toml:"filename"`
}

type Movie struct {
	Title       string
	ReleaseDate time.Time
	Description string
	File        string
	RootURI     string
}

var SeriesDB []Series
var MovieDB []Movie

// this package needs to be used with care - TOML is fast,
// but try to keep this from running hundreds of times needlessly

// TODO: clean up seasons (add 1 if it's 0), but that's not important
func init() {
	loadFromToml()
}

func loadFromToml() {
	walker := fs.Walk(config.SourceFolder)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		rel, err := filepath.Rel(config.SourceFolder, walker.Path())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if rel == "." {
			continue
		}
		// so we find a tv series...
		if filepath.Base(walker.Path()) == "series.toml" {

			var series Series
			if _, err := toml.DecodeFile(walker.Path(), &series); err != nil {
				fmt.Println(err)
			}

			baseURI := strings.TrimLeft(walker.Path(), config.StorageFolder)
			// account for running on non Unix-like platforms
			baseURI = filepath.ToSlash(baseURI)
			series.RootURI = strings.TrimRight(baseURI, "series.toml")

			SeriesDB = append(SeriesDB, series)
		}
		// or we find a movie
		if filepath.Base(walker.Path()) == "movie.toml" {

			var movie Movie
			if _, err := toml.DecodeFile(walker.Path(), &movie); err != nil {
				fmt.Println(err)
			}

			baseURI := strings.TrimLeft(walker.Path(), config.StorageFolder)
			// account for running on non Unix-like platforms
			baseURI = filepath.ToSlash(baseURI)
			movie.RootURI = strings.TrimRight(baseURI, "movie.toml")

			MovieDB = append(MovieDB, movie)
		}
	}
}
