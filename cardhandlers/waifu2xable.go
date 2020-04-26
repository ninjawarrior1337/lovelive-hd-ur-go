package cardhandlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

type Waifu2xAble struct {
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
	var waifu2xCmd *exec.Cmd

	prepareWaifu2xCommand(waifu2xCmd, w.InputPath(), w.OutputPath())

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