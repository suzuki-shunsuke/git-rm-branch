package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/suzuki-shunsuke/git-rm-branch/domain"
)

// InitCommand creates a configuration file if it doesn't exist.
func InitCommand(osIntf domain.OSIntf, ioutilIntf domain.IOUtilIntf) error {
	wd, err := osIntf.Getwd()
	if err != nil {
		return err
	}
	rootDir, err := findRoot(wd, osIntf)
	if err != nil {
		return err
	}
	dst := filepath.Join(rootDir, domain.ConfigFileName)
	if _, err = osIntf.Stat(dst); err == nil {
		return nil
	}
	// create .git-rm-branch
	return ioutilIntf.WriteFile(dst, []byte(domain.InitialConfig), os.ModePerm)
}

func findRoot(startDir string, osIntf domain.OSIntf) (string, error) {
	// find .git
	dir := startDir
	// "/" "" ".."
	for {
		if dir == "" {
			return "", fmt.Errorf("git repository is not found")
		}
		if !filepath.IsAbs(dir) {
			return "", fmt.Errorf("file path must be absolute")
		}
		if _, err := osIntf.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		if dir == "/" {
			return "", fmt.Errorf("git repository is not found")
		}
		dir = filepath.Dir(dir)
	}
}
