package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

func main() {
	warning := color.New(color.FgYellow, color.Bold)

	file, err := os.ReadFile("./package.json")
	if err != nil {
		if os.IsNotExist(err) {
			warning.Fprintln(os.Stderr, "No package.json found in this directory.")
			fmt.Fprintln(os.Stderr, "Run pack in a directory that contains a package.json file.")
		} else {
			warning.Fprintf(os.Stderr, "Couldn't read package.json: %v\n", err)
		}
		os.Exit(1)
	}

	if !gjson.ValidBytes(file) {
		warning.Fprintln(os.Stderr, "package.json isn't valid JSON.")
		os.Exit(1)
	}

	scripts := gjson.GetBytes(file, "scripts")

	if !scripts.Exists() || len(scripts.Map()) == 0 {
		warning.Fprintln(os.Stderr, "No script commands found in package.json.")
		fmt.Fprintln(os.Stderr, "Add a \"scripts\" section to package.json to see commands here.")
		os.Exit(1)
	}

	intro := color.New(color.FgCyan, color.Bold)
	intro.Println("Available script commands in package.json")
	intro.Println("-----------------------------------------")

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
