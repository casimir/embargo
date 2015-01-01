package net

import (
	"log"

	"github.com/casimir/embargo/eval"
)

type Output map[string]eval.Evaluator

func (o Output) Eval(data ...string) string {
	if len(data) < 1 {
		return ""
	}
	idx := data[0]
	i, ok := Out[idx]
	if !ok {
		log.Fatal("net: Unknown interface: ", idx)
	}
	return i.Eval(data[1:]...)
}

var Out = Output{}

func init() {
	il, err := NewInterfaces()
	if err != nil {
		log.Fatalln(err)
	}
	for _, it := range il {
		Out[it.Name] = eval.New(it)
	}
}
