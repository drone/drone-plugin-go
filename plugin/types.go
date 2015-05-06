package plugin

type Clone struct {
	Origin string `json:"origin"`
	Remote string `json:"remote"`
	Branch string `json:"branch"`
	Sha    string `json:"sha"`
	Ref    string `json:"ref"`
	Dir    string `json:"dir"`

	Netrc struct {
		Machine  string `json:"machine"`
		Login    string `json:"login"`
		Password string `json:"user"`
	}

	Keypair struct {
		Public  string `json:"public"`
		Private string `json:"private"`
	}
}

type Repo struct {
	Owner    string `json:"owner"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Language string `json:"language"`
	Private  bool   `json:"private"`
	Link     string `json:"link_url"`
	Clone    string `json:"clone_url"`
	Branch   string `json:"default_branch"`
	Last     Build  `json:"last_build,omitempty"`
}

type Build struct {
	Number      int         `json:"number"`
	State       string      `json:"state"`
	Duration    int64       `json:"duration"`
	Started     int64       `json:"started_at"`
	Finished    int64       `json:"finished_at"`
	Created     int64       `json:"created_at"`
	Updated     int64       `json:"updated_at"`
	Commit      Commit      `json:"head_commit,omitempty"`
	PullRequest PullRequest `json:"pull_request,omitempty"`
	Tasks       []Task      `json:"tasks,omitempty"`
}

type Commit struct {
	Sha       string `json:"sha,omitempty"`
	Ref       string `json:"ref,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Author    Author `json:"author,omitempty"`
	Remote    Remote `json:"repo,omitempty"`
}

type PullRequest struct {
	Number int    `json:"number,omitempty"`
	Title  string `json:"title,omitempty"`
	Source Commit `json:"source,omitempty"`
	Target Commit `json:"target,omitempty"`
}

type Author struct {
	Name     string `json:"name,omitempty"`
	Login    string `json:"login,omitempty"`
	Email    string `json:"email,omitempty"`
	Gravatar string `json:"gravatar_id,omitempty"`
}

type Remote struct {
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Clone    string `json:"clone_url,omitempty"`
}

type Task struct {
	Number   int    `json:"number"`
	State    string `json:"state"`
	ExitCode int    `json:"exit_code"`
	Duration int64  `json:"duration"`
	Started  int64  `json:"started_at"`
	Finished int64  `json:"finished_at"`

	Environment map[string]string `json:"environment,omitempty"`
}

type Links struct {
	Repo   string `json:"repo_url"`
	Commit string `json:"commit_url"`
}
