package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotsync",
	Short: "CLI for sync your dotfiles",
	Long:  `CLI for synchronize your dotfiles, between multiple machines from a git repository`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
