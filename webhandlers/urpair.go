package webhandlers

import (
	"github.com/gofiber/fiber"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"net/http"
	"strconv"
)

func UrPairHandler(ctx *fiber.Ctx) {
	idolized, err := strconv.ParseBool(ctx.Query("idolized"))
	if err != nil {
		idolized = true
	}
	cardResult, err := utils.SelectRandomCard(ctx)
	if err != nil {
		ctx.Status(404)
		ctx.SendString("Failed to select card " + err.Error())
		return
	}

	card := cardhandlers.URPair{
		BaseCard: *cardResult,
		Idolized: idolized,
	}

	if err := card.ProcessImage(); err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.SendString(err.Error())
		return
	}

	ctx.Download(card.OutputPath(), card.FileBaseName)
}
