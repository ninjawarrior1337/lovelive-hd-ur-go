package cardhandlers

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/CardResponse"
	"golang.org/x/sync/errgroup"
	"image"
	"net/http"
	"os"
)

type URPair struct {
	Waifu2xAble Waifu2xAble
	BaseCard    CardResponse.Result
	image1      image.Image
	image2      image.Image
	Idolized    bool
}

func (u *URPair) retrievePair() error {
	waiter := errgroup.Group{}
	var baseCardUrl *string
	var pairCardUrl *string

	if u.Idolized {
		baseCardUrl = u.BaseCard.CleanUrIdolized
		pairCardUrl = u.BaseCard.UrPair.Card.CleanUrIdolized
	} else {
		baseCardUrl = u.BaseCard.CleanUr
		pairCardUrl = u.BaseCard.UrPair.Card.CleanUr
	}

	if baseCardUrl == nil && pairCardUrl == nil {
		return errors.New("not a ur pair")
	}

	waiter.Go(func() error {
		image1Data, err := http.Get("https:" + *baseCardUrl)
		if err != nil {
			return err
		}
		defer image1Data.Body.Close()

		u.image1, _, err = image.Decode(image1Data.Body)
		if err != nil {
			return err
		}
		return nil
	})

	waiter.Go(func() error {
		image2Data, err := http.Get("https:" + *pairCardUrl)
		if err != nil {
			return err
		}
		defer image2Data.Body.Close()

		u.image2, _, err = image.Decode(image2Data.Body)
		if err != nil {
			return err
		}
		return nil
	})

	err := waiter.Wait()
	if err != nil {
		return err
	}

	//Set base name
	u.Waifu2xAble.FileBaseName = fmt.Sprintf("%dx%d%v.png", *u.BaseCard.ID, *u.BaseCard.UrPair.Card.ID, u.Idolized)

	return nil
}

func (u *URPair) stitchPairAndSave() error {
	if _, err := os.Stat(u.Waifu2xAble.InputDir()); err == nil {
		return nil
	}

	baseImage := gg.NewContext(u.image1.Bounds().Dx()+u.image2.Bounds().Dx(), u.image1.Bounds().Dy())

	if u.Idolized {
		switch reverse := *u.BaseCard.UrPair.ReverseDisplayIdolized; reverse {
		case true:
			{
				u.image1, u.image2 = u.image2, u.image1
			}
		}
	} else {
		switch reverse := *u.BaseCard.UrPair.ReverseDisplay; reverse {
		case true:
			{
				u.image1, u.image2 = u.image2, u.image1
			}
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
	if err != nil {
		return
	}
	err = u.stitchPairAndSave()
	if err != nil {
		return
	}
	//err = u.Waifu2xAble.DoWaifu2x()
	//if err != nil {
	//	return
	//}
	return
}
