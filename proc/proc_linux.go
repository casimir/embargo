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
	ret := parts[0]

	switch t {
	case Load1:
		ret = parts[0]
	case Load5:
		ret = parts[1]
	case Load15:
		ret = parts[2]
	}
	// TODO add critical color when > n + runtime.NumCPU
	return ret
}
