package cardhandlers

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/utils"
	"golang.org/x/sync/errgroup"
	"image"
	"net/http"
)

type URPair struct {
	Waifu2xAble
	BaseCard utils.Result
	p1       image.Image
	p2       image.Image
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

	waiter.Go(func() error {
		image1Data, err := http.Get("https:" + baseCardUrl)
		if err != nil {
			return err
		}
		defer image1Data.Body.Close()

		u.p1, _, err = image.Decode(image1Data.Body)
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

		u.p2, _, err = image.Decode(image2Data.Body)
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

func (u *URPair) stitchPairAndWrite() error {
	baseImage := gg.NewContext(u.p1.Bounds().Dx()+u.p2.Bounds().Dx(), u.p1.Bounds().Dy())

	switch u.Idolized {
	case true:
		switch *u.BaseCard.UrPair.ReverseDisplayIdolized {
		case true:
			{
				u.p1, u.p2 = u.p2, u.p1
			}
		}
	case false:
		switch *u.BaseCard.UrPair.ReverseDisplay {
		case true:
			{
				u.p1, u.p2 = u.p2, u.p1
			}
		}
	}

	baseImage.DrawImage(u.p1, 0, 0)
	baseImage.DrawImage(u.p2, u.p1.Bounds().Dx(), 0)

	u.Image = baseImage.Image()

	return nil
}

func (u *URPair) ProcessImage() (err error) {
	err = u.retrievePair()
	if err != nil {
		return
	}
	err = u.stitchPairAndWrite()
	if err != nil {
		return
	}
	err = u.DoWaifu2x()
	if err != nil {
		return
	}
	return
}
