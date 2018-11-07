package infra

import (
	"os"
)

type (
	// OS implements domain.OSIntf .
	OS struct{}
)

// Getwd wraps os.Getwd .
func (osImpl *OS) Getwd() (string, error) {
	return os.Getwd()
}

// Stat wraps os.Stat .
func (osImpl *OS) Stat(dst string) (os.FileInfo, error) {
	return os.Stat(dst)
}
