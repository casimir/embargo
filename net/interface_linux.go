package net

import (
	"os/exec"
	"strings"
)

func ssid(name string) string {
	cmd := exec.Command("iwgetid", "-r", name)
	out, err := cmd.Output()
	if err != nil {
		return "No SSID"
	}
	return strings.Trim(string(out), "\n")
}
