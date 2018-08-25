package cmds

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/infra"
	"github.com/suzuki-shunsuke/git-rm-branch/usecase"
)

var (
	// InitCommand is an init command.
	InitCommand = cli.Command{
		Name:   "init",
		Usage:  "create a configuration file",
		Action: initCommand,
	}
)

func initCommand(c *cli.Context) error {
	return usecase.InitCommand(&infra.OS{}, &infra.IOUtil{})
}
