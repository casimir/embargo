package proc

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	Load1  = "1"
	Load5  = "5"
	Load15 = "15"
)

func load(t string) string {
	raw, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		log.Print("Failed to get load: ", err)
		return "no load"
	}
	parts := strings.Fields(string(raw))

	switch t {
	case Load1:
		return parts[0]
	case Load5:
		return parts[1]
	case Load15:
		return parts[2]
	default:
		return parts[0]
	}
}
