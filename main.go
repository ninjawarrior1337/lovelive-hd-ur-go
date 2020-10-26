package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/webhandlers"
)

func maru(ctx *fiber.Ctx) error {
	ctx.SendFile("./cardhandlers/maruexcite.png")
	return nil
}

func main() {
	router := fiber.New()

	router.Get("/maru", maru)

	router.Use(webhandlers.LimitingMiddleware)
	router.Get("/", webhandlers.NormalCardHandler)
	router.Post("/", webhandlers.GenericHandler)
	router.Get("/urpair", webhandlers.UrPairHandler)
	router.Get("/pfp", webhandlers.PFPHandler)

	panic(router.Listen("0.0.0.0:3000"))
}
