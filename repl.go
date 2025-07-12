package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		command.callback()
	}
}

type clicommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			name:        "help",
			description: "Prints Help the Command Helps",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the pokedex",
			callback:    callbackExit,
		},
	}

}

func cleanInput(str string) []string {
	lowerd := strings.ToLower(str)
	words := strings.Fields(lowerd)
	return words
}
