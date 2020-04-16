package utils

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func addToQueryIfNotExists(query *url.Values, param string, value string) {
	if value != "" {
		query.Add(param, value)
	}
}

func SelectRandomCard(ctx *gin.Context) (*Result, error) {
	parsed, err := url.Parse("https://schoolido.lu/api/cards/")
	if err != nil {
		return nil, err
	}
	q := parsed.Query()
	q.Add("ordering", "random")
	q.Add("expand_ur_pair", "true")
	AddToQueryIfNotExists(&q, "ids", ctx.DefaultQuery("id", ""))
	AddToQueryIfNotExists(&q, "school", ctx.DefaultQuery("school", "Otonokizaka Academy, Uranohoshi Girls' High School"))
	AddToQueryIfNotExists(&q, "rarity", ctx.DefaultQuery("rarity", "SSR,UR"))
	parsed.RawQuery = q.Encode()
	log.Println("Query URL: " + parsed.String())

	resp, err := http.Get(parsed.String())
	if err != nil {
		_ = ctx.AbortWithError(500, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := UnmarshalCardResponse(body)

	return &cardResponse.Results[0], nil
}
