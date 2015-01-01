package proc

import "github.com/casimir/embargo/eval"

type Output struct{}

func (o Output) Load(t string) string {
	return load(t)
}

var Out = eval.New(Output{})
