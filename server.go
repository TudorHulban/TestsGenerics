package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	engine *fiber.App
	n      *node
}

func (s *server) announce(c *fiber.Ctx) error {
	announcerID := c.Params("id")

	id, errCo := strconv.Atoi(announcerID)
	if errCo != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	fmt.Printf("Node with ID: %s announced itself.\n", announcerID)
	s.n.registerNodeID(id)

	fmt.Printf("Neighbors: %d\n", len(*s.n.getRing()))

	return c.JSON(s.n.getRing())
}
