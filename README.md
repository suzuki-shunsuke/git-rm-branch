# git-rm-branch

[![CircleCI](https://circleci.com/gh/suzuki-shunsuke/git-rm-branch.svg?style=svg)](https://circleci.com/gh/suzuki-shunsuke/git-rm-branch)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/git-rm-branch/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/git-rm-branch)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/git-rm-branch)](https://goreportcard.com/report/github.com/suzuki-shunsuke/git-rm-branch)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/git-rm-branch.svg)](https://github.com/suzuki-shunsuke/git-rm-branch)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/git-rm-branch.svg)](https://github.com/suzuki-shunsuke/git-rm-branch/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/git-rm-branch/master/LICENSE)

cli tool to remove merged branches

## Install

```
$ go get suzuki-shunsuke/git-rm-branch
```

## Usage

```
$ git-rm-branch init
$ git-rm-branch run [--local] [--dry-run] [--quiet] [--config <config path>]
```

### init

```
$ git-rm-branch help init
NAME:
   git-rm-branch init - create a configuration file

USAGE:
   git-rm-branch init [arguments...]
```

### run

```
$ git-rm-branch help run
NAME:
   git-rm-branch run - remove merged branches

USAGE:
   git-rm-branch run [command options] [arguments...]

OPTIONS:
   --config value  The path of the configuration file
   --dry-run       don't remove branches but print commands to remove branches
   --quiet         don't print commands
   --local         remove only local branches
```

## Configuration file location

If the `--config` option is not used,
this tool assumes that the configuration file `.git-rm-branch.yml` is in the root directory of the git repository.

## The example of the configuration file

```yaml
local:
  protected:
    - master
    - develop
  merged:
    - upstream/master
remote:
  origin:
    protected:
      - master
    merged:
      - upstream/master
  upstream:
    protected:
      - master
    merged:
      - master
```

## Change Log

See [CHANGELOG.md](CHANGELOG.md).

## License

[MIT](LICENSE)
