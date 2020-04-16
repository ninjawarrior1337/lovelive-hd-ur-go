package main

import (
	"github.com/gofiber/fiber"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/webhandlers"
)

func maru(ctx *fiber.Ctx) {
	ctx.SendFile("maruexcite.png")
}

func main() {
	router := fiber.New()

	router.Get("/maru", maru)

	imageHandling := router.Group("/")
	imageHandling.Use(webhandlers.LimitingMiddleware)
	imageHandling.Get("/", webhandlers.NormalCardHandler)
	imageHandling.Get("/urpair", webhandlers.UrPairHandler)
	imageHandling.Get("/pfp")

	router.Listen("0.0.0.0:3000")
}
