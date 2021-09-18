package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "show and edit dotsync and git config files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dotsyncConfig()
	},
}

var (
	showGit bool
)

type GitConfig struct {
	Url         string
	AllSettings []byte
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&showGit, "git", "g", false, "Show .git/config for dotfiles")
}

func dotsyncConfig() {
	InitConfig(&cfgFile)
	if showGit {
		gitConfig()
	} else {
		showDotsyncConfig()
	}
}

func showDotsyncConfig() {
	fmt.Println(color.GreenString("[ dotsync config ]"))
	configSettings := v.AllSettings()
	for fileName, path := range configSettings {
		fmt.Printf("%s: %s \n", fileName, path)
	}
}

func gitConfig() {
	r, err := git.PlainOpen(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	config, err := r.ConfigScoped(0)
	if err != nil {
		log.Fatal(err)
	}
	configRemoteSection := config.Raw.Sections[1]
	configOriginSubsection := configRemoteSection.Subsection("origin")

	var configAllSettings []byte
	configAllSettings, err = config.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	// Current url in .git/config
	configRemoteUrl := configOriginSubsection.Options.Get("url")

	gconf := GitConfig{configRemoteUrl, configAllSettings}
	gconf.showGitConfig()
}

func (gc *GitConfig) showGitConfig() {
	fmt.Println(color.GreenString("[ git config ]"))
	fmt.Println(string(gc.AllSettings))
}
