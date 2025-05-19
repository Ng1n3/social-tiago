package main

import (
	"net/http"
)

// healthCheck godoc
//
//	@Summary		checks health of API
//	@Description	CHecks if API is up and running
//	@Tags			ops
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	// time.Sleep(time.Second * 3)
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
