package main

import (
	"os/exec"
	"regexp"
	"strings"
)

func runCommand() string {
	var pid string = ""
	var process string = ""
	args := []string{"-root", "\t$0", "_NET_ACTIVE_WINDOW"}
	activeProcess, err := exec.Command("xprop", args...).Output()
	if err != nil {
		return ""
	}

	activeProcessOutput := strings.Split(string(activeProcess), "\n")
	r, _ := regexp.Compile("0x")
	for _, ln := range activeProcessOutput {
		if len(r.FindString(ln)) > 0 {
			fetchPid := strings.Split(ln, "0x")
			pid = "0x" + strings.TrimSpace(fetchPid[1])
			break
		}

	}
	if pid == "" {
		return ""
	}
	args1 := []string{"-id", pid}

	processInfo, err1 := exec.Command("xprop", args1...).Output()
	if err1 != nil {
		return ""
	}

	processInfoOutput := strings.Split(string(processInfo), "\n")

	reg, _ := regexp.Compile("WM_CLASS\\(STRING\\)")
	for _, ln := range processInfoOutput {

		if len(reg.FindString(ln)) > 0 {

			answer := strings.Split(ln, "=")
			names := strings.Split(answer[1], ",")
			process = names[len(names)-1]
			break
		}

	}
	return process
}
