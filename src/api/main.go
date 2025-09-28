package main

import (
	"log"

	"github.com/nitishsharma2825/social/internal"
	"github.com/nitishsharma2825/social/internal/repository"
)

func main() {
	// load the configuration settings
	settings, err := internal.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cfg := config{
		addr: settings.Addr,
	}

	storage := repository.NewPostgresStorage(nil)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
