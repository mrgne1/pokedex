package main

import (
	"fmt"
	"pokedex/commands"
	"strings"
)

func main() {
	var input string
	cmds := commands.GetCommands()
	for {
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)
		parsedInput := strings.Split(input, " ")
		command := strings.ToLower(parsedInput[0])
		c, ok := cmds[command]
		if ok {
			err := c.Callback()
			if err != nil {
				fmt.Printf("**Err**\n%v\n**End**\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %v\n", command)
		}
	}
}

