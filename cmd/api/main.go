package main

import (
	"time"

	"github.com/Ng1n3/social/internal/db"
	"github.com/Ng1n3/social/internal/env"
	"github.com/Ng1n3/social/internal/store"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			GopherSOcial API
//	@version		1.0
//	@description	API for GopherSocial, a social network for gophers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":3050"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:5000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:superpassword@localhost:5434/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3, // 3 days,
		},
	}

	// Logger
	logger := zap.Must(zap.NewDevelopment()).Sugar()
	defer logger.Sync()

	//Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
