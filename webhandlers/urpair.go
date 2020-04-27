package webhandlers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"net/http"
)

func UrPairHandler(ctx *fiber.Ctx) {
	idolized := utils.DetermineIdolizedFromQuery(ctx)
	q := utils.CardQuery{
		IDs:    ctx.Query("id"),
		School: ctx.Query("school"),
		Rarity: ctx.Query("rarity"),
		Name:   ctx.Query("name"),
	}
	cardResult, err := utils.GetCard(q, idolized, true)
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
	ctx.Append("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", card.FileBaseName))
	ctx.SendFile(card.OutputPath())
}
