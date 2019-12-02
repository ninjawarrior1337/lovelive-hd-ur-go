package cardhandlers

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"lovelive-hd-ur/CardResponse"
	"net/http"
	"os"
)

type URPair struct {
	Waifu2xAble Waifu2xAble
	BaseCard    CardResponse.Result
	image1      image.Image
	image2      image.Image
	idolized    bool
}

func (u *URPair) retrievePair() error {
	if u.BaseCard.CleanUrIdolized == nil && u.BaseCard.UrPair.Card.CleanUrIdolized == nil {
		return errors.New("not a ur pair")
	}

	image1Data, err := http.Get("https:" + *u.BaseCard.CleanUrIdolized)
	if err != nil {
		return err
	}
	defer image1Data.Body.Close()

	u.image1, _, err = image.Decode(image1Data.Body)
	if err != nil {
		return err
	}

	image2Data, err := http.Get("https:" + *u.BaseCard.UrPair.Card.CleanUrIdolized)
	if err != nil {
		return err
	}
	defer image2Data.Body.Close()

	u.image2, _, err = image.Decode(image2Data.Body)
	if err != nil {
		return err
	}

	//Set base name
	u.Waifu2xAble.FileBaseName = fmt.Sprintf("%dx%d", *u.BaseCard.ID, *u.BaseCard.UrPair.Card.ID)

	return nil
}

func (u *URPair) stitchPairAndSave() error {
	if _, err := os.Stat(u.Waifu2xAble.InputDir()); err == nil {
		return nil
	}

	baseImage := gg.NewContext(u.image1.Bounds().Dx()+u.image2.Bounds().Dx(), u.image1.Bounds().Dy())

	switch reverse := *u.BaseCard.UrPair.ReverseDisplayIdolized; reverse {
	case true:
		{
			u.image1, u.image2 = u.image2, u.image1
		}
	}

	baseImage.DrawImage(u.image1, 0, 0)
	baseImage.DrawImage(u.image2, u.image1.Bounds().Dx(), 0)
	err := baseImage.SavePNG(u.Waifu2xAble.OutputDir())
	if err != nil {
		return err
	}

	return nil
}

func (u *URPair) ProcessImage() (err error) {
	err = u.retrievePair()
	err = u.stitchPairAndSave()
	err = u.Waifu2xAble.DoWaifu2x()
	return
}
