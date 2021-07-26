package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
  cfgFile string
  dirPath string
  home    string
  v       = viper.NewWithOptions(viper.KeyDelimiter("::"))
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Directory and configuration file initialization",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
    initCommand()
	},
}

func init() {
  rootCmd.AddCommand(initCmd)

  // Define directory for configs
  home, err := os.UserHomeDir()
  if err != nil {
    fmt.Println(err)
  }
  dirPath = home + "/code/go/src/dotsync/dotfiles/"
}

func initCommand() {
	initDir()
	InitConfig(&cfgFile)
}

func initDir() {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		dir := os.Mkdir(dirPath, os.FileMode(0777))
    fmt.Println(dir)
    fmt.Printf("Directory created at %v", dirPath)
	} else {
		var choice string
		fmt.Println("Directory already exists, do you want to rewrite? [Y/n]")
		switch fmt.Scan(&choice); strings.ToLower(choice) {
		case "y":
			err := os.RemoveAll(dirPath)
			if err != nil {
				log.Fatal("Can't remove a directory")
			}
			dir := os.Mkdir(dirPath, os.FileMode(0777))
			log.Println(dir)
		case "n":
			fmt.Println("Cancelled")
		}
	}
}

func InitConfig(c *string) {
  v.AddConfigPath(dirPath)
  v.SetConfigType("yaml")
  v.SetConfigName(".dotsync")
	if err := v.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", v.ConfigFileUsed())
	} else {
    err := v.SafeWriteConfig()
    if err != nil {
      fmt.Println("Error with writing config file")
    }
  }
  *c = v.ConfigFileUsed()
}

