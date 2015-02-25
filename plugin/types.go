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

type User struct {
	Remote string `json:"remote"`
	Login  string `json:"login"`
	Name   string `json:"name"`
	Email  string `json:"email,omitempty"`
}

type Repo struct {
	Remote string `json:"remote"`
	Host   string `json:"host"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	URL    string `json:"url"`
}

type Commit struct {
	Status      string `json:"status"`
	Started     int64  `json:"started_at"`
	Finished    int64  `json:"finished_at"`
	Duration    int64  `json:"duration"`
	Sha         string `json:"sha"`
	Branch      string `json:"branch"`
	PullRequest string `json:"pull_request"`
	Author      string `json:"author"`
	Gravatar    string `json:"gravatar"`
	Timestamp   string `json:"timestamp"`
	Message     string `json:"message"`
}

type Config struct {
	Image    string   `json:"image"`
	Env      []string `json:"env"`
	Script   []string `json:"script"`
	Branches []string `json:"branches"`
	Services []string `json:"services"`
}
