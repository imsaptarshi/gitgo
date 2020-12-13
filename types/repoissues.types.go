package types

type RepoIssues []struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	State  string `json:"state"`
}
