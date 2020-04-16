package webhandlers

import (
	"github.com/gofiber/fiber"
)

var cardJobs = make(chan struct{}, 2)

func LimitingMiddleware(c *fiber.Ctx) {
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
}
