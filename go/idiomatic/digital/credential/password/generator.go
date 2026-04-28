package password

import (
	"fmt"
	"math/rand/v2"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Generator struct {
	Min int
}

func (this Generator) Generate(requirements RequirementsModel) string {
	allowedCharacters := this.calculateAllowed(requirements)
	allowedLength := len(allowedCharacters) - 1
	size := max(this.Min, requirements.Size)
	password := make([]string, size)

	getChar := func(last string) string {
		var found string

		for {
			random := rand.IntN(allowedLength)
			found = allowedCharacters[random]

			if !requirements.NoRepeatedSqences || last == "" || last != found {
				break
			}
		}

		return found
	}

	var last string
	for i := range size {
		password[i] = getChar(last)
		last = password[i]
	}

	return stringer.Join("", password...)
}

func (this Generator) calculateAllowed(requirements RequirementsModel) []string {
	var allowedCharacters []string

	if requirements.Alhpa || requirements.AlhpaLower {
		for i := 'a'; i <= 'z'; i++ {
			allowedCharacters = append(allowedCharacters, fmt.Sprintf("%c", i))
		}
	}

	if requirements.Alhpa || requirements.AlhpaUpper {
		for i := 'A'; i <= 'Z'; i++ {
			allowedCharacters = append(allowedCharacters, fmt.Sprintf("%c", i))
		}
	}

	if requirements.Numeric {
		for i := '0'; i <= '9'; i++ {
			allowedCharacters = append(allowedCharacters, fmt.Sprintf("%c", i))
		}
	}

	if requirements.Symbols {
		allowedCharacters = append(allowedCharacters, "`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", ";", ":", "'", "'", `"`, ",", "<", ">", "?", "/", `\`, "|")
	}

	return allowedCharacters
}
