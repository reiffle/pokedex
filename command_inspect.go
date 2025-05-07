package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must name a pokemon to inspect it")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	//check if pokemon has been caught
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, value := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, value := range pokemon.Types {
		fmt.Printf("  - %s\n", value.Type.Name)
	}

	return nil
}
