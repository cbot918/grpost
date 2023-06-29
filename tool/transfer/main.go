package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const file = "post.json"

func main() {
	PrintStruct(file)
}

func PrintStruct(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("read file failed")
		panic(err)
	}
	// fmt.Println(string(content))
	p := &Post{}
	err = json.Unmarshal(content, &p)
	if err != nil {
		fmt.Println("json unmarshal error")
		panic(err)
	}
	printJson(p)
}

func printJson(target interface{}) {
	jsonData, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
