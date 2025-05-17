package main

import (
	"go-unit-of-work-example/internal/app"
	"go-unit-of-work-example/internal/config"
	"log"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	a := app.New(cfg)

	if err := a.Build(); err != nil {
		log.Fatalf("failed to build app: %s", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}

	log.Printf("App successfully stopeed")
}
