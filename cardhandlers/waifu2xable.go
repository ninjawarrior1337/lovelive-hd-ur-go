package cardhandlers

import (
	"fmt"
	"image"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

type Waifu2xAble struct {
	image.Image
	FileBaseName string
}

func (w *Waifu2xAble) OutputPath() string {
	outpath, _ := filepath.Abs(path.Join("./cardsOut", w.FileBaseName))
	return outpath
}

func (w *Waifu2xAble) InputPath() string {
	inpath, _ := filepath.Abs(path.Join("./cardsIn", w.FileBaseName))
	return inpath
}

func (w *Waifu2xAble) DoWaifu2x() (err error) {
	// Check if file already exists

	_, err = os.Stat(w.OutputPath())
	if err == nil {
		return
	}
	//Continue if it doesnt

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
