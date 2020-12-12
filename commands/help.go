package commands

import "fmt"

func Help() {
	prompts := map[string]string{
		"info": "gitgo <username> info",
		"repo": "gitgo <username> repo <repository_name>",
	}
	for cmd, help := range prompts {
		fmt.Println(cmd, "\t", help)
	}
}
