package phone

// https://en.wikipedia.org/wiki/Hayes_AT_command_set#The_basic_Hayes_command_set
// D - Dial mode

type AtCommand string

type atCommands struct {
	Pause2          AtCommand `json:"pause-2,omitempty" enum:","`
	PauseIndefinite AtCommand `json:"pause-indefinite,omitempty" enum:";"`
}
