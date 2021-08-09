package sync

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func Sync(path string) {
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	status, err := w.Status()
	if err != nil {
		fmt.Println(err)
	}

	var commitLog = map[string]string{}

	for file, statusCode := range status {
		switch code := string(statusCode.Staging); code {
		case "?":
			commitLog[file] = color.GreenString("Added")
		case "M":
			commitLog[file] = color.CyanString("Updated")
		}
	}

	for k, v := range commitLog {
		fmt.Printf("# [ %v ]: %v\n", k, v)
	}

	Add(w)
	Commit(w, r)
	Push(r)
}

func Add(w *git.Worktree) {
	_, err := w.Add(".")
	if err != nil {
		log.Fatal(err)
	}
}

func Commit(w *git.Worktree, r *git.Repository) {
	var commitMessage string
	fmt.Print("Enter commit message: ")
	inputReader := bufio.NewReader(os.Stdin)
	commitMessage, _ = inputReader.ReadString('\n')
	commit, err := w.Commit(commitMessage, &git.CommitOptions{})
	if err != nil {
		log.Fatal(err)
	}

	obj, err := r.CommitObject(commit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj)
}

func Push(r *git.Repository) {
	authSSH, err := ssh.NewPublicKeysFromFile("git", "/Users/madnesstony/.ssh/id_rsa", "")
	err = r.Push(&git.PushOptions{Auth: authSSH, Progress: os.Stdout})
	if err != nil {
		fmt.Println(err)
	}
}
