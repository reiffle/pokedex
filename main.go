package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Create global variable for use in definitions that need it
var commandRegistry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}
type Config struct {
	next     string
	previous string
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	Lower := strings.ToLower(text)
	split_words := strings.Fields(Lower)
	return split_words
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
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

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.next != "" { //Works because the first time config is called, config.next is an empty string
		url = config.next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	rep, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var areas LocationAreaResponse //Note, this is just a single struct, not a slice

	if err = json.Unmarshal(rep, &areas); err != nil {
		return err
	}

	config.next = areas.Next
	config.previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	url := config.previous
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	rep, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var areas LocationAreaResponse

	if err = json.Unmarshal(rep, &areas); err != nil {
		return err
	}
	config.next = areas.Next
	config.previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{}
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
			err := cmd.callback(config)
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

		"map": {
			name:        "map",
			description: "Displays next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func main() {
	repl()
}
