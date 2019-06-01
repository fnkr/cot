package config

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	isInitCommand bool
	binName       string
	cmd           string
	args          []string
	isInitWorkDir bool
	workDir       string
	isInitSudo    bool
	sudo          string
	isInitDocker  bool
	docker        string
	isInitPodman  bool
	podman        string
)

func initCommand() {
	binPath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: os.Executable(): %s", os.Args[0], err.Error())
		os.Exit(1)
	}
	binPath, err = filepath.EvalSymlinks(binPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: filepath.EvalSymlinks(binPath): %s", os.Args[0], err.Error())
		os.Exit(1)
	}
	binName = path.Base(binPath)
	callable := path.Base(os.Args[0])

	if binName == callable {
		// cot echo foo
		if len(os.Args) > 1 {
			cmd = os.Args[1]
		}
		if len(os.Args) > 2 {
			args = os.Args[2:]
		}
	} else {
		// echo foo
		cmd = callable
		if len(os.Args) > 1 {
			args = os.Args[1:]
		}
	}
	isInitCommand = true
}

func BinName() string {
	if !isInitCommand {
		initCommand()
	}

	return binName
}

func CmdAndArgs() []string {
	if !isInitCommand {
		initCommand()
	}

	return append([]string{Cmd()}, Args()...)
}

func Cmd() string {
	if !isInitCommand {
		initCommand()
	}

	return cmd
}

func Args() []string {
	if !isInitCommand {
		initCommand()
	}

	return args
}

func WorkDir() string {
	if !isInitWorkDir {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: os.Getwd(): %s\n", BinName(), err.Error())
			os.Exit(1)
		}
		workDir = wd
		isInitWorkDir = true
	}

	return workDir
}

func Sudo() string {
	if !isInitSudo {
		sudo, _ = exec.LookPath("sudo")
		isInitSudo = true
	}

	return sudo
}

func Docker() string {
	if !isInitDocker {
		docker, _ = exec.LookPath("docker")
		isInitDocker = true
	}

	return docker
}

func Podman() string {
	if !isInitPodman {
		podman, _ = exec.LookPath("podman")
		isInitPodman = true
	}

	return podman
}
