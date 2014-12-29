package net

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

type (
	InterfaceStatus string
	InterfaceType   string

	Interface struct {
		Ip     string
		Name   string
		Ssid   string
		Status InterfaceStatus
		Type   InterfaceType
	}
)

const (
	StatusDown InterfaceStatus = "down"
	StatusUp   InterfaceStatus = "up"

	TypeEthernet InterfaceType = "eth"
	TypeLoop     InterfaceType = "lo"
	TypeWifi     InterfaceType = "wlan"
)

func newInterface(info net.Interface) (i *Interface, err error) {
	i = new(Interface)
	if uint(info.Flags)&1 == 0 {
		i.Status = StatusDown
		return
	}

	i.Name = info.Name
	i.Ssid = ssid(info.Name)
	i.Status = StatusUp
	i.Type = matchType(info.Name)

	la, err := info.Addrs()
	if err != nil {
		return
	}
	if len(la) > 0 {
		i.Ip = strings.Split(la[0].String(), "/")[0]
	} else {
		i.Ip = "None"
	}
	return
}

func NewInterfaces() (li []*Interface, err error) {
	list, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, it := range list {
		i, err := newInterface(it)
		if err != nil {
			log.Println("Net:interface:", err)
		}
		li = append(li, i)
	}
	return
}

func (i *Interface) String() string {
	if i.Status == StatusDown {
		return fmt.Sprintf("%s: %s", i.Name, i.Status)
	}
	ret := fmt.Sprintf("%s: %s", i.Name, i.Ip)
	if i.Type == TypeWifi {
		ret += fmt.Sprintf(" (%s)", i.Ssid)
	}
	return ret
}

func matchType(name string) InterfaceType {
	switch name[0] {
	case 'e':
		return TypeEthernet
	case 'l':
		return TypeLoop
	case 'w':
		return TypeWifi
	default:
		return "Unknown"
	}
}

func ssid(name string) string {
	cmd := exec.Command("iwgetid", "-r", name)
	out, err := cmd.Output()
	if err != nil {
		return "No SSID"
	}
	return strings.Trim(string(out), "\n")
}
