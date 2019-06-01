package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/fnkr/cot/config"
)

func main() {
	if config.Debug() {
		config.PrintConfig(os.Stderr)
	}

	if !config.IsInLimit(config.WorkDir()) {
		fmt.Fprintf(os.Stderr, "%s: error: WorkDir(%s) is not in Limit(%s)\n", config.BinName(), config.WorkDir(), config.LimitString())
		os.Exit(1)
	}

	createTmpDirs()
	writePasswdFile()
	writeGroupFile()

	run := getRun()
	cmd := append(config.Tool(), run.ToolCmdAndArgs()...)

	if config.Debug() {
		fmt.Fprintf(os.Stderr, "cmd = %+v\n", cmd)
	}

	if !config.DryRun() {
		if err := syscall.Exec(cmd[0], cmd, os.Environ()); err != nil {
			fmt.Fprintf(os.Stderr, "%s: error: %s\n", config.BinName(), err.Error())
			os.Exit(1)
		}
	}
}
