package services

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var CONFIG_FILENAME = ".git-rm-branch.yml"

func FindRoot(startDir string) (string, error) {
	// find .git
	dir := startDir
	// "/" "" ".."
	for {
		if dir == "" {
			return "", errors.New("git repository is not found")
		}
		if !filepath.IsAbs(dir) {
			return "", errors.New("file path must be absolute")
		}
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		dir = filepath.Dir(dir)
	}
	return "", errors.New("git repository is not found")
}

func GetStatusCode(err error) int {
	if e2, ok := err.(*exec.ExitError); ok {
		if s, ok := e2.Sys().(syscall.WaitStatus); ok {
			return s.ExitStatus()
		}
	}
	return 1
}
