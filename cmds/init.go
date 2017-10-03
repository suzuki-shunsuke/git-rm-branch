package cmds

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/assets"
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

func core(wd string) error {
	rootDir, err := FindRoot(wd)
	if err != nil {
		return err
	}
	dest := filepath.Join(rootDir, CONFIG_FILENAME)
	if _, err = os.Stat(dest); err == nil {
		return nil
	}
	// create .git-rm-branch
	data, err := assets.Asset("data/git-rm-branch.yml")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, data, os.ModePerm)
}

func Init(c *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	return core(wd)
}
