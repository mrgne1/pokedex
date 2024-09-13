package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/commands"
	"strings"
)

func main() {
	cmds := commands.GetCommands()

	// debugRun(cmds)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error:\n%v\n", err)
		}
		input = strings.ReplaceAll(input, "\n", "")
		args := strings.Split(input, " ")

		dispatch(args, cmds)
	}
}

func dispatch(args []string, cmds map[string]commands.CliCommand) {
	command := args[0]
	c, ok := cmds[command]
	if ok {
		err := c.Callback(args[1:]...)
		if err != nil {
			fmt.Printf("**Err**\n%v\n**End**\n", err)
		}
	} else {
		fmt.Printf("Unknown command: %v\n", command)
	}

}
func debugRun(cmds map[string]commands.CliCommand) {
	testArgs := [][]string{
		{"explore", "canalave-city-area"},
		{"catch", "tentacool"},
		{"catch", "tentacool"},
		{"catch", "tentacool"},
		{"catch", "tentacool"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"catch", "pelipper"},
		{"pokedex"},
		{"exit"},
	}

	for _, args := range testArgs {
		dispatch(args, cmds)
	}
}
