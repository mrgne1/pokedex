package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"map": {
			Name: "map",
			Description: "Display the next twenty map locations",
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "display the previous twenty map locations",
			Callback: commandMapb,
		},
		"help": {
			Name: "help",
			Description: "Display a help message",
			Callback: commandHelp,
		},
		"exit": {
			Name: "exit",
			Description: "Exit the pokedex",
			Callback: commandExit,
		}, 
	}
}
