package webhandlers

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"image"
	"image/jpeg"
	"os"
)

func PFPHandler(ctx *fiber.Ctx) error {
	idolized := utils.DetermineIdolizedFromQuery(ctx)
	q := utils.CardQuery{
		IDs:    ctx.Query("id"),
		School: ctx.Query("school"),
		Rarity: ctx.Query("rarity"),
		Name:   ctx.Query("name"),
	}
	cardResult, err := utils.GetCard(q, idolized, false)
	if err != nil {
		return err
	}

	card := cardhandlers.NormalCard{
		BaseCard: *cardResult,
		Idolized: idolized,
	}

	if err := card.ProcessImage(); err != nil {
		return err
	}

	f, _ := os.Open(card.OutputPath())
	img, _, _ := image.Decode(f)

	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	crop, _ := analyzer.FindBestCrop(img, 256, 256)
	croppedImg := imaging.Crop(img, crop)

	var jpgBuff = new(bytes.Buffer)
	jpeg.Encode(jpgBuff, croppedImg, nil)

	ctx.Append("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", card.Hash()+card.Extension()))
	ctx.Type(card.Extension())
	ctx.Send(jpgBuff.Bytes())
	return nil
}
