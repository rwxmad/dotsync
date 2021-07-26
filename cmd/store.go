package cmd

import (
	"fmt"
  "errors"
  "os"
  "strings"

	"github.com/spf13/viper"
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
  err := viper.Unmarshal(&pathsMap)
  if err != nil {
    fmt.Println("Error while reading configuration file")
  }
  if len(pathsMap) == 0 {
    fmt.Println(ErrConfigEmpty)
  }else {
    storeFiles()
  }
  fmt.Println(pathsMap)
}

func storeFiles() {
  fmt.Println(len(pathsMap))
  for name, path := range pathsMap{
    symlinkFiles(name, path)
  }
}

func symlinkFiles(name, path string) {
  homePath, err := os.UserHomeDir()
  if err != nil {
    fmt.Println(err)
  }
  p := homePath + "/" + path
  fmt.Println("current path", p)

  l := strings.Split(p, "/")
  //l = strings.Join(l[len(l) - 1], "")
  fmt.Println(p, dirPath, l)
  //moveFile(p, l) 

  //fmt.Printf("Symlink %s at path: %s --> created\n", name, path)
}

func moveFile(oldLocation, newLocation string) {
  err := os.Rename(oldLocation, newLocation)
  if err != nil {
    fmt.Println(err)
  }
}

