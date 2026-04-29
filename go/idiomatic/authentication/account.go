package authentication

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type Account struct {
	Id               ider.Id                  `json:"id,omitempty"`
	UsernamePassword UsernamePassword         `json:"username-password,omitempty"`
	Questionnaire    []ChallengeQuestionnaire `json:"security-questions,omitempty"`
}
