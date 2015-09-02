package plugin

// Repo represents a version control repository.
type Repo struct {
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Self     string `json:"self_url"`
	Link     string `json:"link_url"`
	Clone    string `json:"clone_url"`
	Branch   string `json:"default_branch"`
	Private  bool   `json:"private"`
	Trusted  bool   `json:"trusted"`
	Timeout  int64  `json:"timeout"`
}

// System provides important information about the Drone
// server to the plugin.
type System struct {
	Version string   `json:"version"`
	Link    string   `json:"link_url"`
	Plugins []string `json:"plugins"`
	Globals []string `json:"globals"`
}

// Workspace defines the build's workspace inside the
// container. This helps the plugin locate the source
// code directory.
type Workspace struct {
	Root string `json:"root"`
	Path string `json:"path"`

	Netrc *Netrc   `json:"netrc"`
	Keys  *Keypair `json:"keys"`
}

// Keypair represents an RSA public and private key assigned to a
// repository. It may be used to clone private repositories, or as
// a deployment key.
type Keypair struct {
	Public  string `json:"public,omitempty"`
	Private string `json:"private,omitempty"`
}

// Netrc defines a default .netrc file that should be injected
// into the build environment. It will be used to authorize access
// to https resources, such as git+https clones.
type Netrc struct {
	Machine  string `json:"machine"`
	Login    string `json:"login"`
	Password string `json:"user"`
}

// Build represents the process of compiling and testing a changeset,
// typically triggered by the remote system (ie GitHub).
type Build struct {
	Number   int    `json:"number"`
	Status   string `json:"status"`
	Started  int64  `json:"started_at"`
	Finished int64  `json:"finished_at"`

	// Commit that is being built. If this is a pull request it
	// represents the source commit (from the fork)
	Commit *Commit `json:"head_commit"`

	// PullRequest that is being built. Nil if not a pull request.
	PullRequest *PullRequest `json:"pull_request"`
}

// PullRequest represents a proposed revision to
// the repository.
type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`

	// Base represents the base commit on top of which
	// the change was made. It is the target of this
	// pull request.
	Base *Commit `json:"base_commit"`
}

// Commit represents a revision made to the repository.
type Commit struct {
	Sha       string  `json:"sha"`
	Ref       string  `json:"ref"`
	Branch    string  `json:"branch"`
	Message   string  `json:"message"`
	Timestamp string  `json:"timestamp"`
	Remote    string  `json:"remote"`
	Author    *Author `json:"author"`
}

// Author represents the author of a code change.
type Author struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

// Job represents a single job that is being executed as part
// of a Build.
type Job struct {
	ID       int64  `json:"id"`
	Number   int    `json:"number"`
	Status   string `json:"status"`
	ExitCode int    `json:"exit_code"`
	Started  int64  `json:"started_at"`
	Finished int64  `json:"finished_at"`

	Environment map[string]string `json:"environment"`
}
