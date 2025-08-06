package authentication

import "github.com/boundedinfinity/schema/idiomatic/id"

type Account struct {
	Id               id.Id                    `json:"id,omitempty"`
	UsernamePassword UsernamePassword         `json:"username-password,omitempty"`
	Questionnaire    []ChallengeQuestionnaire `json:"security-questions,omitempty"`
}
