package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/casimir/embargo/color"
	"github.com/casimir/embargo/eval"
	"github.com/casimir/embargo/misc"
	"github.com/casimir/embargo/net"
)

// TODO
// proc.disk   → df
// proc.CPU    → /proc/cpuinfo || ps aux (load in '%'?)
// proc.mem    → /proc/meminfo
// proc.load   → /proc/loadavg || uptime
// proc.uptime → uptime
// proc.heavy  → glances ?
// misc.music  → github.com/lann/mpris2-go

func init() {
	eval.DefaultModule = "misc"

	eval.Register("color", color.Out)
	eval.Register("net", net.Out)
	eval.Register(eval.DefaultModule, misc.Out)
}

func main() {
	color.Load(color.FormatDzen)

	sep := "${color.begin blue}|${color.end}"
	line := strings.Join([]string{
		"${color.begin green}${net.wlo1 ssid} ${net.wlo1 ip}${color.end}",
		"${time 'Mon _2'}",
		"${time '15:04'}",
	}, sep)
	for {
		fmt.Println(eval.Eval(line))
		time.Sleep(time.Second)
	}
}
