package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	engine *fiber.App
	n      *node
}

func (s *server) announce(c *fiber.Ctx) error {
	announcerID := c.Params("id")

	fmt.Printf("Node with ID: %s announced itself.\n", announcerID)

	return nil
}
