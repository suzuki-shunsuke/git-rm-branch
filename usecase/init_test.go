package usecase

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/suzuki-shunsuke/git-rm-branch/domain"
)

func TestInitCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("failed to get pwd", func(t *testing.T) {
		mockOSIntf := NewMockOSIntf(ctrl)
		mockOSIntf.EXPECT().Getwd().Return("", fmt.Errorf("failed to get pwd"))
		assert.NotNil(t, InitCommand(mockOSIntf, nil))
	})
	t.Run("git repository is not found", func(t *testing.T) {
		mockOSIntf := NewMockOSIntf(ctrl)
		mockOSIntf.EXPECT().Getwd().Return("/", nil)
		mockOSIntf.EXPECT().Stat("/.git").Return(nil, fmt.Errorf("file is not found"))
		assert.NotNil(t, InitCommand(mockOSIntf, nil))
	})
	t.Run(".git-rm-branch.yml has already existed", func(t *testing.T) {
		mockOSIntf := NewMockOSIntf(ctrl)
		mockOSIntf.EXPECT().Getwd().Return("/", nil)
		mockOSIntf.EXPECT().Stat("/.git").Return(nil, nil)
		mockOSIntf.EXPECT().Stat("/.git-rm-branch.yml").Return(nil, nil)
		assert.Nil(t, InitCommand(mockOSIntf, nil))
	})
	t.Run("create a .git-rm-branch.yml", func(t *testing.T) {
		mockOSIntf := NewMockOSIntf(ctrl)
		mockOSIntf.EXPECT().Getwd().Return("/", nil)
		mockOSIntf.EXPECT().Stat("/.git").Return(nil, nil)
		mockOSIntf.EXPECT().Stat("/.git-rm-branch.yml").Return(nil, fmt.Errorf(
			"/.git-rm-branch.yml is not found"))
		mockIOUtilIntf := NewMockIOUtilIntf(ctrl)
		mockIOUtilIntf.EXPECT().WriteFile(
			"/.git-rm-branch.yml", []byte(domain.InitialConfig),
			os.ModePerm).Return(nil)
		assert.Nil(t, InitCommand(mockOSIntf, mockIOUtilIntf))
	})
}

func Test_findRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOSFileInfo := NewMockOSFileInfo(ctrl)

	mockOSIntf := NewMockOSIntf(ctrl)
	mockOSIntf.EXPECT().Stat("/foo/.git").Return(mockOSFileInfo, nil).AnyTimes()
	mockOSIntf.EXPECT().Stat("/.git").Return(nil, fmt.Errorf("file is not found")).AnyTimes()
	mockOSIntf.EXPECT().Stat("/foo/bar/.git").Return(nil, fmt.Errorf("file is not found")).AnyTimes()

	s, err := findRoot("/foo", mockOSIntf)
	assert.Nil(t, err)
	exp := "/foo"
	assert.Equal(t, s, exp)

	s, err = findRoot("/foo/bar", mockOSIntf)
	assert.Nil(t, err)
	assert.Equal(t, s, exp)

	if _, err := findRoot("", mockOSIntf); err == nil {
		t.Fatal(`git repository is not found`)
	}
	if _, err := findRoot("foo/bar", mockOSIntf); err == nil {
		t.Fatal(`file path must be absolute`)
	}
	if _, err := findRoot("/", mockOSIntf); err == nil {
		t.Fatal(`file is not found`)
	}
}
