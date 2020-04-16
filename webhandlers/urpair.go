package webhandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"net/http"
	"strconv"
)

func UrPairHandler(ctx *gin.Context) {
	idolized, err := strconv.ParseBool(ctx.DefaultQuery("idolized", "true"))
	if err != nil {
		_ = ctx.AbortWithError(500, err)
	}

	cardResult, err := utils.SelectRandomCard(ctx)
	if err != nil {
		_ = ctx.AbortWithError(404, err)
	}

	card := cardhandlers.URPair{
		BaseCard: *cardResult,
		Idolized: idolized,
	}

	if err := card.ProcessImage(); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", card.FileBaseName))
	ctx.File(card.OutputPath())
}
