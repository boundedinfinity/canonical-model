package model

type jsonInterface interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
