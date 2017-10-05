package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/cmds"
)

func main() {
	app := cli.NewApp()
	app.Usage = "remove git's merged branches"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "create a configuration file",
			Action: cmds.Init,
		},
		{
			Name:   "run",
			Usage:  "remove merged branches",
			Action: cmds.Run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Usage: "The path of the configuration file",
				},
				cli.BoolFlag{
					Name:  "dry-run",
					Usage: "don't remove branches but print commands to remove branches",
				},
				cli.BoolFlag{
					Name:  "quiet",
					Usage: "don't print commands",
				},
				cli.BoolFlag{
					Name:  "local",
					Usage: "remove only local branches",
				}},
		},
	}
	app.Run(os.Args)
}
