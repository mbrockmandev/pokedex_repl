package main

import (
	"fmt"
	"os"
)

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "This displays a helpful message.",
			callback: cmdHelp,
		},
		"exit": {
			name:     "exit",
			desc:     "Quits the PokeDex REPL",
			callback: cmdExit,
		},
		"map": {
			name:     "map",
			desc:     "Displays the next 20 locations in the Pokemon world.",
			callback: cmdMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Displays the previous 20 locations in the Pokemon world.",
			callback: cmdMapb,
		},
	}
}

// cb functions
func cmdHelp() error {
	fmt.Println("...")
	fmt.Println("POKEDEX HELP MENU")
	fmt.Println("List of Commands:")
	fmt.Println("...")
	for _, cmd := range commands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.desc)
	}
	fmt.Println("...")
	return nil
}

func cmdExit() error {
	os.Exit(0)
	return nil
}

func cmdMap() error {

	return nil
}

func cmdMapb() error {
	
	return nil
}
