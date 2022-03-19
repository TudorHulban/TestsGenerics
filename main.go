package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func hasher(s string) string {
	return (s + "000000000")[:10]
}

func hashNode(n node) string {
	return hasher(n.id)
}

func main() {
	s := server{
		engine: fiber.New(),
		cache:  newCache(),
	}

	s.engine.Get("/myidis/:id", s.announce)

	log.Fatal(s.engine.Listen(":8080"))
}
