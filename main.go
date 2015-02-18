package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/fortytw2/barrage/api/media"
	"github.com/fortytw2/barrage/config"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

//go:generate bash -c "lessc assets/less/barrage.less | cleancss > static/css/barrage.min.css"
//go:generate bash -c "browserify -t mithrilify assets/js/router.js | uglifyjs -o static/js/barrage.min.js"

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
}

func main() {
	router := httprouter.New()

	// series are different from individual movies, but not by much.
	router.GET("/api/series", media.GetSeriesData)
	router.GET("/api/movies", media.GetMovieData)
	router.GET("/api/series/:id", media.GetSingleSeries)

	router.ServeFiles("/video/*filepath", http.Dir(config.VideoFolder))

	router.NotFound = http.FileServer(rice.MustFindBox("static").HTTPBox()).ServeHTTP

	l.Println("Welcome to barrage. Now listening on localhost, port", config.Port)

	err := http.ListenAndServe(config.Port, httpLogger(router))
	if err != nil {
		panic(err)
	}
}

// cleanly log all HTTP requests.. might not be the cleanest way to do so
func httpLogger(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		router.ServeHTTP(w, req)
		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)
		l.Println(req.Method, req.URL, elapsedTime)
	})
}
