package main

import (
	"fmt"

	"github.com/drone/drone-plugin-go/plugin"
)

func main() {
	repo := plugin.Repo{}
	commit := plugin.Commit{}
	plugin.Param("repo", &repo)
	plugin.Param("commit", &commit)
	plugin.Parse()

	fmt.Printf("hello %s/%s/%v\n", repo.Owner, repo.Name, commit.Sequence)
}
