package cardhandlers

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

func prepareWaifu2xCommand(inputPath, outputPath string) *exec.Cmd {
	var waifu2xArgs = []string{"-i", inputPath, "-o", outputPath, "--noise_level", "3", "--scale_ratio", "2"}
	waifu2xPath, _ := filepath.Abs("./waifu2xwin/waifu2x-caffe-cui.exe")
	cmd := exec.Command(waifu2xPath, waifu2xArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
