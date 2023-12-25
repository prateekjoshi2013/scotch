package main

import (
	"errors"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/prateekjoshi2013/scotch"
)

const Version = "1.0.0"

var sco scotch.Scotch

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGraceFully(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version: %s\n", Version)
	default:
		log.Println(arg2, arg3)
	}

}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) > 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) > 4 {
			arg3 = os.Args[3]
		}

	} else {
		color.Red("Error: Please provide a command")
		showHelp()
		return "", "", "", errors.New("command required")
	}
	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`
	Available commands:
	help		- Show this help
	version		- Show version
	`)
}

func exitGraceFully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	if err != nil {
		color.Red("Error: " + message)
	}
	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}
	os.Exit(0)
}
