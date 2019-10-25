package main

import (
	"os"

	"github.com/fnkr/cot/config"
	"github.com/fnkr/cot/container"
)

func getRun() container.RunCommand {
	volumes := []container.Volume{
		container.Volume{
			HostDir:      config.Tmp() + "/etc/passwd",
			ContainerDir: "/etc/passwd",
			Writable:     config.ToolName() == config.PODMAN,
			SELabel:      false,
		},
		container.Volume{
			HostDir:      config.Tmp() + "/etc/group",
			ContainerDir: "/etc/group",
			Writable:     config.ToolName() == config.PODMAN,
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
		container.Volume{
			HostDir:      config.HomeDir() + "/.ssh/known_hosts",
			ContainerDir: config.HomeDir() + "/.ssh/known_hosts",
			Writable:     false,
			SELabel:      config.SELinuxEnabled(),
		},
		container.Volume{
			HostDir:      config.WorkDir(),
			ContainerDir: config.WorkDir(),
			Writable:     true,
			SELabel:      config.SELinuxEnabled(),
		},
	}

	env := map[string]string{
		"COT_ISOLATED": config.ToolName(),
	}

	if config.SSHAuthSock() != "" {
		containerSSHAuthSock := config.HomeDir() + "/.ssh/auth.sock"
		volumes = append(volumes, container.Volume{
			HostDir:      config.SSHAuthSock(),
			ContainerDir: containerSSHAuthSock,
			Writable:     config.ToolName() == config.PODMAN,
			SELabel:      config.SELinuxEnabled(),
		})
		env["SSH_AUTH_SOCK"] = containerSSHAuthSock
	}

	editor := os.Getenv("EDITOR")
	if editor != "" {
		env["EDITOR"] = editor
	}

	for key, val := range config.Env() {
		env[key] = val
	}

	uidmaps := []container.UIDMap{}

	if config.ToolName() == config.PODMAN {
		uidmaps = append(uidmaps,
			container.UIDMap{
				HostUID:      "0",
				ContainerUID: config.UID(),
				Amount:       "1",
			},
			container.UIDMap{
				HostUID:      "1",
				ContainerUID: "0",
				Amount:       config.UID(),
			},
			// TODO: Fix missing UIDMap
			/*container.UIDMap{
				ContainerUID: "1001",  // config.UID() + 1,
				HostUID:      "1001",  // config.UID() + 1,
				Amount:       "64516", // 65536 - config.UID(),
			},*/
		)
	}

	toolArgs := config.ToolArgs()

	for _, volume := range config.Volumes() {
		toolArgs = append(toolArgs, "--volume="+volume)
	}

	run := container.RunCommand{
		TTY:         config.TTY(),
		Interactive: config.Interactive(),
		CmdAndArgs:  config.CmdAndArgs(),
		Create: container.CreateCommand{
			Image:             config.Image(),
			Rm:                true,
			User:              config.UID() + ":" + config.GID(),
			UIDMaps:           uidmaps,
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
