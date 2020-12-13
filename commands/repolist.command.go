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
	fmt.Println("[")
	for _, r := range u {
		v := reflect.ValueOf(r)
		typeOfS := v.Type()
		fmt.Println("\t{")
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Interface() == nil {
				continue
			}
			fmt.Printf("\t\t%s : %v\n", utils.Green(typeOfS.Field(i).Name), v.Field(i).Interface())
		}
		fmt.Println("\t\b},")
	}
	fmt.Println("]")
}
