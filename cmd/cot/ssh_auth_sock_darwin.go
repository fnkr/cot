// +build darwin

package main

import (
	"fmt"
	"github.com/fnkr/cot/container"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"

	"github.com/fnkr/cot/config"
)

func makeSSHAuthSockAccessible() error {
	testFile := path.Join(config.Tmp(), "ssh_auth_sock_accessible")
	testStat, testErr := os.Stat(testFile)
	if testErr != nil && !os.IsNotExist(testErr) {
		return testErr
	}

	sockFile, err := filepath.EvalSymlinks(config.DOCKER_SOCKET_PATH)
	if err != nil {
		return err
	}
	sockStat, err := os.Stat(sockFile)
	if err != nil {
		return err
	}

	if testErr == nil && !sockStat.ModTime().After(testStat.ModTime()) {
		return nil
	}

	run := container.RunCommand{
		CmdAndArgs: []string{"chmod", "a+w", "/run/ssh-auth.sock"},
		Create: container.CreateCommand{
			Image:        "busybox",
			Rm:           true,
			ReadOnlyRoot: true,
			Net:          "none",
			Volumes: []container.Volume{
				container.Volume{
					HostDir:      config.SSHAuthSock(),
					ContainerDir: "/run/ssh-auth.sock",
					Writable:     true,
				},
			},
		},
	}
	cmd := append(config.Tool(), run.ToolCmdAndArgs(config.ToolName())...)

	if config.Debug() {
		fmt.Fprintf(os.Stderr, "cmd = %+v\n", cmd)
	}

	if !config.DryRun() {
		if _, err := exec.Command(cmd[0], cmd[1:]...).Output(); err != nil {
			return err
		}
	}

	if testErr != nil {
		ioutil.WriteFile(testFile, []byte(""), 0644)
	} else {
		now := time.Now()
		os.Chtimes(testFile, now, now)
	}

	return nil
}
