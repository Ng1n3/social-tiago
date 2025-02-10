package main

import (
	"log"

	"github.com/Ng1n3/social/internal/db"
	"github.com/Ng1n3/social/internal/env"
	"github.com/Ng1n3/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:superpassword@localhost:5434/social?sslmode=disable")
	if addr == "" {
		log.Fatal("DB_ADDR environment variable is not set")
	}

	conn, err := db.New(addr, 3, 3, "15m")

	if err != nil {
		log.Fatal(err)
	}
	store := store.NewStorage(conn)
	defer conn.Close()
	db.Seed(store)
}
