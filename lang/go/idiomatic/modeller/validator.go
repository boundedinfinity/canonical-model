package modeller

type Contraint func() error

type Validator interface {
	Validate() error
}
