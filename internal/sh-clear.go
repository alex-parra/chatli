package internal

import (
	"os"
	"os/exec"
	"runtime"
)

func shClear() {
	sh := func(name string, arg ...string) {
		cmd := exec.Command(name, arg...)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	switch runtime.GOOS {
	case "windows":
		sh("cmd", "/c", "cls")
	default: // darwin, linux
		sh("clear")
	}
}
