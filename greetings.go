package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Greeting holds a language and formal and informal greeting in those languages
type Greeting struct {
	Language string `json:"language"`
	Formal   string `json:"formal"`
	Informal string `json:"informal"`
}

var greetingMapping map[string]Greeting

func initGreetingMapping() {
	greetingfile, err := os.Open("greetings.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var greetingArray []Greeting

	bytes, err := ioutil.ReadAll(greetingfile)

	json.Unmarshal(bytes, &greetingArray)

	greetingMapping = make(map[string]Greeting, len(greetingArray))

	for _, r := range greetingArray {
		greetingMapping[r.Language] = r
	}
}
