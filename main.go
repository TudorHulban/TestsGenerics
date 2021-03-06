package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/TudorHulban/log"
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
		l:      log.NewLogger(log.DEBUG, os.Stdout, true),
	}

	s.engine.Get("/"+urlAnnounce+"/:id", s.announce)
	s.engine.Get("/"+urlRing, s.logRing)
	s.engine.Get("/"+urlPartitions+"/:id", s.partitions)

	if s.n.isRoot() {
		fmt.Printf("Root node with ID %d started.\n", id)
	} else {
		go s.n.announceTo(s.n.rootID)
		go s.n.registerNodeID(s.n.rootID)
	}

	fmt.Printf("Node listening on port %s.\n", s.n.listenOn())

	s.engine.Listen(s.n.listenOn())
}
