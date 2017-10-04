package cmds

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/assets"
	"github.com/suzuki-shunsuke/git-rm-branch/services"
)

func core(wd string) error {
	rootDir, err := services.FindRoot(wd)
	if err != nil {
		return err
	}
	dest := filepath.Join(rootDir, services.CONFIG_FILENAME)
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
