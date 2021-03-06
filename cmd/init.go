package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxmad/dotsync/internal/sync"

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

	// Define directory for config files
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	// FIXME: return default path for home dir
	dirPath = home + "/code/go/src/dotsync/dotfiles/"
}

func initCommand() {
	initDir()
	sync.InitGit(dirPath)
	InitConfig(&cfgFile)
}

func initDir() {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.FileMode(0777))
		if err != nil {
			log.Fatal("Failed to create directory")
		}
		fmt.Printf("Directory created at %v \n", dirPath)
	} else {
		var choice string
		fmt.Print("Directory already exists, do you want to rewrite? [Y/n] ")
		switch fmt.Scan(&choice); strings.ToLower(choice) {
		case "y":
			err := os.RemoveAll(dirPath)
			if err != nil {
				log.Fatal("Can't remove a directory")
			}
			err = os.Mkdir(dirPath, os.FileMode(0777))
			if err != nil {
				log.Fatal("Can't create a directory")
			}
		case "n":
			fmt.Println("Cancelled")
		}
	}
}

// InitConfig function
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
		err = v.ReadInConfig()
		if err != nil {
			log.Fatal("dotsync configuration file not found, use \"dotsync init\"")
		}
	}
	*c = v.ConfigFileUsed()
}
