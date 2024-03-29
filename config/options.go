package config

import (
	"fmt"
	"github.com/fnkr/cot/container"
	"os"
	"strings"
)

var (
	isInitImage                     bool
	image                           string
	isInitTmp                       bool
	tmp                             string
	isInitReadOnlyRoot              bool
	readOnlyRoot                    bool
	isInitNet                       bool
	net                             string
	isInitTTY                       bool
	tty                             bool
	isInitInteractive               bool
	interactive                     bool
	isInitLimit                     bool
	limit                           []string
	isInitLimitString               bool
	limitString                     string
	isInitMakeSSHAuthSockAccessible bool
	makeSSHAuthSockAccessible       bool
	mountSSHKnownHosts              bool
	isInitMountSSHKnownHosts        bool
	isInitShell                     bool
	shell                           string
	isInitCPUs                      bool
	cpus                            string
	isInitMemory                    bool
	memory                          string
	isInitMemoryReservation         bool
	memoryReservation               string
	isInitCapAdd                    bool
	capAdd                          []string
	isInitCapDrop                   bool
	capDrop                         []string
	isInitEnv                       bool
	env                             map[string]string
	isInitVolumes                   bool
	volumes                         []container.Volume
	isInitCustomWorkingDirVolume    bool
	customWorkingDirVolume          bool
	isInitDebug                     bool
	debug                           bool
	isInitDryRun                    bool
	dryRun                          bool
)

func EnvPrefix() string {
	return "COT"
}

func Image() string {
	if !isInitImage {
		image = os.Getenv(EnvPrefix() + "_IMAGE")
		if image == "" {
			image = "ghcr.io/fnkr/cot"
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
		if ToolName() == container.PODMAN {
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
			if ToolName() == container.PODMAN {
				net = "slirp4netns"
			} else if ToolName() == container.DOCKER {
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

func MakeSSHAuthSockAccessible() bool {
	if !isInitMakeSSHAuthSockAccessible {
		makeSSHAuthSockAccessible = boolFromEnv(EnvPrefix()+"_MAKE_SSH_AUTH_SOCK_ACCESSIBLE", true)
		isInitMakeSSHAuthSockAccessible = true
	}

	return makeSSHAuthSockAccessible
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

func volumeFromString(volStr string) container.Volume {
	volSlice := strings.Split(volStr, ":")

	hostDir := volSlice[0]
	if len(volSlice) < 2 {
		fmt.Fprintf(os.Stderr, "%s: error: unable to parse volume definition: %s\n", BinName(), volStr)
		os.Exit(1)
	}
	containerDir := volSlice[1]

	var options []string
	if len(volSlice) > 2 {
		options = strings.Split(volSlice[2], ",")
	}

	writable := true

	for _, option := range options {
		switch option {
		case "ro":
			writable = false
		}
	}

	return container.Volume{
		HostDir:      hostDir,
		ContainerDir: containerDir,
		Writable:     writable,
		SELabel:      SELinuxEnabled(),
	}
}

func Volumes() []container.Volume {
	if !isInitVolumes {
		for _, volume := range listFromEnvs(EnvPrefix() + "_VOLUME_") {
			if volume == "" {
				continue
			}
			volumes = append(volumes, volumeFromString(volume))
		}
		isInitVolumes = true
	}

	return volumes
}

func CustomWorkingDirVolume() bool {
	if !isInitCustomWorkingDirVolume {
		for _, volume := range Volumes() {
			if volume.HostDir == WorkDir() {
				customWorkingDirVolume = true
				break
			}
		}
		isInitCustomWorkingDirVolume = true
	}

	return customWorkingDirVolume
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
