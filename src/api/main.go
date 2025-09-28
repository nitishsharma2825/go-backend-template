package main

import (
	"log"

	"github.com/nitishsharma2825/social/internal"
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

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
