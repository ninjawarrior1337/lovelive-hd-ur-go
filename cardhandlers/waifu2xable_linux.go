package cardhandlers

import (
	"os/exec"
)

func prepareWaifu2xCommand(inputPath, outputPath string) *exec.Cmd {
	var waifu2xArgs = []string{"-i", inputPath, "-o", outputPath, "--noise-level", "3", "--scale-ratio", "2"}
	return exec.Command("/usr/bin/waifu2x-converter-cpp", waifu2xArgs...)
}
