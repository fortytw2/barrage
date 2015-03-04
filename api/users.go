package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	js, err := json.Marshal(map[string]string{"wow": "wow"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	js, err := json.Marshal(map[string]string{"wow": "wow"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	js, err := json.Marshal(map[string]string{"wow": "wow"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
