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
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Language    string `json:"language"`
	Private     bool   `json:"private"`
	Self        string `json:"self_url"`
	Link        string `json:"link_url"`
	Clone       string `json:"clone_url"`
	Branch      string `json:"default_branch"`
	Timeout     int64  `json:"timeout"`
	Trusted     bool   `json:"trusted"`
	PostCommit  bool   `json:"post_commits"`
	PullRequest bool   `json:"pull_requests"`
	Created     int64  `json:"created_at"`
	Updated     int64  `json:"updated_at"`
}

type Commit struct {
	Sequence     int    `json:"sequence"`
	State        string `json:"state"`
	Started      int64  `json:"started_at"`
	Finished     int64  `json:"finished_at"`
	Sha          string `json:"sha"`
	Ref          string `json:"ref"`
	PullRequest  string `json:"pull_request,omitempty"`
	Branch       string `json:"branch"`
	Author       string `json:"author"`
	Gravatar     string `json:"gravatar"`
	Timestamp    string `json:"timestamp"`
	Message      string `json:"message"`
	SourceRemote string `json:"source_remote,omitempty"`
	SourceBranch string `json:"source_branch,omitempty"`
	SourceSha    string `json:"source_sha,omitempty"`
	Created      int64  `json:"created_at"`
	Updated      int64  `json:"updated_at"`

	Builds []*Build `json:"builds,omitempty"`
}

type Build struct {
	State    string `json:"state"`
	ExitCode int    `json:"exit_code"`
	Sequence int    `json:"sequence"`
	Duration int64  `json:"duration"`
	Started  int64  `json:"started_at"`
	Finished int64  `json:"finished_at"`
	Created  int64  `json:"created_at"`
	Updated  int64  `json:"updated_at"`

	Environment map[string]string `json:"environment"`
}
