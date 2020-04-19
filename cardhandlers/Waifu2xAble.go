package cardhandlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"syscall"
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
	if runtime.GOOS == "windows" {
		var waifu2xArgs = []string{"-i", w.InputPath(), "-o", w.OutputPath(), "--noise_level", "3", "--scale_ratio", "2"}
		waifu2xPath, _ := filepath.Abs("./waifu2xwin/waifu2x-caffe-cui.exe")
		waifu2xCmd = exec.Command(waifu2xPath, waifu2xArgs...)
		waifu2xCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		var waifu2xArgs = []string{"-i", w.InputPath(), "-o", w.OutputPath(), "--noise-level", "3", "--scale-ratio", "2"}
		waifu2xCmd = exec.Command("/usr/bin/waifu2x-converter-cpp", waifu2xArgs...)
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
