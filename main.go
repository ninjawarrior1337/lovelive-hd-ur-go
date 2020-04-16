package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/webhandlers"
)

func maru(ctx *gin.Context) {
	ctx.File("maruexcite.png")
}

func main() {
	router := gin.Default()

	router.GET("/maru", maru)

	imageHandling := router.Group("/")
	imageHandling.Use(webhandlers.LimitingMiddleware)
	imageHandling.GET("/", webhandlers.NormalCardHandler)
	imageHandling.GET("/urpair", webhandlers.UrPairHandler)
	imageHandling.GET("/pfp")

	router.Run("0.0.0.0:3000")
}
