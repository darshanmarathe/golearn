package main

import (
	"fmt"
)

var Commands map[string]string

func main() {
	fmt.Println("Demo for maps")
	LinuxCommands()
	for k, v := range Commands {
		fmt.Println(k, "=>", v)
	}

	WindowsCommands()
	for k, v := range Commands {
		fmt.Println(k, "=>", v)
	}

	delete(Commands, "Dir")
	fmt.Println("Delete key and print agian")
	for k, v := range Commands {
		fmt.Println(k, "=>", v)
	}
}

func LinuxCommands() {
	Commands = make(map[string]string)

	Commands["ls"] = "Show all files and folders in directory"
	Commands["cat"] = "read a file"
	Commands["touch"] = "create a new file"

}

func WindowsCommands() {
	Commands = map[string]string{
		"Dir":             "Show all files and folders in dir",
		"cd":              "Change Directory",
		"md":              "Make Directory",
		"type <filename>": "Reads a file",
	}

	Commands["refreshenv"] = "refreshes the cmd env"
}
