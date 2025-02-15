package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("Internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Errorw("Internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSON(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Warnf("Bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Errorw("Not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusNotFound, "not found")
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Errorf("Conflic response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSON(w, http.StatusConflict, err.Error())
}
