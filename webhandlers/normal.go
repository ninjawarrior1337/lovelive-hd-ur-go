package webhandlers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"strconv"
)

func NormalCardHandler(ctx *fiber.Ctx) {
	idolized := utils.DetermineIdolizedFromQuery(ctx)
	q := utils.CardQuery{
		IDs:    ctx.Query("id"),
		School: ctx.Query("school"),
		Rarity: ctx.Query("rarity"),
		Name:   ctx.Query("name"),
	}
	cardResult, err := utils.GetCard(q, idolized, false)
	if err != nil {
		ctx.SendStatus(404)
		ctx.SendString("Failed to select card " + err.Error())
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
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.Append("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", card.FileBaseName))
	ctx.SendFile(card.OutputPath())
	return
}
