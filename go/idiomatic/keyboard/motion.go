package keyboard

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type Motion struct {
	Key     []*Key
	Name    string
	Action  string
	Options []*Motion
}

func match(target string) func(*Key) bool {
	return func(key *Key) bool {
		return key.Match(target)
	}
}

func (this Motion) Match(target string) bool {
	return slicer.ContainsFunc(match(target), this.Key...)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var Motions = motions{
	Motion: &Motion{
		Name:   "Motion",
		Action: "motion",
	},
	MoveUp: &Motion{
		Name:   "Move Up",
		Action: "move up",
		Key:    []*Key{Keys.UpperI, Keys.LowerI, Keys.ArrowUp},
	},
	MoveDown: &Motion{
		Name:   "Move Down",
		Action: "move down",
		Key:    []*Key{Keys.UpperK, Keys.LowerK, Keys.ArrowDown},
	},
	MoveLeft: &Motion{
		Name:   "Move Left",
		Action: "move left",
		Key:    []*Key{Keys.UpperJ, Keys.LowerJ, Keys.ArrowLeft},
	},
	MoveRight: &Motion{
		Name:   "Move Right",
		Action: "move right",
		Key:    []*Key{Keys.UpperL, Keys.LowerL, Keys.ArrowRight},
	},
	Back: &Motion{
		Name:   "Back",
		Action: "back",
		Key:    []*Key{Keys.UpperB, Keys.LowerB},
	},
}

type motions struct {
	Motion    *Motion
	MoveUp    *Motion
	MoveDown  *Motion
	MoveLeft  *Motion
	MoveRight *Motion
	Back      *Motion
}

func (this motions) Copy(motion *Motion) *Motion {
	copy := &Motion{
		Key:     motion.Key,
		Name:    motion.Name,
		Action:  motion.Action,
		Options: []*Motion{},
	}

	return copy
}

func (this motions) New(motion *Motion, options ...*Motion) *Motion {
	newMotion := this.Copy(motion)

	for _, option := range options {
		newOption := this.Copy(option)
		newOption.Options = append(newOption.Options, newMotion)
		newMotion.Options = append(newMotion.Options, newOption)
	}

	return newMotion
}
