package cmds

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/git-rm-branch/models"
	"github.com/suzuki-shunsuke/git-rm-branch/services"
)

func runCore() error {
	// remove branches
	// remove local branches
	//   git branch --merged master | grep -vE $EXCLUDED_BRANCHS | xargs -I % git branch -d %
	// remove remote branches
	//   git branch -r --merged master | sed -e "s/ *\(.*\) */\1/" | grep "^origin/" | sed -e "s/origin\///" | grep -v "HEAD -> " |  grep -vE $EXCLUDED_BRANCHS | xargs -I% git push --delete origin %
	// find configuration file
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	rootDir, err := services.FindRoot(wd)
	if err != nil {
		return err
	}
	dest := filepath.Join(rootDir, services.CONFIG_FILENAME)
	if _, err = os.Stat(dest); err != nil {
		return errors.New("configuration file is not found")
	}
	// read configuration file
	buf, err := ioutil.ReadFile(dest)
	if err != nil {
		return err
	}
	// parse configuration file
	var cfg models.Cfg
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return err
	}
	// remove local branches
	removedBranchCandidates := map[string]string{}
	// list branches
	//   git branch --merged master
	for _, branch := range cfg.Local {
		out, err := exec.Command("git", "branch", "--merged", branch).Output()
		if err != nil {
			return err
		}
		for _, s := range strings.Split(string(out), "\n") {
			s = strings.TrimSpace(s)
			if s != "" && (len(s) < 2 || s[:2] != "* ") {
				removedBranchCandidates[s] = ""
			}
		}
	}
	// exclude protected branches
	gitBranchCmdArgs := []string{"branch", "-d"}
	for branch, _ := range removedBranchCandidates {
		isRemoved := true
		for _, b := range cfg.Local {
			if branch == b {
				isRemoved = false
				break
			}
		}
		if isRemoved {
			gitBranchCmdArgs = append(gitBranchCmdArgs, branch)
		}
	}
	if len(gitBranchCmdArgs) == 2 {
		fmt.Println("no branch is removed")
		return nil
	}
	// remove branches
	//   git branch -d %
	fmt.Printf("git %s\n", strings.Join(gitBranchCmdArgs, " "))
	_, err = exec.Command("git", gitBranchCmdArgs...).Output()
	return err
}

func Run(c *cli.Context) error {
	err := runCore()
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return err
}
