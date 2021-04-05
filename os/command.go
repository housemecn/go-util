package os

import (
	"os"
	"os/exec"
)

// RunBashCommand run bash command
func RunBashCommand(cmd string) error {
	command := exec.Command("bash", "-c", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
