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
			Writable:     false,
			SELabel:      false,
		},
		container.Volume{
			HostDir:      config.Tmp() + "/etc/group",
			ContainerDir: "/etc/group",
			Writable:     false,
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

	env := map[string]string{}

	if config.SSHAuthSock() != "" {
		containerSSHAuthSock := config.HomeDir() + "/.ssh/auth.sock"
		volumes = append(volumes, container.Volume{
			HostDir:      config.SSHAuthSock(),
			ContainerDir: containerSSHAuthSock,
			SELabel:      config.SELinuxEnabled(),
		})
		env["SSH_AUTH_SOCK"] = containerSSHAuthSock
	}

	env["COT_CONTAINER"] = config.ToolName()
	env["EDITOR"] = os.Getenv("EDITOR")

	uidmaps := []container.UIDMap{}

	if config.ToolName() == "podman" {
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

	run := container.RunCommand{
		TTY:         config.TTY(),
		Interactive: config.Interactive(),
		CmdAndArgs:  config.CmdAndArgs(),
		Create: container.CreateCommand{
			Image:             config.Image(),
			Rm:                true,
			User:              config.UID() + ":" + config.GID(),
			UIDMaps:           uidmaps,
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
		ExtraToolArgs: config.ToolArgs(),
	}

	return run
}
