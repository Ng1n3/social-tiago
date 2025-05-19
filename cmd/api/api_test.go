package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Ng1n3/social/internal/ratelimiter"
)

func TestRateLimitterMiddleware(t *testing.T) {
	cfg := config{
		rateLimiter: ratelimiter.Config{
			RequestPerTimeFrame: 2,
			TimeFrame:           time.Second * 5,
			Enabled:             true,
		},
		addr: ":5000",
	}

	app := newTestApplication(t, cfg)

	app.rateLimiter = ratelimiter.NewFixeWindowLimiter(
		cfg.rateLimiter.RequestPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	if app.rateLimiter == nil {
		t.Fatal("rate limiter was not initialized")
	}

	ts := httptest.NewServer(app.mount())
	defer ts.Close()

	client := &http.Client{}
	mockIP := "192.168.1.1"
	marginOfError := 2

	for i := 0; i < cfg.rateLimiter.RequestPerTimeFrame+marginOfError; i++ {
		req, err := http.NewRequest("GET", ts.URL+"/v1/health", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		req.Header.Set("X-Forwarded-For", mockIP)

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}

		defer resp.Body.Close()

		if i < cfg.rateLimiter.RequestPerTimeFrame {
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", resp.Status)
			}
		} else {
			if resp.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected status too many requests; got %v", resp.Status)
			}
		}
	}
}
