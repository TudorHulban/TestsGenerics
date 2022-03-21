package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ptrID := flag.Int("id", 123, "node own ID")
	ptrRootID := flag.Int("rootid", 123, "joining root with ID")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Printf("ID of starting node not provided\n")
		os.Exit(1)
	}

	node := func(ptr *int) *node {
		if ptr != nil {
			return newRoot(*ptrID)
		}

		return newNode(*ptrID)
	}

	s := server{
		engine: fiber.New(),
		n:      node(ptrRootID),
	}

	s.engine.Get("/myidis/:id", s.announce)

	if s.n.isRoot {
		fmt.Printf("Root node with ID %d started.\n", ptrID)
	}

	fmt.Printf("Node listening on port %s", s.n.listenOn())

	log.Fatal(s.engine.Listen(s.n.listenOn()))
}
