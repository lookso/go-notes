package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type student struct {
	Age    int    `json:"age,omitempty"`
	Name   string `json:"name,omitempty"`
	School string `json:"school,omitempty"`
}

var st = student{
	Age:    10,
	Name:   "john smith",
	School: "a high school",
}

var dic = map[string]int{
	"age":    0,
	"name":   1,
	"school": 2,
}

var filters = []string{
	"name",
	"school",
}

func initStudentElems(st *student, fields []string) bool {
	v := reflect.Indirect(reflect.ValueOf(st))
	for _, field := range fields {
		idx, exist := dic[field]
		if !exist {
			return false
		}
		vf := v.Field(idx)
		vf.Set(reflect.Zero(vf.Type()))
	}
	return true
}

func printMarshalIdent(st *student) error {
	b, err := json.MarshalIndent(st, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func main() {
	printMarshalIdent(&st)
	initStudentElems(&st, filters)
	printMarshalIdent(&st)
}
