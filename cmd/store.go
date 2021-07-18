package cmd

import (
	"fmt"
  "os"
  "log"
  "strings"

	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store, symlink and sync dotfiles",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- Store is complete ---")

	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
