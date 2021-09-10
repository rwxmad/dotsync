package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Writes the file to the configuration file",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("You need to specify config file")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		addFile(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addFile(files []string) {
	InitConfig(&cfgFile)

	for _, file := range files {
		cfgAbsPath, err := filepath.Abs(file)
		if err != nil {
			fmt.Println("Error with filepath")
		}

		pathSlice := strings.Split(cfgAbsPath, "/")
		pathSlice = append(pathSlice[3:])
		filePath := strings.Join(pathSlice, "/")

		fileNameSlice := strings.Split(file, "/")
		fileName := fileNameSlice[len(fileNameSlice)-1]

		if v.Get(fileName) == nil {
			v.Set(fileName, filePath)
			v.WriteConfig()
			fmt.Printf("# [ %s ] was successfully written to the config file \n", color.GreenString(fileName))
		} else {
			fmt.Printf("# [ %s ] is already written to the config file \n", color.CyanString(fileName))
		}
	}
}
