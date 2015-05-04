package main

import (
	"fmt"

	"github.com/drone/drone-plugin-go/plugin"
)

func main() {
	repo := plugin.Repo{}
	build := plugin.Build{}
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Parse()

	fmt.Printf("hello %s/%s/%v\n", repo.Owner, repo.Name, build.Number)
}
