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
		if v.Field(i).Interface() == nil || /*( [] && !utils.Contains(typeOfS.Field(i).Name, ops)) ||*/ utils.Filter(typeOfS.Field(i).Name) {
			continue
		}
		fmt.Printf("%s: %v\n", utils.Green(typeOfS.Field(i).Name), v.Field(i).Interface())
	}
}
