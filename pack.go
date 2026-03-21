package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func main() {
	intro := color.New(color.FgCyan, color.Bold)
	intro.Println("Available script commands in package.json")
	intro.Println("-----------------------------------------")

	file, err := os.ReadFile("./package.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %v\n", err)
		os.Exit(1)
	}

	scripts := gjson.GetBytes(file, "scripts")

	title := color.New(color.FgHiWhite, color.Bold)
	command := color.New(color.FgWhite)

	length := 0
	scripts.ForEach(func(key, value gjson.Result) bool {
		if n := utf8.RuneCountInString(key.String()); n > length {
			length = n
		}
		return true
	})

	scripts.ForEach(func(key, value gjson.Result) bool {
		title.Print(fmt.Sprintf("%*s", length, key.String()))
		command.Printf(" %s\n", value.String())
		return true
	})
}
