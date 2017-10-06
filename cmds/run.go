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

func rmLocalBranch(cfg models.Cfg, isDryRun, isQuiet bool) error {
	// remove local branches
	removedBranchCandidates := map[string]string{}
	// list branches
	protectedBranches := cfg.Local["protected"]
	mergedBranches := cfg.Local["merged"]
	for _, branch := range mergedBranches {
		out, err := exec.Command("git", "branch", "--merged", branch).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "git branch --merged %s\n%s", branch, out)
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
		for _, b := range protectedBranches {
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
		if !isQuiet {
			fmt.Println("No local branch is removed.")
		}
		return nil
	}
	// remove branches
	if isDryRun {
		fmt.Printf("[Dry Run] git %s\n", strings.Join(gitBranchCmdArgs, " "))
	} else {
		if !isQuiet {
			fmt.Printf("git %s\n", strings.Join(gitBranchCmdArgs, " "))
		}
		out, err := exec.Command("git", gitBranchCmdArgs...).CombinedOutput()
		if err != nil {
			if isQuiet {
				fmt.Fprintf(os.Stderr, "git %s\n%s", strings.Join(gitBranchCmdArgs, " "), out)
			} else {
				fmt.Fprint(os.Stderr, string(out))
			}
			return err
		}
	}
	return nil
}

func rmRemoteBranch(cfg models.Cfg, isDryRun, isQuiet bool) error {
	// remove remote branches
	// list branches
	for remote, remoteVal := range cfg.Remote {
		protectedBranches := remoteVal["protected"]
		mergedBranches := remoteVal["merged"]
		removedBranchCandidates := map[string]string{}
		for _, branch := range mergedBranches {
			out, err := exec.Command("git", "branch", "-r", "--merged", branch).CombinedOutput()
			if err != nil {
				fmt.Fprintf(os.Stderr, "git branch -r --merged %s\n%s", branch, out)
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
			for _, b := range protectedBranches {
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
			if !isQuiet {
				fmt.Println("No remote branch is removed.")
			}
			continue
		}
		// remove branches
		if isDryRun {
			fmt.Printf("[Dry Run] git %s\n", strings.Join(gitPushCmdArgs, " "))
		} else {
			if !isQuiet {
				fmt.Printf("git %s\n", strings.Join(gitPushCmdArgs, " "))
			}
			out, err := exec.Command("git", gitPushCmdArgs...).CombinedOutput()
			if err != nil {
				if isQuiet {
					fmt.Fprintf(os.Stderr, "git %s\n%s", strings.Join(gitPushCmdArgs, " "), out)
				} else {
					fmt.Fprint(os.Stderr, string(out))
				}
				return err
			}
		}
	}
	return nil
}

func getCfg(cfgFilePath string) (*models.Cfg, error) {
	if _, err := os.Stat(cfgFilePath); err != nil {
		return nil, errors.New("Configuration file is not found.")
	}
	// read configuration file
	buf, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return nil, err
	}
	// parse configuration file
	var cfg models.Cfg
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}

func findCfg(wd string) (string, error) {
	rootDir, err := services.FindRoot(wd)
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, services.CONFIG_FILENAME), nil
}

func runCore(isDryRun, isQuiet, isOnlyLocal bool, cfgFilePath string) error {
	// remove branches
	// find configuration file
	if isDryRun && isQuiet {
		return errors.New("Both --dry-run and --quiet options must not be used.")
	}
	if cfgFilePath == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		cfgFilePath, err = findCfg(wd)
		if err != nil {
			return err
		}
	}
	cfg, err := getCfg(cfgFilePath)
	if err != nil {
		return err
	}
	err = rmLocalBranch(*cfg, isDryRun, isQuiet)
	if err != nil {
		return err
	}
	if isOnlyLocal {
		return nil
	}
	return rmRemoteBranch(*cfg, isDryRun, isQuiet)
}

func Run(c *cli.Context) error {
	isDryRun := c.Bool("dry-run")
	isQuiet := c.Bool("quiet")
	isOnlyLocal := c.Bool("local")
	cfgFilePath := c.String("config")
	err := runCore(isDryRun, isQuiet, isOnlyLocal, cfgFilePath)
	if err != nil {
		return cli.NewExitError(err, services.GetStatusCode(err))
	}
	return err
}
