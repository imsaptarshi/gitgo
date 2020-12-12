package main

import (
	"encoding/json"
	"fmt"
	types "gitgo/types"
	"gitgo/utils"
	"log"
	"os"
)

func getUserInfo(u string) types.User {
	var user types.User
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/users/%s", u)), &user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return user
}

func getRepoInfo(owner string, name string) types.Repo {
	var repo types.Repo
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return repo
}

func printHelp() {
	// TODO
}

func main() {
	if len(os.Args) < 3 {
		printHelp()
		return
	}

	user := os.Args[1]
	op := os.Args[2]

	if user == "repo" {
		// TODO
		return
	}

	if op == "info" {
		u := getUserInfo(user)
		utils.Display(u)
	} else if op == "repo" {
		if len(os.Args) == 4 {
			n := os.Args[3]
			getRepoInfo(user, n)
		}
	}
}
