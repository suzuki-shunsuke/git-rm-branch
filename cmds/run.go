package cmds

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/infra"
	"github.com/suzuki-shunsuke/git-rm-branch/usecase"
)

var (
	// RunCommand is a run command.
	RunCommand = cli.Command{
		Name:   "run",
		Usage:  "remove merged branches",
		Action: runCommand,
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
	}
)

func runCommand(c *cli.Context) error {
	isDryRun := c.Bool("dry-run")
	isQuiet := c.Bool("quiet")
	isOnlyLocal := c.Bool("local")
	cfgFilePath := c.String("config")
	err := usecase.RunCommand(
		isDryRun, isQuiet, isOnlyLocal, cfgFilePath,
		&infra.OS{}, &infra.IOUtil{}, &infra.Exec{})
	if err != nil {
		return cli.NewExitError(err, infra.GetStatusCode(err))
	}
	return err
}
