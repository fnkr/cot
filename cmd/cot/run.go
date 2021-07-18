package main

import (
	"fmt"
	"os"

	"github.com/fnkr/cot/config"
	"github.com/fnkr/cot/container"
)

func getRun() container.RunCommand {
	volumes := []container.Volume{
		container.Volume{
			HostDir:      config.Tmp() + "/etc/passwd",
			ContainerDir: "/etc/passwd",
			Writable:     config.ToolName() == container.PODMAN,
			SELabel:      false,
		},
		container.Volume{
			HostDir:      config.Tmp() + "/etc/group",
			ContainerDir: "/etc/group",
			Writable:     config.ToolName() == container.PODMAN,
			SELabel:      false,
		},
		container.Volume{
			HostDir:      "/etc/hosts",
			ContainerDir: "/etc/hosts",
			Writable:     false,
			SELabel:      false,
		},
		container.Volume{
			HostDir:      config.Tmp() + "/tmp",
			ContainerDir: "/tmp",
			Writable:     true,
			SELabel:      config.SELinuxEnabled(),
		},
		container.Volume{
			HostDir:      config.Tmp() + "/home",
			ContainerDir: config.HomeDir(),
			Writable:     true,
			SELabel:      config.SELinuxEnabled(),
		},
	}

	if !config.CustomWorkingDirVolume() {
		volumes = append(volumes, container.Volume{
			HostDir:      config.WorkDir(),
			ContainerDir: config.WorkDir(),
			Writable:     true,
			SELabel:      config.SELinuxEnabled(),
		})
	}

	env := map[string]string{
		"COT_ISOLATED": config.ToolName(),
	}

	if config.SSHAuthSock() != "" {
		containerSSHAuthSock := config.HomeDir() + "/.ssh/auth.sock"
		volumes = append(volumes, container.Volume{
			HostDir:      config.SSHAuthSock(),
			ContainerDir: containerSSHAuthSock,
			Writable:     config.ToolName() == container.PODMAN,
			SELabel:      config.SELinuxEnabled(),
		})
		env["SSH_AUTH_SOCK"] = containerSSHAuthSock

		if config.MakeSSHAuthSockAccessible() {
			if err := makeSSHAuthSockAccessible(); err != nil {
				fmt.Fprintf(os.Stderr, "%s: error: %s\n", config.BinName(), err.Error())
				os.Exit(1)
			}
		}
	}

	if config.MountSSHKnownHosts() {
		volumes = append(volumes, container.Volume{
			HostDir:      config.HomeDir() + "/.ssh/known_hosts",
			ContainerDir: config.HomeDir() + "/.ssh/known_hosts",
			Writable:     false,
			SELabel:      config.SELinuxEnabled(),
		})
	}

	if editor := os.Getenv("EDITOR"); editor != "" {
		env["EDITOR"] = editor
	}

	for key, val := range config.Env() {
		env[key] = val
	}

	toolArgs := config.ToolArgs()

	for _, volume := range config.Volumes() {
		volumes = append(volumes, volume)
	}

	groups := []string{}
	for _, group := range config.AddGroupsFinal() {
		groups = append(groups, group.Gid)
	}

	run := container.RunCommand{
		TTY:         config.TTY(),
		Interactive: config.Interactive(),
		CmdAndArgs:  config.CmdAndArgs(),
		Create: container.CreateCommand{
			Image:             config.Image(),
			Rm:                true,
			User:              config.UID() + ":" + config.GID(),
			GroupAdd:          groups,
			ReadOnlyRoot:      config.ReadOnlyRoot(),
			Net:               config.Network(),
			Volumes:           volumes,
			Env:               env,
			WorkDir:           config.WorkDir(),
			CapDrop:           config.CapDrop(),
			CapAdd:            config.CapAdd(),
			CPUs:              config.CPUs(),
			Memory:            config.Memory(),
			MemoryReservation: config.MemoryReservation(),
		},
		ExtraToolArgs: toolArgs,
	}

	return run
}
