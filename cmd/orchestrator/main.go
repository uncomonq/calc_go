package main

import (
	"log"

	"github.com/uncomonq/calc_go/internal/application"
)

func main() {
	ap := application.NewOrchestrator()
	log.Println("Starting Orchestrator on port", ap.Config.Addr)
	if err := ap.RunServer(); err != nil {
		log.Fatal(err)
	}
}