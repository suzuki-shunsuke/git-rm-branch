package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/scylladb/go-set/strset"
	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/git-rm-branch/internal/domain"
)

func Test_rmRemoteBranch(t *testing.T) {
	assert.Nil(t, rmRemoteBranch(&domain.Cfg{}, true, false, nil))
}

func Test_listRemoteRemovedBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockExecIntf := NewMockExecIntf(ctrl)
	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "-r", "--merged", "master",
	).Return(nil, fmt.Errorf("command error"))
	_, err := listRemoteRemovedBranches("origin", []string{"master"}, mockExecIntf)
	assert.NotNil(t, err)
	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "-r", "--merged", "master",
	).Return([]byte(`origin/devel`), nil)
	s, err := listRemoteRemovedBranches("origin", []string{"master"}, mockExecIntf)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"devel"}, s.List())
}

func Test_findCfg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOSIntf := NewMockOSIntf(ctrl)
	mockOSIntf.EXPECT().Stat("/foo/.git").Return(nil, nil).AnyTimes()
	mockOSIntf.EXPECT().Stat("/.git").Return(nil, fmt.Errorf("file is not found")).AnyTimes()

	s, err := findCfg("/foo", mockOSIntf)
	assert.Nil(t, err)
	exp := "/foo/.git-rm-branch.yml"
	assert.Equal(t, s, exp)

	if _, err := findCfg("/", mockOSIntf); err == nil {
		t.Fatal("file is not found")
	}
}

func Test_getCfg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOSIntf := NewMockOSIntf(ctrl)
	cfgFilePath := "/foo/.git-rm-branch.yml"
	mockOSIntf.EXPECT().Stat(cfgFilePath).Return(
		nil, fmt.Errorf("file is not found"))

	if _, err := getCfg(cfgFilePath, mockOSIntf, nil); err == nil {
		t.Fatal("file is not found")
	}
	mockIOUtilIntf := NewMockIOUtilIntf(ctrl)
	cfgFilePath = "/.git-rm-branch.yml"
	mockOSIntf.EXPECT().Stat(cfgFilePath).Return(nil, nil).AnyTimes()
	mockIOUtilIntf.EXPECT().ReadFile(cfgFilePath).Return(
		nil, fmt.Errorf("failed to read config file"))
	if _, err := getCfg(cfgFilePath, mockOSIntf, mockIOUtilIntf); err == nil {
		t.Fatal("failed to read config file")
	}
	mockIOUtilIntf.EXPECT().ReadFile(cfgFilePath).Return([]byte(domain.InitialConfig), nil)
	if _, err := getCfg(cfgFilePath, mockOSIntf, mockIOUtilIntf); err != nil {
		t.Fatal(err)
	}
}

func TestRunCommand(t *testing.T) {
	assert.NotNil(t, RunCommand(true, true, false, "", nil, nil, nil))
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOSIntf := NewMockOSIntf(ctrl)
	mockOSIntf.EXPECT().Getwd().Return("", fmt.Errorf("failed to get pwd"))
	assert.NotNil(t, RunCommand(true, false, false, "", mockOSIntf, nil, nil))
	mockOSIntf.EXPECT().Getwd().Return("/", nil)
	mockOSIntf.EXPECT().Stat("/.git").Return(nil, fmt.Errorf("file is not found"))
	assert.NotNil(t, RunCommand(true, false, false, "", mockOSIntf, nil, nil))
	mockOSIntf.EXPECT().Stat("/.git-rm-branch.yml").Return(
		nil, fmt.Errorf("file is not found"))
	assert.NotNil(t, RunCommand(
		true, false, false, "/.git-rm-branch.yml", mockOSIntf, nil, nil))
}

func Test_execRemoveBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockExecIntf := NewMockExecIntf(ctrl)
	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "-d", "devel",
	).Return(nil, nil)
	assert.Nil(t, execRemoveBranches([]string{"branch", "-d", "devel"}, false, mockExecIntf))

	mockExecIntf.EXPECT().CommandCombinedOutput(
		"git", "branch", "-d", "devel",
	).Return(nil, fmt.Errorf("failed to remove branches")).AnyTimes()
	assert.NotNil(t, execRemoveBranches([]string{"branch", "-d", "devel"}, false, mockExecIntf))
	assert.NotNil(t, execRemoveBranches([]string{"branch", "-d", "devel"}, true, mockExecIntf))
}

func Test_getRemoveBranchCmdArgs(t *testing.T) {
	arr := getRemoveBranchCmdArgs([]string{"branch", "-d"}, strset.New("master", "devel"), []string{"master"})
	assert.ElementsMatch(t, arr, []string{"branch", "-d", "devel"})
}
