package utils

import (
	"fmt"
	"github.com/gofiber/fiber"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SelectRandomCard(ctx *fiber.Ctx) (*Result, error) {
	parsed, err := url.Parse("https://schoolido.lu/api/cards/")
	if err != nil {
		return nil, err
	}
	q := parsed.Query()

	q.Add("ordering", "random")
	q.Add("expand_ur_pair", "true")
	q.Add("school", "Otonokizaka Academy, Uranohoshi Girls' High School")
	q.Add("rarity", "SSR,UR")

	if id := ctx.Query("id"); id != "" {
		q.Add("ids", ctx.Query("id"))
	}
	if school := ctx.Query("school"); school != "" {
		q.Add("school", school)
	}
	if rarity := ctx.Query("rarity"); rarity != "" {
		q.Add("rarity", rarity)
	}
	if name := ctx.Query("name"); name != "" {
		q.Add("name", name)
	}

	parsed.RawQuery = q.Encode()
	log.Println("Query URL: " + parsed.String())

	resp, err := http.Get(parsed.String())
	if err != nil {
		ctx.Status(500)
		ctx.SendString(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := UnmarshalCardResponse(body)

	if len(cardResponse.Results) == 0 {
		return nil, fmt.Errorf("no cards found that match query")
	}

	return &cardResponse.Results[0], nil
}
