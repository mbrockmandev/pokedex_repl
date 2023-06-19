package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name     string
	desc     string
	callback func() error
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		words := normalizeInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		cmd, ok := commands()[cmdName]

		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func normalizeInput(text string) []string {
	result := strings.ToLower(text)
	words := strings.Fields(result)
	return words
}

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
