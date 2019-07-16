package container

import (
	"github.com/fnkr/cot/config"
	"strconv"
	"strings"
)

type RunCommand struct {
	TTY           bool
	Interactive   bool
	Create        CreateCommand
	CmdAndArgs    []string
	ExtraToolArgs []string
}

type CreateCommand struct {
	Image             string
	Rm                bool
	User              string
	UIDMaps           []UIDMap
	ReadOnlyRoot      bool
	Net               string
	WorkDir           string
	Volumes           []Volume
	Env               map[string]string
	CapDrop           []string
	CapAdd            []string
	CPUs              string
	Memory            string
	MemoryReservation string
}

type Volume struct {
	HostDir      string
	ContainerDir string
	Writable     bool
	SELabel      bool
}

type UIDMap struct {
	HostUID      string
	ContainerUID string
	Amount       string
}

func (rc *RunCommand) ToolCmdAndArgs() []string {
	return append([]string{"run"}, rc.ToolArgs()...)
}

func (rc *RunCommand) ToolArgs() (args []string) {
	args = append(args,
		"--tty="+strconv.FormatBool(rc.TTY),
		"--interactive="+strconv.FormatBool(rc.Interactive),
		"--rm="+strconv.FormatBool(rc.Create.Rm),
		"--read-only="+strconv.FormatBool(rc.Create.ReadOnlyRoot),
	)

	if rc.Create.User != "" {
		args = append(args, "--user="+rc.Create.User)
	}

	for _, uidmap := range rc.Create.UIDMaps {
		args = append(args, uidmap.ToolArg())
	}

	if rc.Create.Net != "" {
		args = append(args, "--net="+rc.Create.Net)
	}

	if rc.Create.WorkDir != "" {
		args = append(args, "--workdir="+rc.Create.WorkDir)
	}

	for _, volume := range rc.Create.Volumes {
		args = append(args, volume.ToolArg())
	}

	for key, val := range rc.Create.Env {
		args = append(args, "--env="+key+"="+val)
	}

	for _, val := range rc.Create.CapDrop {
		args = append(args, "--cap-drop="+val)
	}

	for _, val := range rc.Create.CapAdd {
		args = append(args, "--cap-add="+val)
	}

	if rc.Create.CPUs != "" {
		args = append(args, "--cpus="+rc.Create.CPUs)
	}

	if rc.Create.Memory != "" {
		args = append(args, "--memory="+rc.Create.Memory)
	}

	if rc.Create.MemoryReservation != "" {
		args = append(args, "--memory-reservation="+rc.Create.MemoryReservation)
	}

	args = append(args, rc.ExtraToolArgs...)
	args = append(args, "--", config.Image())
	args = append(args, rc.CmdAndArgs...)

	return
}

func (vol *Volume) ToolArg() string {
	arg := "--volume=" + vol.HostDir + ":" + vol.ContainerDir
	var opts []string

	if vol.Writable {
		opts = append(opts, "rw")
	} else {
		opts = append(opts, "ro")
	}

	if vol.SELabel {
		opts = append(opts, "z")
	}

	return arg + ":" + strings.Join(opts, ",")
}

func (uidmap *UIDMap) ToolArg() string {
	return "--uidmap=" + uidmap.ContainerUID + ":" + uidmap.HostUID + ":" + uidmap.Amount
}
