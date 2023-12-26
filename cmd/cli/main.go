package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/prateekjoshi2013/scotch"
)

const Version = "1.0.0"

var sco scotch.Scotch

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGraceFully(err)
	}

	setup()

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version: %s\n", Version)
	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGraceFully(err)
		}
		message = "Migration completed successfully"
	case "make":
		if arg2 == "" {
			exitGraceFully(errors.New("make requires a subcommand: (migration | model |handler)"))
		}
		err = doMake(arg2, arg3)
		if err != nil {
			exitGraceFully(err)
		}
	default:
		showHelp()
	}
	exitGraceFully(nil, message)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) > 2 {
			arg2 = os.Args[2]
		}

		if len(os.Args) > 3 {
			arg3 = os.Args[3]
		}

	} else {
		color.Red("Error: Please provide a command")
		showHelp()
		return "", "", "", errors.New("command required")
	}
	color.Red("Command: %s %s %s", arg1, arg2, arg3)
	return arg1, arg2, arg3, nil
}



func exitGraceFully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	if err != nil {
		color.Red("Error: " + err.Error())
	}
	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished")
	}
	os.Exit(0)
}
