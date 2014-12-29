 package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/fortytw2/barrage/app/home"
	"github.com/fortytw2/barrage/app/detail"
	"github.com/fortytw2/barrage/cache"
	"github.com/fortytw2/barrage/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

//go:generate bash -c "lessc assets/css/main.less | cleancss > static/css/main.min.css"

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
}

func main() {
	router := mux.NewRouter()

	cache.ParseFiles(config.Config.SourceFolder)

	router.HandleFunc("/", home.GetHomePage).Methods("GET")
	router.HandleFunc("/about", home.GetAboutPage).Methods("GET")
	router.HandleFunc("/detail", detail.GetDetailPage).Methods("GET")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(rice.MustFindBox("static").HTTPBox())))
	http.Handle("/video/", http.StripPrefix("/video/", http.FileServer(http.Dir(config.Config.StorageFolder))))
	http.Handle("/", httpLogger(router))

	l.Println("Welcome to barrage. Now listening on localhost, port", config.Config.Port)

	http.ListenAndServe(config.Config.Port, nil)

	// this is a test
}

func httpLogger(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		router.ServeHTTP(w, req)
		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)
		l.Println(req.Method, req.URL, elapsedTime)
	})
}
