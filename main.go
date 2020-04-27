package main

import (
	"github.com/gofiber/fiber"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/webhandlers"
)

func maru(ctx *fiber.Ctx) {
	ctx.SendFile("./cardhandlers/maruexcite.png")
}

func main() {
	router := fiber.New()

	router.Get("/maru", maru)

	router.Use(webhandlers.LimitingMiddleware)
	router.Get("/", webhandlers.NormalCardHandler)
	router.Get("/urpair", webhandlers.UrPairHandler)
	router.Get("/pfp", webhandlers.PFPHandler)

	router.Listen("0.0.0.0:3000")
}
