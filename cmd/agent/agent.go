package main

import (
	"log"

	"github.com/uncomonq/calc_go/internal/application"
)

func main() {
	agent := application.NewAgent()
	log.Println("Starting Agent...")
	agent.Run()
}