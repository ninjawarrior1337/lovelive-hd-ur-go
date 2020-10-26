package cardhandlers

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

type Waifu2xAble struct {
	image.Image
}

func (w *Waifu2xAble) Hash() string {
	imgHash, _ := goimagehash.DifferenceHash(w.Image)
	return imgHash.ToString()
}

func (w *Waifu2xAble) HasAlpha() bool {
	im := w.Image
	if op, ok := im.(interface{ Opaque() bool }); ok {
		return !op.Opaque()
	}

	for y := im.Bounds().Min.Y; y < im.Bounds().Max.Y; y++ {
		for x := im.Bounds().Max.X; x < im.Bounds().Max.X; x++ {
			if _, _, _, a := im.At(x, y).RGBA(); a != 0xffff {
				return true
			}
		}
	}
	return false
}

func (w *Waifu2xAble) Extension() string {
	if w.HasAlpha() {
		return ".png"
	} else {
		return ".jpg"
	}
}

func (w *Waifu2xAble) OutputPath() string {
	outpath, _ := filepath.Abs(path.Join("./imgOut", w.Hash()))
	return outpath + w.Extension()
}

func (w *Waifu2xAble) InputPath() string {
	inpath, _ := filepath.Abs(path.Join("./imgIn", w.Hash()))
	return inpath + w.Extension()
}

func (w *Waifu2xAble) DoWaifu2x() (err error) {
	// Check if file already exists

	_, err = os.Stat(w.OutputPath())
	if err == nil {
		return
	}
	//Continue if it doesnt

	//Write Card To Input Folder
	f, err := os.OpenFile(w.InputPath(), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if w.Extension() == ".png" {
		png.Encode(f, w)
	}
	if w.Extension() == ".jpg" {
		jpeg.Encode(f, w, &jpeg.Options{Quality: 90})
	}

	// waifu2xCmd is defined by if the built file is for windows or linux.
	var waifu2xCmd *exec.Cmd

	waifu2xCmd = prepareWaifu2xCommand(w.InputPath(), w.OutputPath())

	if waifu2xCmd == nil {
		return fmt.Errorf("what the frick just happened, I can only imageine that this is running on macOS")
	}

	log.Println("Conversion Command: " + waifu2xCmd.String())
	waifu2xOut, err := waifu2xCmd.Output()
	if err != nil {
		switch runtime.GOOS {
		case "windows":
			log.Println(waifu2xOut)
		default:
			log.Println(string(waifu2xOut))
		}
		return err
	}
	fmt.Println(string(waifu2xOut))
	_, err = os.Stat(w.OutputPath())
	if err != nil {
		return err
	}
	return
}
