package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
)

func RepoPulls(owner string, name string, ops ...string) {
	var repo types.RepoPulls
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", owner, name)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	displayRepoPulls(repo, name, ops...)
}

func displayRepoPulls(u types.RepoPulls, name string, ops ...string) {

	fmt.Printf("%v/pulls : %v", utils.Green(name), utils.White(len(u)))
	for k, r := range u {
		v := reflect.ValueOf(r)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			if v.Field(i).Interface() == nil || utils.Filter(typeOfS.Field(i).Name) {
				continue
			}
			x := ""
			switch typeOfS.Field(i).Name {
			case "Title":
				if k == len(u)-1 {
					x = utils.Yellow(fmt.Sprintf("\n     └── Title %v", utils.White(v.Field(i).Interface())))
				} else {
					x = utils.Yellow(fmt.Sprintf("\n     ├── Title %v", utils.White(v.Field(i).Interface())))
				}
			case "Number":
				if k == len(u)-1 {
					x = utils.Yellow(fmt.Sprintf("\n         ├── #%v", utils.White(v.Field(i).Interface())))
				} else {
					x = utils.Yellow(fmt.Sprintf("\n     │   ├── #%v", utils.White(v.Field(i).Interface())))
				}

			case "State":
				if k == len(u)-1 {
					x = utils.Yellow(fmt.Sprintf("\n         └── State %v", utils.White(v.Field(i).Interface())))
				} else {
					x = utils.Yellow(fmt.Sprintf("\n     │   └── State %v", utils.White(v.Field(i).Interface())))
				}
			}
			if x == "" {
				continue
			}
			fmt.Printf("%s", x)
		}
	}
	fmt.Println()
}

//https://api.github.com/repos/milan090/courseyard/pulls