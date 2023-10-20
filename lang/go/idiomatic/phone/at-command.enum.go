package phone

// https://en.wikipedia.org/wiki/Hayes_AT_command_set#The_basic_Hayes_command_set
// D - Dial mode

//go:generate enumer -path=./at-command.enum.go

type AtCommand string

type atCommands struct {
	Pause2Seconds   AtCommand `json:"pause-2,omitempty" enum:","`
	PauseIndefinite AtCommand `json:"pause-indefinite,omitempty" enum:";"`
}
