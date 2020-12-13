package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
)

func RepoList(owner string) {
	var repo types.RepoList
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/users/%s/repos", owner)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	displayRepoList(repo)
}

func displayRepoList(u types.RepoList, ops ...string) {

	for _, r := range u {
		v := reflect.ValueOf(r)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Interface() == nil || utils.Filter(typeOfS.Field(i).Name) {
				continue
			}
			x := ""
			switch typeOfS.Field(i).Name {
			case "StargazersCount":
				x = utils.Yellow(fmt.Sprintf("\n- Stars %d", v.Field(i).Interface()))
			case "Name":
				x = utils.Green(fmt.Sprintf("%v: ", v.Field(i).Interface()))
			case "Forks":
				x = utils.Yellow(fmt.Sprintf("\n- Forks %v ", v.Field(i).Interface()))
				// default:
				// x = typeOfS.Field(i).Name
			}
			if x == "" {
				continue
			}
			fmt.Printf("%s", x)
		}
		fmt.Print("\n")
	}
}
