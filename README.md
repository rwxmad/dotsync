# dotsync

### CLI for managing and syncing your dotfiles

### Commands

| Command      | Flags | Arguments | Description |
| ------------ | ----- | --------- | --------- |
| dotsync init | - | - |Initialize a directory for dot files and create configuration file with configs list|
| dotsync add | - | file |Adds the specified file to the dotsync config|
| dotsync store | - | - |Move configuration files to sync folder and create symlinks|


### Config
You need to create a file in .yaml format or generate it using the `dotsync init` command.
Paths must be specified from the home directory, you do not need to register the home directory.
The config file should have the following structure:

```yaml 
file_name: absolute/path/to/file
# example
.vimrc: .vimrc
.tmux.conf: .tmux.conf
.editorconfig: code/.editorconfig
```
