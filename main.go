package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Create global variable for use in definitions that need it
var commandRegistry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	Lower := strings.ToLower(text)
	split_words := strings.Fields(Lower)
	return split_words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	//This is why I needed that global variable defined but not initialized
	//Otherwise I'd have a loop where commandRegistry needed to be initialized
	//But it couldn't be initialized without this function being defined
	if len(commandRegistry) == 0 {
		fmt.Println("No commands in registry")
		return nil
	}

	for _, cmd := range commandRegistry {
		fmt.Println(cmd.name + ": " + cmd.description)
	}
	return nil
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		cmd, exists := commandRegistry[command]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// special function that automatically, implicitly executes without being called
func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	repl()
}
