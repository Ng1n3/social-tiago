package main

import (
	"net/http"
	"testing"

	"github.com/Ng1n3/social/internal/auth"
)

func TestGetUser(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()
	testToken, err := app.authenticator.GenerateToken(auth.TestClaims)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequests(req, mux)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)

	})
	t.Run("Should allow authenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequests(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}
