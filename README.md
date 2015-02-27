drone-plugin-go
===============

This is a package with simple support for writing Drone plugins in Go.

## Overview

Plugins are executable files run by Drone to customize the build lifecycle. Plugins receive input data from `stdin` and writes results to `stdout`

```sh
./slack-plugin <<EOF
{
    "repo" : {
        "host": "github.com",
        "owner": "foo",
        "name": "bar"
    },
    "commit" : {
        "status": "Success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "sha": "9f2849d5",
        "branch": "master",
        "pull_request": "800",
        "author": "john.smith@gmail.com",
        "message": "Update the Readme"
    },
    "clone" : {
        "branch": "master",
        "remote": "git://github.com/drone/drone",
        "dir": "/drone/src/github.com/drone/drone",
        "ref": "refs/heads/master",
        "sha": "436b7a6e2abaddfd35740527353e78a227ddcb2c"
    }
}
EOF
```

To read this data you can use the `plugin` package to declare and parse parameters. If you've worked with the `flag` package this should look familiar:

```Go
var repo = plugin.Repo{}
var clone = plugin.Clonse{}
var commit = plugin.Commit{}
var config = struct {
	URL      string `json:"webhook_url"`
    Username string `json:"username"`
    Channel  string `json:"channel"`
}{}

plugin.Param(&repo)
plugin.Param(&commit)
plugin.Param(&config)
plugin.Parse()
```

### Shared Volumes

The repository clone directory (ie `clone.dir` input parameter) will be shared across all plugins. This means that anything files in your repository directory or subdirectories are accessible to plugins. This is useful for plugins that analyze or archive files, such as an S3 plugin.
