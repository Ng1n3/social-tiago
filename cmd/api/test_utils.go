package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ng1n3/social/internal/auth"
	"github.com/Ng1n3/social/internal/store"
	"github.com/Ng1n3/social/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	// logger := zap.Must(zap.NewProduction()).Sugar()
	logger := zap.NewNop().Sugar()
	mockstore := store.NewMockStore()
	// mockCacheStore := cache.NewMockStore()
	mockCacheStore := &cache.MockUserStore{}
	testAuth := &auth.TestAuthenticator{}

	return &application{
		logger:        logger,
		store:         mockstore,
		cacheStorage:  cache.Storage{Users: mockCacheStore},
		authenticator: testAuth,
		config: cfg,
	}
}

func executeRequests(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected response code %d. Got %d", expected, actual)
	}
}
