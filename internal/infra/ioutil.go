package infra

import (
	"io/ioutil"
	"os"
)

type (
	// IOUtil implements domain.IOUtilIntf .
	IOUtil struct{}
)

// ReadFile wraps io/ioutil.ReadFile .
func (ioutilImpl *IOUtil) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// WriteFile wraps io/ioutil.WriteFile .
func (ioutilImpl *IOUtil) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
