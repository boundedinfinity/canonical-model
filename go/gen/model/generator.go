package model

type Generator interface {
	Language() Language
	Generate(config Config) error
}
