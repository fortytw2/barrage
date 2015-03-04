package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/fortytw2/barrage/api"
	"github.com/fortytw2/barrage/config"
	"github.com/fortytw2/barrage/models"
	"github.com/julienschmidt/httprouter"
)

func runWeb() {
	router := httprouter.New()

	db := models.OpenDB()

	// series are different from individual movies, but not by much.
	router.GET("/api/series", api.GetSeries(db))
	router.GET("/api/series/:id", api.GetSeriesDetail(db))

	router.ServeFiles("/video/*filepath", http.Dir(config.VideoFolder))

	router.NotFound = http.FileServer(rice.MustFindBox("static").HTTPBox()).ServeHTTP

	log.Println("Welcome to barrage. Now listening on localhost, port", config.Port)

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
		log.Println(req.Method, req.URL, elapsedTime)
	})
}
