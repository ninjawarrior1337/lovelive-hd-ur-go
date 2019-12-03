package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lovelive-hd-ur/CardResponse"
	"lovelive-hd-ur/cardhandlers"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
)

var cardJobs = make(chan struct{}, 2)

func selectRandomCard() (*CardResponse.Result, error) {

}

func LimitingMiddleware(c *gin.Context) {
	select {
	case cardJobs <- struct{}{}:
		{
			defer func() { <-cardJobs }()
			c.Next()
		}
	default:
		{
			c.AbortWithStatus(503)
		}
	}
}

func root(ctx *gin.Context) {
	lsOut, _ := exec.Command("ls").Output()

	_, _ = fmt.Fprint(ctx.Writer, string(lsOut))
}

func maru(ctx *gin.Context) {
	ctx.File("maruexcite.png")
}

func normalCards(ctx *gin.Context) {
	parsed, _ := url.Parse("https://schoolido.lu/api/cards/")
	q := parsed.Query()
	q.Add("ids", ctx.Query("id"))
	parsed.RawQuery = q.Encode()
	fmt.Println(parsed)
	resp, err := http.Get(parsed.String())
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := CardResponse.UnmarshalCardResponse(body)

	if cardResponse.Results[0].CleanUrIdolized == nil {
		_ = ctx.AbortWithError(http.StatusNotFound, errors.New("card has no idolized ur"))
		return
	}

	card := cardhandlers.NormalCard{
		Waifu2xAble: cardhandlers.Waifu2xAble{
			FileBaseName: strconv.FormatInt(*cardResponse.Results[0].ID, 10) + ".png",
		},
		BaseCard: cardResponse.Results[0],
	}

	if err := card.ProcessImage(); err != nil {
		_ = ctx.AbortWithError(500, err)
		return
	}

	ctx.File(card.Waifu2xAble.OutputDir())
	return
}

func urPairs(c *gin.Context) {

}

func main() {
	router := gin.Default()

	router.GET("/list", root)
	router.GET("/maru", maru)

	imageHandling := router.Group("/")
	imageHandling.Use(LimitingMiddleware)
	imageHandling.GET("/", normalCards)
	imageHandling.GET("/urpair", urPairs)

	router.Run("0.0.0.0:5005")
	// http.HandleFunc("/", root)
	// http.HandleFunc("/ggg", ggg)
	// http.HandleFunc("/cards", cards)

	// http.ListenAndServe(":5005", nil)
}
