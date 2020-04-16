package utils

import (
	"github.com/gofiber/fiber"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var defaultSchools = "Otonokizaka Academy, Uranohoshi Girls' High School"
var defaultRarities = "SSR,UR"

func addToQueryIfNotExists(query *url.Values, param string, value string) {
	if value != "" {
		query.Add(param, value)
	}
}

func SelectRandomCard(ctx *fiber.Ctx) (*Result, error) {
	parsed, err := url.Parse("https://schoolido.lu/api/cards/")
	if err != nil {
		return nil, err
	}
	q := parsed.Query()
	q.Add("ordering", "random")
	q.Add("expand_ur_pair", "true")
	addToQueryIfNotExists(&q, "ids", ctx.Query("id"))
	if school := ctx.Query("school"); school != "" {
		addToQueryIfNotExists(&q, "school", school)
	} else {
		addToQueryIfNotExists(&q, "school", defaultSchools)
	}

	if rarity := ctx.Query("rarity"); rarity != "" {
		addToQueryIfNotExists(&q, "rarity", rarity)
	} else {
		addToQueryIfNotExists(&q, "rarity", defaultRarities)
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

	return &cardResponse.Results[0], nil
}
