package main

import (
	"game/internal/events"
	"log"
)

func main() {
	params := events.StartGame()
	log.Println(params)
}
