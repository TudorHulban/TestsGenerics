package main

import "github.com/gofiber/fiber/v2"

type server struct {
	engine *fiber.App
	n      *node
	cache  *cache
}

func (s *server) announce(c *fiber.Ctx) error {
	// announcerID := c.Params("id")

	return nil
}
