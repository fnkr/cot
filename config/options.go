package config

import (
	"fmt"
	"os"
	"strings"
)

var (
	isInitEnvPrefix          bool
	envPrefix                string
	isInitImage              bool
	image                    string
	isInitTmp                bool
	tmp                      string
	isInitReadOnlyRoot       bool
	readOnlyRoot             bool
	isInitNet                bool
	net                      string
	isInitTTY                bool
	tty                      bool
	isInitInteractive        bool
	interactive              bool
	isInitLimit              bool
	limit                    []string
	isInitLimitString        bool
	limitString              string
	isInitSSHAuthSock        bool
	sshAuthSock              string
	mountSSHKnownHosts       bool
	isInitMountSSHKnownHosts bool
	isInitShell              bool
	shell                    string
	isInitCPUs               bool
	cpus                     string
	isInitMemory             bool
	memory                   string
	isInitMemoryReservation  bool
	memoryReservation        string
	isInitCapAdd             bool
	capAdd                   []string
	isInitCapDrop            bool
	capDrop                  []string
	isInitEnv                bool
	env                      map[string]string
	isInitVolumes            bool
	volumes                  []string
	isInitDebug              bool
	debug                    bool
	isInitDryRun             bool
	dryRun                   bool
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

func ReadOnlyRoot() bool {
	if !isInitReadOnlyRoot {
		def := true
		if ToolName() == PODMAN {
			def = false
		}
		readOnlyRoot = boolFromEnv(EnvPrefix()+"_READ_ONLY_ROOT", def)
		isInitReadOnlyRoot = true
	}

	return readOnlyRoot
}

func Network() string {
	if !isInitNet {
		net = os.Getenv(EnvPrefix() + "_NET")
		if net == "" {
			if ToolName() == PODMAN {
				net = "slirp4netns"
			} else if ToolName() == DOCKER {
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

func MountSSHKnownHosts() bool {
	if !isInitMountSSHKnownHosts {
		mountSSHKnownHosts = boolFromEnv(EnvPrefix()+"_MOUNT_SSH_KNOWN_HOSTS", true)
		isInitMountSSHKnownHosts = true
	}

	return mountSSHKnownHosts
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
			cpus = CPUsDefault()
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
			memoryReservation = MemoryReservationDefault()
		}
		isInitMemoryReservation = true
	}

	return memoryReservation
}

func Env() map[string]string {
	if !isInitEnv {
		env = listFromEnvs(EnvPrefix() + "_ENV_")
		isInitEnv = true
	}

	return env
}

func Volumes() []string {
	if !isInitVolumes {
		for _, volume := range listFromEnvs(EnvPrefix() + "_VOLUME_") {
			if volume == "" {
				continue
			}
			volumes = append(volumes, volume)
		}
		isInitVolumes = true
	}

	return volumes
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
