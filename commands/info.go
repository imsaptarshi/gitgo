package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	types "github.com/saptarshibasu15/gitgo/types"
	"github.com/saptarshibasu15/gitgo/utils"
)

func Info(u string) {
	var user types.User
	jsonErr := json.Unmarshal(utils.Fetch(fmt.Sprintf("https://api.github.com/users/%s", u)), &user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	displayInfo(user)
}

func displayInfo(u types.User) {
	v := reflect.ValueOf(u)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil || utils.Filter(typeOfS.Field(i).Name) {
			continue
		}
		fmt.Printf("%s: %v\n", utils.Green(typeOfS.Field(i).Name), v.Field(i).Interface())
	}
}
