package services

import (
	"errors"
	"os"
	"path/filepath"
)

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
