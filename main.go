package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := server{
		engine: fiber.New(),
		n:      newNode(7, true),
	}

	s.engine.Get("/myidis/:id", s.announce)

	log.Fatal(s.engine.Listen(s.n.listenOn()))
}
