package main

import (
	"os"

	"github.com/saptarshibasu15/gitgo/commands"
)

func main() {
	if len(os.Args) < 3 {
		commands.Help()
		return
	}

	user, op := os.Args[1], os.Args[2]

	switch op {
	case "info":
		commands.Info(user)
	case "repo": // g rishit repo name
		if len(os.Args) > 3 {
			if len(os.Args) <= 4 {
				name := os.Args[3]

				commands.Repo(user, name, os.Args[4:]...)
			}
		} else {
			// commands.ListRepo(user)
		}
	}
}
