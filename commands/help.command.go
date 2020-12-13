package commands

import (
	"fmt"

	"github.com/saptarshibasu15/gitgo/utils"
)

func Help() {
	prompts := map[string]string{
		"info": "gitgo <username> info",
		"repo": "gitgo <username> repo <repository_name>",
	}
	for cmd, help := range prompts {
		fmt.Printf("%s \t %s\n", utils.Green(cmd), help)
	}
}
