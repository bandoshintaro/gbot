package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Webhook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	//js, err := json.Marshal(genes)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	//w.Write(js)
}
