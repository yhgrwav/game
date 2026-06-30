package main

import (
	"game/internal/events"
	"log"
)

func main() {
	hero := events.StartGame()
	log.Println(hero)
}
