package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const jsonc = "user.json"

func main() {
	s := GetStructFromJson(jsonc, &User{})
	// InsertUsers(s)
	PrintJson(s)
}

func GetStructFromJson(file string, tstruct *User) *User {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file failed")
		panic(err)
	}
	return marshaler(content, tstruct)
}

func marshaler(content []byte, target *User) *User {
	err := json.Unmarshal(content, &target)
	if err != nil {
		fmt.Println("json unmarshal error")
		panic(err)
	}
	return target
}

func PrintJson(target *User) {
	jsonData, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
