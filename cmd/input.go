package cmd

import "os"

func IsInputFromPipe() bool {
	info, _ := os.Stdin.Stat()
	return info.Mode()&os.ModeCharDevice != os.ModeCharDevice
}
