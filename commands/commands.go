package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"map": {
			Name:        "map",
			Description: "Display the next twenty map locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "display the previous twenty map locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "list all Pokemon in the area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon's details",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Display all caught pokemon",
			Callback:    commandPokedex,
		},
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}
}
