package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func DetermineIdolizedFromQuery(ctx *fiber.Ctx) bool {
	idolized, err := strconv.ParseBool(ctx.Query("idolized"))
	if err != nil {
		idolized = true
	}
	return idolized
}

func GenerateURL(query CardQuery) *url.URL {
	parsed, err := url.Parse("https://schoolido.lu/api/cards/")
	if err != nil {
		panic("what the frick just happened")
	}
	q := parsed.Query()

	q.Add("ordering", "random")
	q.Add("expand_ur_pair", "true")
	q.Add("school", "Otonokizaka Academy, Uranohoshi Girls' High School")
	q.Add("rarity", "SSR,UR")

	if id := query.IDs; id != "" {
		q.Add("ids", id)
	}
	if school := query.School; school != "" {
		q.Add("school", school)
	}
	if rarity := query.Rarity; rarity != "" {
		q.Add("rarity", rarity)
	}
	if name := query.Name; name != "" {
		q.Add("name", name)
	}

	parsed.RawQuery = q.Encode()
	log.Println("Query URL: " + parsed.String())

	return parsed
}

func GetCard(query CardQuery, idolized, urpair bool) (*Result, error) {
	parsed := GenerateURL(query)
	resp, err := http.Get(parsed.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := UnmarshalCardResponse(body)

	if len(cardResponse.Results) == 0 {
		return nil, fmt.Errorf("no cards found that match query")
	}
	var selectedCard *Result

	for _, result := range cardResponse.Results {
		if urpair {
			if result.UrPair.Card == nil {
				continue
			}
			if idolized {
				if result.CleanUrIdolized != "" && result.UrPair.Card.CleanUrIdolized != "" {
					selectedCard = &result
					break
				}
			} else {
				if result.CleanUr != "" && result.UrPair.Card.CleanUr != "" {
					selectedCard = &result
					break
				}
			}
		} else {
			if idolized {
				if result.CleanUrIdolized != "" {
					selectedCard = &result
					break
				}
			} else {
				if result.CleanUr != "" {
					selectedCard = &result
					break
				}
			}
		}
	}

	if selectedCard == nil {
		return nil, &CardNotFoundError{}
	}

	return selectedCard, nil
}
