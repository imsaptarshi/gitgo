package main

import (
	"encoding/json"
	"fmt"
	types "gitgo/types"
	"gitgo/utils"
	"log"
	"os"
	"reflect"
)

func getUserInfo(u string) types.User {
	var user types.User
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/users/%s", u)), &user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return user
}

func getRepoInfo(owner string, name string) types.Repo {
	var repo types.Repo
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name)), &repo)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return repo
}

//...
func Display(u types.User) {
	v := reflect.ValueOf(u)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil {
			continue
		}
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}
func printHelp() {
	// TODO
}

func main() {

	if len(os.Args) < 3 {
		printHelp()
	}

	user := os.Args[1]
	op := os.Args[2]

	if user == "repo" {
		return
	}

	if op == "info" {
		u := getUserInfo(user)
		Display(u)
	} else if op == "repo" {
		if len(os.Args) == 4 {
			n := os.Args[3]
			getRepoInfo(user, n)
		} // g rishit repo name
	}
}
