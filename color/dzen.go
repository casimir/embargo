package color

type Dzen struct {
	End string
}

func (d Dzen) Begin(color string) string {
	return "^fg(" + color + ")"
}

func newDzen() Dzen {
	return Dzen{End: "^fg()"}
}
