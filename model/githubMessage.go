package model

type GithubMessage struct {
	HookId		int64		`json:"hook_id"`
	Repository	Repository	`json:"repository"`
	Pusher		Pusher		`json:"pusher"`
	HeadCommit	Commit		`json:"head_commit"`
}

type Pusher struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
}

type Repository struct {
	FullName	string		`json:"full_name" gorm:"full_name;primary_key"`
	Name		string		`json:"name" gorm:"-"`
	Subscribers	[]Subscriber	`gorm:"many2many:repository_subscribers"`
}

type Commit struct {
	ID			string		`json:"id"`
	TreeId		string		`json:"tree_id"`
	Message		string		`json:"message"`
	Timestamp	string		`json:"timestamp"`
	Url			string		`json:"url"`
	Author		Author		`json:"author"`
	Committer	Committer	`json:"committer"`
	Added		[]string	`json:"added"`
	Removed		[]string	`json:"removed"`
	Modified	[]string	`json:"modified"`
}

type Author struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Username	string		`json:"username"`
}

type Committer struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Username	string		`json:"username"`
}