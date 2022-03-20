package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func hasher(s string) string {
	return (s + "000000000")[:10]
}

func hashNode(n node) string {
	return hasher(strconv.Itoa(n.id))
}

func main() {
	s := server{
		engine: fiber.New(),
		n:      newNode(7),
		cache:  newCache(),
	}

	s.engine.Get("/myidis/:id", s.announce)

	log.Fatal(s.engine.Listen(s.n.listenOn()))
}
