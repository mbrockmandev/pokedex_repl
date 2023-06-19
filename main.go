package main

import (
	"bufio"
	"fmt"
	"os"
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
