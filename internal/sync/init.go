package sync

import (
  "log"

  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/config"
)

func InitGit() {

  r, err := git.PlainInit("./dotfiles", false)
  if err != nil {
    log.Fatal(err)
  }

  _, err = r.CreateRemote(&config.RemoteConfig{
    URLs:  []string{"git@github.com:themadnesstony/test.git"},
    Name: "origin",
  })
  if err != nil {
    log.Fatal(err)
  }
}
