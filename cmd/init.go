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
  dirPath string = "./dotfiles"
  defaultSettings = map[string]string {
    ".tmux.conf": "~/.tmux.conf",
    ".vim": "~/.vim",
  }
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
}

func initCommand() {
	viper.AutomaticEnv() // read in environment variables that match
  fmt.Println(os.UserHomeDir())
  //home, err := os.UserHomeDir()
  //fmt.Printf("Home dir is -> %s \n", home)
  //cobra.CheckErr(err)
	initDir()
	initConfig()
}

func initDir() {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		dir := os.Mkdir(dirPath, os.FileMode(0777))
    fmt.Println(dir)
    fmt.Printf("Directory created at %v", dirPath)
	} else {
		var choice string
		fmt.Println("Directory already exists, do you want to rewite? [Y/n]")
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

func initConfig() {
  viper.AddConfigPath(dirPath)
  viper.SetConfigType("yaml")
  viper.SetConfigName(".dotsync")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {

    viper.SetDefault("dotfiles", defaultSettings)

    err := viper.SafeWriteConfig()
    if err != nil {
    fmt.Println("Error with writing config file")
    }
    cfgFile = viper.ConfigFileUsed()
    fmt.Println(viper.ConfigFileUsed())
    fmt.Println(viper.AllSettings())
  }
}

