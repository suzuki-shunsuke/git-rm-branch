package handler

import (
	"os"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/internal/domain"
)

// Main is called at cmd/git-rm-branch/main.go .
func Main() {
	app := cli.NewApp()
	app.Usage = "remove git's merged branches"
	app.Version = domain.Version

	app.Commands = []cli.Command{
		InitCommand,
		RunCommand,
	}
	app.Run(os.Args)
}
