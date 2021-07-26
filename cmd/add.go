package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
  "strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a file to the synchronized directory and writes it to the configuration file",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("You need to specify config file")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		addFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addFile(c string) {
	InitConfig(&cfgFile)

	cfgAbsPath, err := filepath.Abs(c)
	if err != nil {
		fmt.Println("Error with filepath")
	}

  pathSlice := strings.Split(cfgAbsPath, "/")
  pathSlice = append(pathSlice[3:])
  filePath := strings.Join(pathSlice, "/")

	v.Set(c, filePath)
	v.WriteConfig()
}
