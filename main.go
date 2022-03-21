package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 7, "node own ID")
	flag.Parse()

	node := func() *node {
		if len(flag.Args()) == 0 {
			return newRoot(id)
		}

		rootID, errConv := strconv.Atoi(flag.Arg(0))
		if errConv != nil {
			fmt.Printf("Conversion of root ID: %s", errConv.Error())
			os.Exit(1)
		}

		return newNode(id, rootID)
	}

	s := server{
		engine: fiber.New(),
		n:      node(),
	}

	s.engine.Get("/myidis/:id", s.announce)

	if s.n.isRoot() {
		fmt.Printf("Root node with ID %d started.\n", id)
	}

	fmt.Printf("Node listening on port %s.\n", s.n.listenOn())

	log.Fatal(s.engine.Listen(":" + s.n.listenOn()))
}
