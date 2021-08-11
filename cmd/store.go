package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/themadnesstony/dotsync/internal/sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store, symlink and sync dotfiles",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Store()
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}

var (
	pathsMap = make(map[string]string)
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
		log.Fatal("Config file is empty")
	} else {
		storeFiles()
	}
}

func storeFiles() {
	for name, path := range pathsMap {
		symlinkFile(name, path)
	}
	sync.Sync(dirPath)
}

func symlinkFile(name, path string) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	oldPath := homePath + "/" + path
	newPath := dirPath + name

	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		fmt.Printf("# [File %s at the path: %s ]: %s\n", name, oldPath, color.RedString("Not found"))
	} else {
		if _, err := os.Stat(newPath); !os.IsNotExist(err) {
			fmt.Printf("# [Symlink to the %s at the path: %s ]: %s\n", name, newPath, color.CyanString("Exists"))
		} else {
			moveFile(oldPath, newPath)
			err = os.Symlink(newPath, oldPath)
			if err != nil {
				fmt.Printf("# [Symlink to the %s at the path: %s ]: %s\n", name, newPath, color.RedString("Error"))
			} else {
				fmt.Printf("# [Symlink to the %s at the path: %s ]: %s\n", name, newPath, color.GreenString("Done"))
			}
		}
	}
}

func moveFile(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		fmt.Println(err)
	}
}
