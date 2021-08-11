package sync

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func InitGit(path string) {

	r, err := git.PlainInit(path, false)
	if err != nil {
		log.Fatal(err)
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		URLs: []string{"git@github.com:rwxmad/test.git"},
		Name: "origin",
	})
	if err != nil {
		log.Fatal(err)
	}
}
