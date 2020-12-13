package commands

import (
	"fmt"

	"github.com/saptarshibasu15/gitgo/utils"
)

func Help() {
	utils.AsciiBanner()
	prompts := map[string]string{
		"help":                 "gitgo or gitgo help",
		"info":                 "gitgo <username> info",
		"repo":                 "gitgo <username> repo <repository_name>",
		"repo *":               "gitgo <username> repo *",
		"repo (-i | --issues)": "gitgo <username> repo <repository_name> (-i | --issues)",
	}
	for cmd, help := range prompts {
		fmt.Printf("%s \t %s\n", utils.Green(cmd), help)
	}
}
