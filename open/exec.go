// +build linux

package open

import (
	"os/exec"
)

func open(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}
