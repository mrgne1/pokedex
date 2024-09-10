package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
