package main

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}
