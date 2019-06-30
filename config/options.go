package config

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	isInitEnvPrefix         bool
	envPrefix               string
	isInitImage             bool
	image                   string
	isInitTmp               bool
	tmp                     string
	isInitNet               bool
	net                     string
	isInitTTY               bool
	tty                     bool
	isInitInteractive       bool
	interactive             bool
	isInitLimit             bool
	limit                   []string
	isInitLimitString       bool
	limitString             string
	isInitSSHAuthSock       bool
	sshAuthSock             string
	isInitShell             bool
	shell                   string
	isInitDebug             bool
	cpus                    string
	isInitCPUs              bool
	memory                  string
	isInitMemory            bool
	memoryReservation       string
	isInitMemoryReservation bool
	capAdd                  []string
	isInitCapAdd            bool
	capDrop                 []string
	isInitCapDrop           bool
	debug                   bool
	isInitDryRun            bool
	dryRun                  bool
)

func EnvPrefix() string {
	if !isInitEnvPrefix {
		envPrefix = strings.ToUpper(BinName())
		isInitEnvPrefix = true
	}

	return envPrefix
}

func Image() string {
	if !isInitImage {
		image = os.Getenv(EnvPrefix() + "_IMAGE")
		if image == "" {
			image = "fnkr/cot"
		}
		isInitImage = true
	}

	return image
}

func Tmp() string {
	if !isInitTmp {
		tmp = os.Getenv(EnvPrefix() + "_TMP")
		if tmp == "" {
			tmp = fmt.Sprintf("/tmp/%s-%s-%s", BinName(), ToolName(), UID())
		}
		isInitTmp = true
	}

	return tmp
}

func Network() string {
	if !isInitNet {
		net = os.Getenv(EnvPrefix() + "_NET")
		if net == "" {
			if ToolName() == "podman" {
				net = "slirp4netns"
			} else if ToolName() == "docker" {
				net = "bridge"
			} else {
				fmt.Fprintf(os.Stderr, "%s: error: not implemented: ToolName(%s)\n", BinName(), ToolName())
				os.Exit(1)
			}
		}
		isInitNet = true
	}

	return net
}

func TTY() bool {
	if !isInitTTY {
		tty = boolFromEnv(EnvPrefix()+"_TTY", true)
		isInitTTY = true
	}

	return tty
}

func Interactive() bool {
	if !isInitInteractive {
		interactive = boolFromEnv(EnvPrefix()+"_INTERACTIVE", true)
		isInitInteractive = true
	}

	return interactive
}

func Limit() []string {
	if !isInitLimit {
		limit = listFromEnv(EnvPrefix()+"_LIMIT", ":")
		isInitLimit = true
	}

	return limit
}

func LimitString() string {
	if !isInitLimitString {
		limitString = strings.Join(Limit(), ":")
		isInitLimitString = true
	}

	return limitString
}

func SSHAuthSock() string {
	if !isInitSSHAuthSock {
		sshAuthSock = os.Getenv("SSH_AUTH_SOCK")
		isInitSSHAuthSock = true
	}

	return sshAuthSock
}

func Shell() string {
	if !isInitShell {
		shell = os.Getenv(EnvPrefix() + "_SHELL")
		if shell == "" {
			shell = "/bin/sh"
		}
		isInitShell = true
	}

	return shell
}

func CPUs() string {
	if !isInitCPUs {
		cpus = os.Getenv(EnvPrefix() + "_CPUS")
		if cpus == "" {
			if ToolName() == "docker" {
				cpus = strconv.FormatFloat(float64(runtime.NumCPU())/1.25, 'f', 6, 64) // 80%
			}
		}
		isInitCPUs = true
	}

	return cpus
}

func Memory() string {
	if !isInitMemory {
		memory = os.Getenv(EnvPrefix() + "_MEMORY")
		isInitMemory = true
	}

	return memory
}

func CapAdd() []string {
	if !isInitCapAdd {
		capAdd = listFromEnv(EnvPrefix()+"_CAP_ADD", ",")
		isInitCapAdd = true
	}
	return capAdd
}

func CapDrop() []string {
	if !isInitCapDrop {
		capDrop = listFromEnv(EnvPrefix()+"_CAP_DROP", ",")
		if len(capDrop) == 0 {
			capDrop = []string{"ALL"}
		}
		isInitCapDrop = true
	}
	return capDrop
}

func MemoryReservation() string {
	if !isInitMemoryReservation {
		memoryReservation = os.Getenv(EnvPrefix() + "_MEMORY_RESERVATION")
		if memoryReservation == "" {
			if ToolName() == "docker" {
				memoryReservation = "1g"
			}
		}
		isInitMemoryReservation = true
	}

	return memoryReservation
}

func Debug() bool {
	if !isInitDebug {
		debug = boolFromEnv(EnvPrefix()+"_DEBUG", false)
		isInitDebug = true
	}

	return debug
}

func DryRun() bool {
	if !isInitDryRun {
		dryRun = boolFromEnv(EnvPrefix()+"_DRY_RUN", false)
		isInitDryRun = true
	}

	return dryRun
}
