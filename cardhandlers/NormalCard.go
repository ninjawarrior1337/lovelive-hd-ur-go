package cardhandlers

import (
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/CardResponse"
	"io"
	"net/http"
	"os"
)

type NormalCard struct {
	Waifu2xAble
	BaseCard CardResponse.Result
	Idolized bool
}

func (card *NormalCard) writeBaseCard() error {
	//Check if card already exists
	if _, err := os.Stat(card.InputPath()); err == nil {
		return nil
	}
	//If not, re-download
	var cardData *http.Response
	switch card.Idolized {
	case true:
		var err error
		if card.BaseCard.CleanUrIdolized == nil {
			return &CardNotFoundError{*card.BaseCard.ID, card.Idolized}
		}
		cardData, err = http.Get("https:" + *card.BaseCard.CleanUrIdolized)
		if err != nil {
			return err
		}
	case false:
		var err error
		if card.BaseCard.CleanUr == nil {
			return &CardNotFoundError{*card.BaseCard.ID, card.Idolized}
		}
		cardData, err = http.Get("https:" + *card.BaseCard.CleanUr)
		if err != nil {
			return err
		}
	}
	defer cardData.Body.Close()

	f, err := os.Create(card.InputPath())
	if err != nil {
		return err
	}

	_, _ = io.Copy(f, cardData.Body)

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
