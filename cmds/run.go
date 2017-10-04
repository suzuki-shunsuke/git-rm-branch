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

func rmLocalBranch(cfg models.Cfg) error {
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
		fmt.Println("no local branch is removed")
		return nil
	}
	// remove branches
	//   git branch -d %
	fmt.Printf("git %s\n", strings.Join(gitBranchCmdArgs, " "))
	_, err := exec.Command("git", gitBranchCmdArgs...).Output()
	return err
}

func rmRemoteBranch(cfg models.Cfg) error {
	// remove remote branches
	//   git branch -r --merged master | sed -e "s/ *\(.*\) */\1/" | grep "^origin/" | sed -e "s/origin\///" | grep -v "HEAD -> " |  grep -vE $EXCLUDED_BRANCHS | xargs -I% git push --delete origin %
	// list branches
	//   git branch -r --merged master
	// origin/HEAD -> origin/master
	// origin/master
	for remote, branches := range cfg.Remote {
		removedBranchCandidates := map[string]string{}
		for _, branch := range branches {
			out, err := exec.Command("git", "branch", "-r", "--merged", branch).Output()
			if err != nil {
				return err
			}
			for _, s := range strings.Split(string(out), "\n") {
				s = strings.TrimSpace(s)
				if strings.HasPrefix(s, fmt.Sprintf("%s/", remote)) {
					s = s[len(remote)+1:]
					if !strings.Contains(s, "->") {
						removedBranchCandidates[s] = ""
					}
				}
			}
		}
		// exclude protected branches
		gitPushCmdArgs := []string{"push", "--delete", remote}
		for branch, _ := range removedBranchCandidates {
			isRemoved := true
			for _, b := range cfg.Remote[remote] {
				if branch == b {
					isRemoved = false
					break
				}
			}
			if isRemoved {
				gitPushCmdArgs = append(gitPushCmdArgs, branch)
			}
		}
		if len(gitPushCmdArgs) == 3 {
			fmt.Println("no remote branch is removed")
			continue
		}
		// remove branches
		fmt.Printf("git %s\n", strings.Join(gitPushCmdArgs, " "))
		_, err := exec.Command("git", gitPushCmdArgs...).Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func getCfg(wd string) (*models.Cfg, error) {
	rootDir, err := services.FindRoot(wd)
	if err != nil {
		return nil, err
	}
	dest := filepath.Join(rootDir, services.CONFIG_FILENAME)
	if _, err = os.Stat(dest); err != nil {
		return nil, errors.New("configuration file is not found")
	}
	// read configuration file
	buf, err := ioutil.ReadFile(dest)
	if err != nil {
		return nil, err
	}
	// parse configuration file
	var cfg models.Cfg
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}

func runCore() error {
	// remove branches
	// find configuration file
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	cfg, err := getCfg(wd)
	if err != nil {
		return err
	}
	err = rmLocalBranch(*cfg)
	if err != nil {
		return err
	}
	return rmRemoteBranch(*cfg)
}

func Run(c *cli.Context) error {
	err := runCore()
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return err
}
