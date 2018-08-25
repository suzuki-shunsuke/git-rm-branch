package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/cmds"
	"github.com/suzuki-shunsuke/git-rm-branch/domain"
)

func main() {
	app := cli.NewApp()
	app.Usage = "remove git's merged branches"
	app.Version = domain.Version

	app.Commands = []cli.Command{
		cmds.InitCommand,
		cmds.RunCommand,
	}
	app.Run(os.Args)
}
