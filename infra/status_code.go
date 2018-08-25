package infra

import (
	"os/exec"
	"syscall"
)

// GetStatusCode gets a status code from an error object.
func GetStatusCode(err error) int {
	if e2, ok := err.(*exec.ExitError); ok {
		if s, ok := e2.Sys().(syscall.WaitStatus); ok {
			return s.ExitStatus()
		}
	}
	return 1
}
