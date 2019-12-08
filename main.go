package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/CardResponse"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
)

var cardJobs = make(chan struct{}, 2)

func addToQueryIfNotExists(query *url.Values, param string, value string) {
	if value != "" {
		query.Add(param, value)
	}
}

func selectRandomCard(ctx *gin.Context) (*CardResponse.Result, error) {
	parsed, err := url.Parse("https://schoolido.lu/api/cards/")
	if err != nil {
		return nil, err
	}
	q := parsed.Query()
	q.Add("ordering", "random")
	q.Add("expand_ur_pair", "true")
	addToQueryIfNotExists(&q, "ids", ctx.DefaultQuery("id", ""))
	addToQueryIfNotExists(&q, "school", ctx.DefaultQuery("school", "Otonokizaka Academy, Uranohoshi Girls' High School"))
	addToQueryIfNotExists(&q, "rarity", ctx.DefaultQuery("rarity", "SSR,UR"))
	parsed.RawQuery = q.Encode()
	fmt.Println(parsed)

	resp, err := http.Get(parsed.String())
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := CardResponse.UnmarshalCardResponse(body)

	return &cardResponse.Results[0], nil
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
	idolized, err := strconv.ParseBool(ctx.DefaultQuery("idolized", "true"))
	cardResult, err := selectRandomCard(ctx)
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
		return
	}

	ctx.File(card.Waifu2xAble.OutputDir())
	return
}

func urPairs(ctx *gin.Context) {
	idolized, err := strconv.ParseBool(ctx.DefaultQuery("idolized", "true"))
	if err != nil {
		_ = ctx.AbortWithError(500, err)
	}
	cardResult, err := selectRandomCard(ctx)
	if err != nil {
		_ = ctx.AbortWithError(404, err)
	}

	card := cardhandlers.URPair{
		Waifu2xAble: cardhandlers.Waifu2xAble{},
		BaseCard:    *cardResult,
		Idolized:    idolized,
	}

	if err := card.ProcessImage(); err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.File(card.Waifu2xAble.OutputDir())
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
