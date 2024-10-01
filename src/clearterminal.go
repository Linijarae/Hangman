package Hangman

import (
	"os"
	"os/exec"
	"runtime"
)

//Clear le terminal
func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
