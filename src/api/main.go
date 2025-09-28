package main

import (
	"log"

	"github.com/nitishsharma2825/social/internal"
	"github.com/nitishsharma2825/social/internal/db"
	"github.com/nitishsharma2825/social/internal/repository"
)

func main() {
	// load the configuration settings
	settings, err := internal.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// config for the application
	cfg := config{
		addr:   settings.Addr,
		dbConn: settings.Database.GetConnectionString(),
	}

	// create the db connection pool
	db, err := db.New(cfg.dbConn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// Ideally, storage should be linked to service layer not directly to app layer
	storage := repository.NewPostgresStorage(db)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
