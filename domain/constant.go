package domain

const (
	// ConfigFileName is a configuration file name.
	ConfigFileName = ".git-rm-branch.yml"
	// InitialConfig is a configuration created by init command.
	InitialConfig = `---
local:
  protected:
  - master
  merged:
  - origin/master
remote:
  origin:
    protected:
    - master
    merged:
    - origin/master`
)
