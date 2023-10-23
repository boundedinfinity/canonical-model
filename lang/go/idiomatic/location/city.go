package location

type City string

func (t City) String() string {
	return string(t)
}
