package main

import (
	"os/exec"
	"strings"
)

func runCommand() string {
	args := []string{"-root", "\t$0", "_NET_ACTIVE_WINDOW"}
	activeProcess, err := exec.Command("xprop", args...).Output()
	if err != nil {
		return ""
	}

	s := strings.Split(string(activeProcess), "#")
	pid := strings.TrimSpace(s[1])

	args1 := []string{"-id ", pid}

	processInfo, err1 := exec.Command("xprop", args1...).Output()
	if err1 != nil {
		return ""
	}

	s1 := strings.Split(string(processInfo), "WM_CLASS(STRING) =")
	s2 := strings.Split(s1[1], "WM_ICON_NAME(STRING)")
	s3 := strings.Split(s2[0], ",")
	return strings.TrimSpace(s3[1])
}
