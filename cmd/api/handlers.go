package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) serveHTML(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "index.html")
}

func (app *application) serveApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("bundle request", r.URL.Path[1:])
	http.ServeFile(w, r, r.URL.Path[1:])
}

func (app *application) checkUserStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("bundle request", r.URL.Path[1:])
	http.ServeFile(w, r, r.URL.Path[1:])
}
