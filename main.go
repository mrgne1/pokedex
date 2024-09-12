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

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error:\n%v\n", err)
		}
		input = strings.ReplaceAll(input, "\n", "")
		args := strings.Split(input, " ")

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
}

