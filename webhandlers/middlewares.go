package webhandlers

import (
	"github.com/gofiber/fiber/v2"
)

var cardJobs = make(chan struct{}, 2)

func LimitingMiddleware(c *fiber.Ctx) error {
	select {
	case cardJobs <- struct{}{}:
		{
			defer func() { <-cardJobs }()
			c.Next()
		}
	default:
		{
			c.Status(503)
		}
	}
	return nil
}
