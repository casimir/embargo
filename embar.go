package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/casimir/embargo/color"
	"github.com/casimir/embargo/eval"
	"github.com/casimir/embargo/misc"
	"github.com/casimir/embargo/net"
	"github.com/casimir/embargo/proc"
	"gopkg.in/alecthomas/kingpin.v1"
)

// TODO
// proc.disk   → df
// proc.CPU    → /proc/cpuinfo || ps aux (load in '%'?)
// proc.mem    → /proc/meminfo
// proc.uptime → /proc/uptime
// proc.heavy  → glances ?
// misc.music  → github.com/lann/mpris2-go

var (
	aFormat   = kingpin.Flag("format", "Format type of the output.").Short('f').Default(color.FormatNone).Enum(color.FormatNone, color.FormatDzen, color.FormatTerm)
	aInterval = kingpin.Flag("interval", "Seconds between refresh.").Short('i').Default("1s").Duration()
	aPrint    = kingpin.Flag("print", "Print one line and quit.").Short('p').Bool()
	aLine     = kingpin.Arg("line", "Line to evaluate.").Default(defaultLine()).String()
)

func init() {
	eval.DefaultModule = "misc"

	eval.Register("color", color.Out)
	eval.Register("net", net.Out)
	eval.Register("proc", proc.Out)
	eval.Register(eval.DefaultModule, misc.Out)
}

func defaultLine() string {
	sep := "${color.begin blue} · ${color.end}"
	return strings.Join([]string{
		"${color.begin green}${net.wlo1 ssid} ${net.wlo1 ip}${color.end}",
		"${proc.load 1}",
		"${time 'Mon _2'}",
		"${time '15:04'}",
	}, sep)
}

func main() {
	kingpin.Parse()
	color.Load(*aFormat)
	if *aPrint {
		fmt.Println(eval.Eval(*aLine))
	} else {
		for {
			fmt.Println(eval.Eval(*aLine))
			time.Sleep(*aInterval)
		}
	}
}
