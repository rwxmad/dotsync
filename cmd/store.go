package cmd

import (
	"fmt"
  "errors"
  "os"
  "log"

	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store, symlink and sync dotfiles",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    Store()
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}

var (
  pathsMap = make(map[string]string)
  ErrConfigEmpty = errors.New("Config file is empty")
)

func Store() {
  getPaths()
}

func getPaths() {
  InitConfig(&cfgFile)
  err := v.Unmarshal(&pathsMap)
  if err != nil {
    log.Fatal("Error while reading configuration file")
  }
  if len(pathsMap) == 0 {
    fmt.Println(ErrConfigEmpty)
  }else {
    storeFiles()
  }
}

func storeFiles() {
  for name, path := range pathsMap{
    symlinkFiles(name, path)
  }
}

func symlinkFiles(name, path string) {
  homePath, err := os.UserHomeDir()
  if err != nil {
    fmt.Println(err)
  }

  oldPath := homePath + "/" + path
  newPath := dirPath + name

  moveFile(oldPath, newPath)
  os.Symlink(newPath, oldPath)

  fmt.Printf("# [Symlink to the %s at the path: %s] ===> was created\n", name, path)
}

func moveFile(oldLocation, newLocation string) {
  err := os.Rename(oldLocation, newLocation)
  if err != nil {
    fmt.Println(err)
  }
}

