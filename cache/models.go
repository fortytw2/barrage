package cache

import (
  "time"
  "github.com/kr/fs"
  "github.com/BurntSushi/toml"
  "path/filepath"
  "os"
  "fmt"
)

type Series struct {
  Title string
  Description string
  Group string
  Episodes []Episode
}

type Episode struct {
  Title string
  ReleaseDate time.Time `toml:"date"`
  Description string
  File string `toml:"filename"`
}

type Movie struct {
  Title string
  ReleaseDate time.Time
  Description string
  File string
  Group string
}

func ParseFiles(path string) {
  walker := fs.Walk(path)
  for walker.Step() {
    if err := walker.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        continue
    }
    rel, err := filepath.Rel(path, walker.Path())
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
    if rel == "." {
      continue
    }
    // so we find a tv series...
    if filepath.Base(walker.Path()) == "tv.toml" {

      var series Series
      if _, err := toml.DecodeFile(walker.Path(), &series); err != nil {
        fmt.Println(err)
      }

      fmt.Printf("%+v", series)
    }
    // or we find a movie
    if filepath.Base(walker.Path()) == "movie.toml" {

      var movie Movie
      if _, err := toml.DecodeFile(walker.Path(), &movie); err != nil {
        fmt.Println(err)
      }

      fmt.Printf("%+v", movie)
    }
  }
}