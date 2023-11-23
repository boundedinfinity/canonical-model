package authentication

import "github.com/boundedinfinity/schema/idiomatic/id"

type ChallengeQuestion struct {
	Id   id.Id  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Text string `json:"text,omitempty"`
}

type ChallengeAnswer struct {
	Id   id.Id  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Text string `json:"text,omitempty"`
}

type ChallengeQuestionnaire struct {
	Id        id.Id             `json:"id,omitempty"`
	Question  ChallengeQuestion `json:"question,omitempty"`
	Correct   ChallengeAnswer   `json:"correct,omitempty"`
	Incorrect []ChallengeAnswer `json:"incorrect,omitempty"`
}
