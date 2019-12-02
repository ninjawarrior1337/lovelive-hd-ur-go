package cardhandlers

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Waifu2xAble struct {
	FileBaseName string
}

func (w *Waifu2xAble) OutputDir() string {
	return path.Join("cardsOut", w.FileBaseName)
}

func (w *Waifu2xAble) InputDir() string {
	return path.Join("cardsIn", w.FileBaseName)
}

func (w *Waifu2xAble) DoWaifu2x() (err error) {
	_, err = os.Stat(w.OutputDir())
	if err == nil {
		return
	}
	waifu2xCmd := exec.Command("/usr/bin/waifu2x-converter-cpp", "-i", w.InputDir(), "-o", w.OutputDir(), "--noise-level 3", "--scale-ratio 2")
	waifu2xOut, _ := waifu2xCmd.Output()
	fmt.Println(string(waifu2xOut))
	_, err = os.Stat(w.OutputDir())
	if err != nil {
		return err
	}
	return
}
