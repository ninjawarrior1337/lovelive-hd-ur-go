package webhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"strconv"
)

func NormalCardHandler(ctx *fiber.Ctx) {
	idolized, err := strconv.ParseBool(ctx.Query("idolized"))
	if err != nil {
		idolized = true
	}
	cardResult, err := utils.SelectRandomCard(ctx)
	if err != nil {
		ctx.SendStatus(404)
		ctx.SendString("Failed to select random card " + err.Error())
		return
	}
	card := cardhandlers.NormalCard{
		Waifu2xAble: cardhandlers.Waifu2xAble{
			FileBaseName: strconv.FormatInt(*cardResult.ID, 10) + strconv.FormatBool(idolized) + ".png",
		},
		BaseCard: *cardResult,
		Idolized: idolized,
	}

	if err := card.ProcessImage(); err != nil {
		ctx.SendStatus(500)
		ctx.JSON(gin.H{"error": err.Error()})
		return
	}

	ctx.Download(card.OutputPath(), card.FileBaseName)
	return
}
