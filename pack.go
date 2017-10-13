package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"github.com/douglarek/leftpad"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func main() {
	intro := color.New(color.FgCyan, color.Bold)
	intro.Println("Available script commands in package.json")
	intro.Println("-----------------------------------------")

	file, e := ioutil.ReadFile("./package.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	json := string(file)
	scripts := gjson.Get(json, "scripts")

	title := color.New(color.FgHiWhite, color.Bold)
	command := color.New(color.FgWhite)

	length := 0
	scripts.ForEach(func(key, value gjson.Result) bool {
		if utf8.RuneCountInString(key.String()) > length {
			length = utf8.RuneCountInString(key.String())
		}
		return true
	})

	scripts.ForEach(func(key, value gjson.Result) bool {
		title.Print(leftpad.Leftpad(key.String(), length, ' '))
		command.Printf(" %s\n", value.String())
		return true
	})
}
