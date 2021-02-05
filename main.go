package main

import (
	"os"

	"github.com/saptarshibasu15/gitgo/commands"
)

func main() {
	if len(os.Args) == 1 {
		commands.Help()
		return
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "login" {
			commands.Login()
		}
		return
	}

	user, op := os.Args[1], os.Args[2]

	switch op {
	case "info":
		commands.Info(user)
	case "repo": // g rishit repo name
		if len(os.Args) > 3 {
			if len(os.Args) <= 5 && os.Args[3] != "*" {
				name := os.Args[3]
				if len(os.Args) == 5 {
					if os.Args[4] == "-i" || os.Args[4] == "--issues" {
						commands.RepoIssues(user, name, os.Args[5:]...) //g user repo name -i/--issues
					}
          if os.Args[4] == "-pr" || os.Args[4] == "--pulls" {
						commands.RepoPulls(user, name, os.Args[5:]...) //g user repo name -i/--issues
					}
				} else {
					commands.Repo(user, name, os.Args[4:]...)
				}
			} else {

				commands.RepoList(user)
			}
		}
	}
}
