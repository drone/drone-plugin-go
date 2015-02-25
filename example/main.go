package main

import (
	"fmt"

	"github.com/drone/drone-plugin-go/plugin"
)

type Repo struct {
	Host  string `json:"host"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

func main() {
	repo := Repo{}
	plugin.Param("repo", &repo)
	plugin.Parse()

	fmt.Printf("git://%s/%s/%s.git\n", repo.Host, repo.Owner, repo.Name)
}
