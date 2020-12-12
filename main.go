package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
<<<<<<< HEAD
=======
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
>>>>>>> 665f9b145d4b34658cf8ec6e92dfa985eadaddf2
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

<<<<<<< HEAD
=======
//Display ...
func Display(u interface{}) {
	v := reflect.ValueOf(u)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil {
			continue
		}
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

>>>>>>> 665f9b145d4b34658cf8ec6e92dfa985eadaddf2
func printHelp() {
	prompts := map[string]string{
		"info": "gitgo <username> info",
		"repo": "gitgo <username> repo <repository_name>",
	}
	for cmd, help := range prompts {
		fmt.Println(cmd, "\t", help)
	}
}

func main() {
	if len(os.Args) < 3 {
		printHelp()
<<<<<<< HEAD
		return
	}

	user := os.Args[1]
	op := os.Args[2]

	if user == "repo" {
		// TODO
		return
	}

	if op == "info" {
		u := getUserInfo(user)
		utils.Display(u)
	} else if op == "repo" {
		if len(os.Args) == 4 {
			n := os.Args[3]
			getRepoInfo(user, n)
		}
=======
		return
	}

	user, op := os.Args[1], os.Args[2]

	switch op {
	case "info":
		res := getUserInfo(user)
		Display(res)
	case "repo": // g rishit repo name
		name := os.Args[3]
		Display(getRepoInfo(user, name))
>>>>>>> 665f9b145d4b34658cf8ec6e92dfa985eadaddf2
	}
}
