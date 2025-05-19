package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("Internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Errorw("Internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Warnf("Bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {

	app.logger.Warnw("forbidden request request", "method", r.Method, "path", r.URL.Path, "error")
	writeJsonError(w, http.StatusForbidden, "forbidden")
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorw("Not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusNotFound, "not found")
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorf("Conflic response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusConflict, err.Error())
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logger.Errorf("unauthorized  basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJsonError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path, "error", "Rate limit exceeded")

	w.Header().Set("Retry-After", retryAfter)
	writeJsonError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}
