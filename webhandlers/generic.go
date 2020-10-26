package webhandlers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ninjawarrior1337/lovelive-hd-ur-go/cardhandlers"
	"image"
)

func ImageFromForm(ctx *fiber.Ctx) (image.Image, error) {
	fH, err := ctx.FormFile("image")
	if err != nil {
		return nil, err
	}
	f, err := fH.Open()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ImageFromB64(ctx *fiber.Ctx) (image.Image, error) {
	b64Str := ctx.Body()
	b64Buf := bytes.Buffer{}
	b64Buf.Write(b64Str)
	r := base64.NewDecoder(base64.RawURLEncoding, &b64Buf)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func GenericHandler(ctx *fiber.Ctx) error {
	var im image.Image

	if img, err := ImageFromForm(ctx); err == nil {
		im = img
	} else if img, err := ImageFromB64(ctx); err == nil {
		im = img
	}

	if im == nil {
		return errors.New("please send an image as either base64 or as formdata")
	}

	w := cardhandlers.Waifu2xAble{Image: im}

	err := w.DoWaifu2x()
	if err != nil {
		return err
	}

	ctx.Append("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", w.Hash()+w.Extension()))
	ctx.SendFile(w.OutputPath())
	return nil
}
