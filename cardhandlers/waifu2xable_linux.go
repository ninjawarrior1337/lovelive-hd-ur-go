package cardhandlers

import (
	"os/exec"
)

func prepareWaifu2xCommand(inputPath, outputPath string) *exec.Cmd {
	var waifu2xArgs = []string{"-i", inputPath, "-o", outputPath, "-n", "3", "-s", "2"}
	return exec.Command("/usr/bin/waifu2x-ncnn-vulkan", waifu2xArgs...)
}
