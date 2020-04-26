package cardhandlers

import (
	"os/exec"
)

func prepareWaifu2xCommand(cmd *exec.Cmd, inputPath, outputPath string) {
	var waifu2xArgs = []string{"-i", inputPath, "-o", outputPath, "--noise-level", "3", "--scale-ratio", "2"}
	cmd = exec.Command("/usr/bin/waifu2x-converter-cpp", waifu2xArgs...)
}
