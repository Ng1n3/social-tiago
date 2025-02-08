package main

import (
	"log"

	"github.com/Ng1n3/social/internal/env"
	"github.com/Ng1n3/social/internal/store"
)

func main() {
	cfg := config{
    addr: env.GetString("ADDR", ":3050"),
	}
  
  store := store.NewStorage(nil)

	app := &application{
		config: cfg,
    store: store,
	}


	mux := app.mount()
	log.Fatal(app.run(mux))
}
