package webhandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"strconv"
)

func NormalCardHandler(ctx *gin.Context) {
	idolized, err := strconv.ParseBool(ctx.DefaultQuery("idolized", "true"))
	cardResult, err := utils.SelectRandomCard(ctx)
	if err != nil {
		_ = ctx.AbortWithError(404, err)
	}
	card := cardhandlers.NormalCard{
		Waifu2xAble: cardhandlers.Waifu2xAble{
			FileBaseName: strconv.FormatInt(*cardResult.ID, 10) + strconv.FormatBool(idolized) + ".png",
		},
		BaseCard: *cardResult,
		Idolized: idolized,
	}

	if err := card.ProcessImage(); err != nil {
		_ = ctx.AbortWithError(500, err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", card.FileBaseName))
	ctx.File(card.OutputPath())
	return
}
