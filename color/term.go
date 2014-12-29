package color

import "fmt"

const (
	Reset = 0
)

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func fg(c int) int { return c + 30 }

var colors = map[string]int{
	"black":   fg(Black),
	"red":     fg(Red),
	"green":   fg(Green),
	"yellow":  fg(Yellow),
	"blue":    fg(Blue),
	"magenta": fg(Magenta),
	"cyan":    fg(Cyan),
	"white":   fg(White),
}

type Term struct {
	End string
}

func (t Term) Begin(color string) string {
	code, ok := colors[color]
	if !ok {
		return ""
	}
	return fmt.Sprintf("[%dm", code)
}

func newTerm() Term {
	return Term{End: fmt.Sprintf("[%dm", Reset)}
}
