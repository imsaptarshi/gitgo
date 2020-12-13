package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
)

func Repo(owner string, name string, ops ...string) {
	var repo types.Repo
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	displayRepo(repo, ops...)
}

func displayRepo(u types.Repo, ops ...string) {
	v := reflect.ValueOf(u)
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
			if v.Field(8).Interface() == true {
				x = utils.LRed(fmt.Sprintf("%v: ", v.Field(i).Interface()))
			} else {
				x = utils.Green(fmt.Sprintf("%v: ", v.Field(i).Interface()))
			}
		case "Forks":
			x = utils.Yellow(fmt.Sprintf("\n- Forks %v ", v.Field(i).Interface()))
		}
		if x == "" {
			continue
		}
		fmt.Printf("%s", x)
	}
}
