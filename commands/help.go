package commands

import "fmt"

func commandHelp(_ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}
	fmt.Println("")
	fmt.Println()

	return nil
}
