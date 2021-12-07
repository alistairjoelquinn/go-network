package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	router.GET("/", app.serveHTML)
	router.GET("/favicon.ico", app.serveFavicon)
	router.GET("/bundle.js", app.serveApp)

	// API routes
	router.GET("/api/auth/id", app.checkUserStatus)

	return router
}
