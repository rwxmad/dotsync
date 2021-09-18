package sync

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func InitGit(path string) {

	fmt.Print("Enter repository URL: ")
	inputReader := bufio.NewReader(os.Stdin)
	remoteUrl, _ := inputReader.ReadString('\n')

	r, err := git.PlainInit(path, false)
	if err != nil {
		log.Fatal(err)
	}

	remote, err := createRemote(r, remoteUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(remote)
}

func createRemote(r *git.Repository, url string) (*git.Remote, error) {
	rem, err := r.CreateRemote(&config.RemoteConfig{
		URLs: []string{url},
		Name: "origin",
	})
	if err != nil {
		return nil, err
	}
	return rem, err
}
