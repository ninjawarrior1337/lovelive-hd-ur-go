package cardhandlers

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"golang.org/x/sync/errgroup"
	"image"
	"net/http"
	"os"
)

type URPair struct {
	Waifu2xAble
	BaseCard utils.Result
	image1   image.Image
	image2   image.Image
	Idolized bool
}

func (u *URPair) retrievePair() error {
	waiter := errgroup.Group{}
	var baseCardUrl string
	var pairCardUrl string

	switch u.Idolized {
	case true:
		baseCardUrl = u.BaseCard.CleanUrIdolized
		pairCardUrl = u.BaseCard.UrPair.Card.CleanUrIdolized
	case false:
		baseCardUrl = u.BaseCard.CleanUr
		pairCardUrl = u.BaseCard.UrPair.Card.CleanUr
	}
	//Set base name
	u.FileBaseName = fmt.Sprintf("%dx%d%v.png", *u.BaseCard.ID, *u.BaseCard.UrPair.Card.ID, u.Idolized)

	waiter.Go(func() error {
		image1Data, err := http.Get("https:" + baseCardUrl)
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
		image2Data, err := http.Get("https:" + pairCardUrl)
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
		return fmt.Errorf("failed to obtain images: %s", err)
	}

	return nil
}

func (u *URPair) stitchPairAndSave() error {
	if _, err := os.Stat(u.InputPath()); err == nil {
		return err
	}

	baseImage := gg.NewContext(u.image1.Bounds().Dx()+u.image2.Bounds().Dx(), u.image1.Bounds().Dy())

	switch u.Idolized {
	case true:
		switch *u.BaseCard.UrPair.ReverseDisplayIdolized {
		case true:
			{
				u.image1, u.image2 = u.image2, u.image1
			}
		}
	case false:
		switch *u.BaseCard.UrPair.ReverseDisplay {
		case true:
			{
				u.image1, u.image2 = u.image2, u.image1
			}
		}
	}

	baseImage.DrawImage(u.image1, 0, 0)
	baseImage.DrawImage(u.image2, u.image1.Bounds().Dx(), 0)
	err := baseImage.SavePNG(u.InputPath())
	if err != nil {
		return fmt.Errorf("failed to save image: %s", err)
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
	err = u.DoWaifu2x()
	if err != nil {
		return
	}
	return
}
