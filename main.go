package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/cmds"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "create a configuration file",
			Action: cmds.Init,
		},
		{
			Name:   "run",
			Usage:  "remove branches",
			Action: cmds.Run,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "dry-run",
				},
				cli.BoolFlag{
					Name: "quiet",
				}},
		},
	}
	app.Run(os.Args)
}
