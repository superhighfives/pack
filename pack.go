package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

// listScripts writes the available script commands from a package.json to
// stdout, or an actionable message to stderr, and returns a process exit code.
// data is the raw package.json contents and readErr is any error from reading
// it (nil if the file was read successfully).
func listScripts(stdout, stderr io.Writer, data []byte, readErr error) int {
	warning := color.New(color.FgYellow, color.Bold)

	if readErr != nil {
		if os.IsNotExist(readErr) {
			warning.Fprintln(stderr, "No package.json found in this directory.")
			fmt.Fprintln(stderr, "Run pack in a directory that contains a package.json file.")
		} else {
			warning.Fprintf(stderr, "Couldn't read package.json: %v\n", readErr)
		}
		return 1
	}

	if !gjson.ValidBytes(data) {
		warning.Fprintln(stderr, "package.json isn't valid JSON.")
		return 1
	}

	scripts := gjson.GetBytes(data, "scripts")

	if scripts.Exists() && !scripts.IsObject() {
		warning.Fprintln(stderr, `The "scripts" field in package.json isn't an object.`)
		return 1
	}

	if !scripts.Exists() || len(scripts.Map()) == 0 {
		warning.Fprintln(stderr, "No script commands found in package.json.")
		fmt.Fprintln(stderr, `Add a "scripts" section to package.json to see commands here.`)
		return 1
	}

	intro := color.New(color.FgCyan, color.Bold)
	intro.Fprintln(stdout, "Available script commands in package.json")
	intro.Fprintln(stdout, "-----------------------------------------")

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
		title.Fprint(stdout, fmt.Sprintf("%*s", length, key.String()))
		command.Fprintf(stdout, " %s\n", value.String())
		return true
	})

	return 0
}

func main() {
	data, err := os.ReadFile("./package.json")
	os.Exit(listScripts(os.Stdout, os.Stderr, data, err))
}
