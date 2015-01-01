package misc

import "github.com/casimir/embargo/eval"

type Output struct{}

func (o Output) Time(layout string) string {
	return formatTime(layout)
}

var Out = eval.New(Output{})
