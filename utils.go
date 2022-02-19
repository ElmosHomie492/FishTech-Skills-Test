package main

import (
	"bytes"
	"net"
	"os/exec"
)

func runWhoisCommand(args ...string) bytes.Buffer {
	var out bytes.Buffer

	cmd := exec.Command("whois", args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Run()

	return out
}

func ipValid(ip string) bool {
	ipObj := net.ParseIP(ip)

	if ipObj == nil {
		return false
	} else {
		return true
	}
}
