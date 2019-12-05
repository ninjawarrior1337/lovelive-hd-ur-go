package cardhandlers

import (
	"errors"
	"io/ioutil"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/CardResponse"
	"net/http"
	"os"
)

type NormalCard struct {
	Waifu2xAble Waifu2xAble
	BaseCard    CardResponse.Result
	idolized    bool
}

func (card *NormalCard) writeBaseCard() error {
	if card.BaseCard.CleanUrIdolized == nil {
		return errors.New("selected card has no UR")
	}
	cardData, err := http.Get("https:" + *card.BaseCard.CleanUrIdolized)
	if err != nil {
		return err
	}
	defer cardData.Body.Close()

	cardDataReader, err := ioutil.ReadAll(cardData.Body)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(card.Waifu2xAble.InputDir(), cardDataReader, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (card *NormalCard) ProcessImage() (err error) {
	err = card.writeBaseCard()
	err = card.Waifu2xAble.DoWaifu2x()
	return
}
