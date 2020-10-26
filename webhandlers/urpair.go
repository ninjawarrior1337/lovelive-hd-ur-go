package webhandlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
)

func UrPairHandler(ctx *fiber.Ctx) error {
	idolized := utils.DetermineIdolizedFromQuery(ctx)
	q := utils.CardQuery{
		IDs:    ctx.Query("id"),
		School: ctx.Query("school"),
		Rarity: ctx.Query("rarity"),
		Name:   ctx.Query("name"),
	}
	cardResult, err := utils.GetCard(q, idolized, true)
	if err != nil {
		return err
	}

	card := cardhandlers.URPair{
		BaseCard: *cardResult,
		Idolized: idolized,
	}

	if err := card.ProcessImage(); err != nil {
		return err
	}

	ctx.Append("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", card.Hash()+card.Extension()))
	ctx.SendFile(card.OutputPath())
	return nil
}
