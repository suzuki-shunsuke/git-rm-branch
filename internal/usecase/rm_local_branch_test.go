package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/git-rm-branch/internal/domain"
)

func Test_listLocalMergedBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockExecIntf := NewMockExecIntf(ctrl)
	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "--merged", "origin/master",
	).Return([]byte(`devel`), nil)
	s, err := listLocalMergedBranches([]string{"origin/master"}, mockExecIntf)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"devel"}, s.List())

	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "--merged", "origin/master",
	).Return([]byte(`
* hoge`), nil)
	s, err = listLocalMergedBranches([]string{"origin/master"}, mockExecIntf)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{}, s.List())

	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "--merged", "origin/master",
	).Return([]byte(`a`), nil)
	s, err = listLocalMergedBranches([]string{"origin/master"}, mockExecIntf)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"a"}, s.List())
}

func Test_rmLocalBranch(t *testing.T) {
	assert.Nil(t, rmLocalBranch(&domain.Cfg{}, true, false, nil))
	cfg := &domain.Cfg{
		Local: map[string][]string{
			"merged": {"origin/master"},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockExecIntf := NewMockExecIntf(ctrl)
	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "--merged", "origin/master",
	).Return(nil, fmt.Errorf("command error"))
	assert.NotNil(t, rmLocalBranch(cfg, true, false, mockExecIntf))
}
