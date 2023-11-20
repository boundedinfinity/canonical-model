package modeller

type Contraint func() error

type Validator interface {
	Validate(groups ...string) error
}
