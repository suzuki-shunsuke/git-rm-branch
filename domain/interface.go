package domain

import (
	"os"
	"time"
)

type (
	// OSIntf is an interface of os package.
	OSIntf interface {
		Getwd() (string, error)
		Stat(string) (os.FileInfo, error)
	}

	// IOUtilIntf is an interface of io/ioutil package.
	IOUtilIntf interface {
		WriteFile(filename string, data []byte, perm os.FileMode) error
		ReadFile(filename string) ([]byte, error)
	}

	// ExecIntf is an interface of os/exec package.
	ExecIntf interface {
		CommandCombinedOutput(name string, arg ...string) ([]byte, error)
	}

	// OSFileInfo is an interface of os.FileInfo .
	OSFileInfo interface {
		Name() string
		Size() int64
		Mode() os.FileMode
		ModTime() time.Time
		IsDir() bool
		Sys() interface{}
	}
)
