package usa

var Forms = forms{}

type forms struct{}

func (t forms) F1040() Form1040 {
	return Form1040{}
}

func (t forms) F1040_V() Form1040_V {
	return Form1040_V{}
}

func (t forms) W2() FormW2 {
	return FormW2{}
}

func (t forms) W4() FormW4 {
	return FormW4{}
}
