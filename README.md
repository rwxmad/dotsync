# dotsync

### CLI for managing and syncing your dotfiles

[![Go Report Card](https://goreportcard.com/badge/github.com/rwxmad/dotsync)](https://goreportcard.com/report/github.com/rwxmad/dotsync)

### Usage
```bash
dotsync [command] [args] [flag] [value]
```

### Commands

| Command      | Flags | Value for flag | Arguments | Description |
| ------------ | ----- | --------- | --------- | ------------ |
| `init` | - |                |-| Initialize a directory for dotfiles and create configuration file with configs list |
| `add` | - |  |Paths to files| Adds the specified file to the dotsync config |
| `store` | - |  |-| Move configuration files to sync folder and create symlinks |
| `config` | [--git, -g] - show .git/config | - |-| Show and edit .dotsync.yaml, .git/config files |


### Config
You need to create a file in .yaml format or generate it using the `dotsync init` command.
Paths must be specified from the home directory, you do not need to register the home directory.
The config file should have the following structure:

```yaml 
file_name: path/from/home/dir
# example
.vimrc: .vimrc
.tmux.conf: .tmux.conf
.editorconfig: code/.editorconfig
```

