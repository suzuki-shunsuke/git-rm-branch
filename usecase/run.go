package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/scylladb/go-set/strset"

	"github.com/suzuki-shunsuke/git-rm-branch/domain"
)

// RunCommand removes merged branches.
func RunCommand(isDryRun, isQuiet, isOnlyLocal bool, cfgFilePath string, osIntf domain.OSIntf, ioutilIntf domain.IOUtilIntf, execIntf domain.ExecIntf) error {
	if isDryRun && isQuiet {
		return fmt.Errorf("Both --dry-run and --quiet options must not be used.")
	}
	if cfgFilePath == "" {
		wd, err := osIntf.Getwd()
		if err != nil {
			return err
		}
		cfgFilePath, err = findCfg(wd, osIntf)
		if err != nil {
			return err
		}
	}

	cfg, err := getCfg(cfgFilePath, osIntf, ioutilIntf)
	if err != nil {
		return err
	}
	err = rmLocalBranch(cfg, isDryRun, isQuiet, execIntf)
	if err != nil {
		return err
	}
	if isOnlyLocal {
		return nil
	}
	return rmRemoteBranch(cfg, isDryRun, isQuiet, execIntf)
}

func listRemoteRemovedBranches(remote string, branches []string, execIntf domain.ExecIntf) (*strset.Set, error) {
	candidates := strset.New()
	for _, branch := range branches {
		out, err := execIntf.CommandCombinedOutput("git", "branch", "-r", "--merged", branch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "git branch -r --merged %s\n%s", branch, out)
			return nil, err
		}
		for _, s := range strings.Split(string(out), "\n") {
			s = strings.TrimSpace(s)
			if strings.HasPrefix(s, fmt.Sprintf("%s/", remote)) {
				s = s[len(remote)+1:]
				if !strings.Contains(s, "->") {
					candidates.Add(s)
				}
			}
		}
	}
	return candidates, nil
}

func rmRemoteBranch(cfg *domain.Cfg, isDryRun, isQuiet bool, execIntf domain.ExecIntf) error {
	// remove remote branches
	// list branches
	for remote, remoteVal := range cfg.Remote {
		removedBranchCandidates, err := listRemoteRemovedBranches(remote, remoteVal["merged"], execIntf)
		if err != nil {
			return err
		}
		// exclude protected branches
		gitPushCmdArgs := getRemoveBranchCmdArgs(
			[]string{"push", "--delete", remote},
			removedBranchCandidates, remoteVal["protected"])
		if len(gitPushCmdArgs) == 3 {
			if !isQuiet {
				fmt.Println("No remote branch is removed.")
			}
			continue
		}
		// remove branches
		if isDryRun {
			fmt.Printf("[Dry Run] git %s\n", strings.Join(gitPushCmdArgs, " "))
			return nil
		}
		return execRemoveBranches(gitPushCmdArgs, isQuiet, execIntf)
	}
	return nil
}

func findCfg(wd string, osIntf domain.OSIntf) (string, error) {
	rootDir, err := findRoot(wd, osIntf)
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDir, domain.ConfigFileName), nil
}

func getCfg(cfgFilePath string, osIntf domain.OSIntf, ioutilIntf domain.IOUtilIntf) (*domain.Cfg, error) {
	if _, err := osIntf.Stat(cfgFilePath); err != nil {
		return nil, fmt.Errorf("Configuration file is not found.")
	}
	// read configuration file
	buf, err := ioutilIntf.ReadFile(cfgFilePath)
	if err != nil {
		return nil, err
	}
	// parse configuration file
	var cfg domain.Cfg
	err = yaml.Unmarshal(buf, &cfg)
	return &cfg, err
}

func execRemoveBranches(cmdArgs []string, isQuiet bool, execIntf domain.ExecIntf) error {
	if !isQuiet {
		fmt.Printf("git %s\n", strings.Join(cmdArgs, " "))
	}
	out, err := execIntf.CommandCombinedOutput("git", cmdArgs...)
	if err != nil {
		if isQuiet {
			fmt.Fprintf(os.Stderr, "git %s\n%s", strings.Join(cmdArgs, " "), out)
		} else {
			fmt.Fprint(os.Stderr, string(out))
		}
		return err
	}
	return nil
}

func getRemoveBranchCmdArgs(cmdArgs []string, candidates *strset.Set, protectedBranches []string) []string {
	candidates.Each(func(branch string) bool {
		isRemoved := true
		// if branch is not included in protected branches,
		// branch is added to cmdArgs
		for _, b := range protectedBranches {
			if branch == b {
				isRemoved = false
				break
			}
		}
		if isRemoved {
			cmdArgs = append(cmdArgs, branch)
		}
		return true
	})
	return cmdArgs
}
