package cardhandlers

import (
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"image"
	"net/http"
)

type NormalCard struct {
	Waifu2xAble
	BaseCard utils.Result
	Idolized bool
}

func (card *NormalCard) writeBaseCard() error {
	var cardData *http.Response
	switch card.Idolized {
	case true:
		var err error
		cardData, err = http.Get("https:" + card.BaseCard.CleanUrIdolized)
		if err != nil {
			return err
		}
	case false:
		var err error
		cardData, err = http.Get("https:" + card.BaseCard.CleanUr)
		if err != nil {
			return err
		}
	}
	defer cardData.Body.Close()

	i, _, err := image.Decode(cardData.Body)
	if err != nil {
		return err
	}

	card.Image = i

	return nil
}

func (card *NormalCard) ProcessImage() (err error) {
	err = card.writeBaseCard()
	if err != nil {
		return
	}
	err = card.DoWaifu2x()
	if err != nil {
		return
	}
	return
}
