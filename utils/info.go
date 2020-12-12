package utils

import (
	"fmt"
	"gitgo/types"
	"reflect"
)

func Display(u types.User) {
	v := reflect.ValueOf(u)
	typeof := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Interface() == nil || field.Interface() == "" || IsInFilter(typeof.Field(i).Name) {
			continue
		}
		fmt.Printf("%s: %v\n", typeof.Field(i).Name, field.Interface())
	}
}
