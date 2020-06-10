package main

import (
	"os/exec"
	"regexp"
	"strings"
)

// this function is linux specific which return active window name on the desktop
func getLinuxActiveWindowName() string {
	var (
		pid     string
		process string
	)
	// xprop command is used to fetch the active window process id
	cmd1Args := []string{"-root", "\t$0", "_NET_ACTIVE_WINDOW"}
	activeProcess, cmd1Err := exec.Command("xprop", cmd1Args...).Output()
	if cmd1Err != nil {
		return process
	}
	// parsing the output of xprop line by line
	activeProcessOutput := strings.Split(string(activeProcess), "\n")

	// process id starts with 0x hence use it in reqex to fetch process id
	r, _ := regexp.Compile("0x")
	for _, ln := range activeProcessOutput {

		if len(r.FindString(ln)) > 0 {
			fetchPid := strings.Split(ln, "0x")
			pid = "0x" + strings.TrimSpace(fetchPid[1])
			break
		}

	}
	// if fail to fetch process id return blank
	if pid == "" {
		return process
	}

	// get the entire information of process by process id
	cmd2Args := []string{"-id", pid}
	processInfo, cmd2Err := exec.Command("xprop", cmd2Args...).Output()
	if cmd2Err != nil {
		return process
	}
	// Parse the output of above command line by line
	processInfoOutput := strings.Split(string(processInfo), "\n")
	reg, _ := regexp.Compile("WM_CLASS\\(STRING\\)")
	for _, ln := range processInfoOutput {

		if len(reg.FindString(ln)) > 0 {
			answer := strings.Split(ln, "=")
			names := strings.Split(answer[1], ",")
			// assign process name
			process = names[len(names)-1]
			break
		}

	}
	return process
}
