package usecase

import (
	"fmt"
	"os"
	"strings"

	"github.com/scylladb/go-set/strset"

	"github.com/suzuki-shunsuke/git-rm-branch/internal/domain"
)

func listLocalMergedBranches(branches []string, execIntf domain.ExecIntf) (*strset.Set, error) {
	ret := strset.New()
	for _, branch := range branches {
		out, err := execIntf.CommandCombinedOutput(
			"git", "branch", "--merged", branch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "git branch --merged %s\n%s", branch, out)
			return nil, err
		}
		for _, s := range strings.Split(string(out), "\n") {
			s = strings.TrimSpace(s)
			switch {
			case s == "":
				continue
			case len(s) == 1:
				ret.Add(s)
				continue
			case s[:2] == "* ":
				continue
			default:
				ret.Add(s)
			}
		}
	}
	return ret, nil
}

// remove local branches
func rmLocalBranch(cfg *domain.Cfg, isDryRun, isQuiet bool, execIntf domain.ExecIntf) error {
	// list branches
	removedBranchCandidates, err := listLocalMergedBranches(
		cfg.Local["merged"], execIntf)
	if err != nil {
		return err
	}
	// exclude protected branches
	gitBranchCmdArgs := getRemoveBranchCmdArgs(
		[]string{"branch", "-d"}, removedBranchCandidates, cfg.Local["protected"])
	if len(gitBranchCmdArgs) == 2 {
		if !isQuiet {
			fmt.Println("No local branch is removed.")
		}
		return nil
	}
	// remove branches
	if isDryRun {
		fmt.Printf("[Dry Run] git %s\n", strings.Join(gitBranchCmdArgs, " "))
		return nil
	}
	return execRemoveBranches(gitBranchCmdArgs, isQuiet, execIntf)
}
