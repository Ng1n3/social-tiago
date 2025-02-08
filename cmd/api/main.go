package main

import (
	"log"

	"github.com/Ng1n3/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":3050"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
