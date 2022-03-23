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

	errRe := s.n.registerNodeID(id)
	if errRe != nil {
		return c.Status(http.StatusInternalServerError).SendString(errRe.Error())
	}

	fmt.Printf("Neighbors: %d\n", len(*s.n.getRing(s.n.getNodeData()))-1)

	return c.JSON(s.n.getRing(s.n.getNodeData()))
}

func (s *server) logRing(c *fiber.Ctx) error {
	s.n.neighborsTo(c)

	return nil
}
