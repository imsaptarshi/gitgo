package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
)

func Repo(owner string, name string) {
	var repo types.Repo
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	displayRepo(repo)
}

func displayRepo(u types.Repo) {
	v := reflect.ValueOf(u)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil {
			continue
		}
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}
