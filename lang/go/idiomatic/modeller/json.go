package modeller

type jsonInterface interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

func (t *modellerBuilder[T]) Json() jsonInterface {

	return nil
}
