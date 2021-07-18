package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add current file in configutaion list",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    addFile()	
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addFile() {
  fmt.Println(cfgFile)
  fmt.Println(viper.ConfigFileUsed())
  fmt.Println("add file is here")
}
