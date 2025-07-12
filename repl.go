package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	var args []string
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			log.Println(err)
		}
	}
}

type clicommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"map": {
			name:        "map",
			description: "List next location areas",
			callback:    callBackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "go back to previous map",
			callback:    callBackMapb,
		},
		"explore": {
			name:        "explore {location area name}",
			description: "explore pokemon in area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon area name}",
			description: "catch pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "inspect catched pokemon",
			callback:    callbackInspect,
		},
	}

}

func cleanInput(str string) []string {
	lowerd := strings.ToLower(str)
	words := strings.Fields(lowerd)
	return words
}
