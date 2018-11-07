package infra

import (
	"os/exec"
)

type (
	// Exec implements domain.ExecIntf.
	Exec struct{}
)

// CommandCombinedOutput executes a command.
func (execImpl *Exec) CommandCombinedOutput(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).CombinedOutput()
}
