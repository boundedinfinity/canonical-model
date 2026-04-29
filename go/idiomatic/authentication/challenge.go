package authentication

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type ChallengeQuestion struct {
	Id   ider.Id `json:"id,omitempty"`
	Code string  `json:"code,omitempty"`
	Text string  `json:"text,omitempty"`
}

type ChallengeAnswer struct {
	Id   ider.Id `json:"id,omitempty"`
	Code string  `json:"code,omitempty"`
	Text string  `json:"text,omitempty"`
}

type ChallengeQuestionnaire struct {
	Id        ider.Id           `json:"id,omitempty"`
	Question  ChallengeQuestion `json:"question,omitempty"`
	Correct   ChallengeAnswer   `json:"correct,omitempty"`
	Incorrect []ChallengeAnswer `json:"incorrect,omitempty"`
}
