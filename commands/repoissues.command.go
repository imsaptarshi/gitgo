package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
)

func RepoIssues(owner string, name string, ops ...string) {
	var repo types.RepoIssues
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, name)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	displayRepoIssues(repo, name, ops...)
}

func displayRepoIssues(u types.RepoIssues, name string, ops ...string) {

	fmt.Printf("%v/issues : %v", utils.Green(name), utils.White(len(u)))
	for _, r := range u {
		v := reflect.ValueOf(r)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			if v.Field(i).Interface() == nil || utils.Filter(typeOfS.Field(i).Name) {
				continue
			}
			x := ""
			switch typeOfS.Field(i).Name {
			case "Title":
				x = utils.LBlue(fmt.Sprintf("\n     └── Title %v", utils.White(v.Field(i).Interface())))
			case "Number":
				x = utils.Yellow(fmt.Sprintf("\n           ├── # %v", utils.White(v.Field(i).Interface())))

			case "State":
				x = utils.Yellow(fmt.Sprintf("\n           └── State %v", utils.White(v.Field(i).Interface())))
			}
			if x == "" {
				continue
			}
			fmt.Printf("%s", x)
		}
	}
	fmt.Println()
}
