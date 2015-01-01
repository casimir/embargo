package color

import (
	"log"

	"github.com/casimir/embargo/eval"
)

const (
	FormatNone = "none"
	FormatDzen = "dzen2"
	FormatTerm = "term"
)

type None struct {
	End string
}

func (n None) Begin(color string) string { return "" }

type Output struct {
	module eval.Evaluator
}

func (o Output) Eval(data ...string) string {
	return o.module.Eval(data...)
}

var Out = &Output{module: eval.New(None{})}

func Load(format string) {
	switch format {
	case FormatDzen:
		log.Println("Color mode: Dzen")
		Out.module = eval.New(newDzen())
	case FormatTerm:
		log.Println("Color mode: ANSI")
		Out.module = eval.New(newTerm())
	default:
		Out.module = eval.New(None{})
	}
}
