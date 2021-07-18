// +build linux

package config

import (
	"github.com/fnkr/cot/container"
	"os"
	"runtime"
	"strconv"
)

var (
	isInitSSHAuthSock        bool
	sshAuthSock              string
)

func CPUsDefault() string {
	if ToolName() == container.DOCKER {
		return strconv.FormatFloat(float64(runtime.NumCPU())/1.25, 'f', 6, 64) // 80%
	}
	return ""
}

func MemoryReservationDefault() string {
	if ToolName() == container.DOCKER {
		return "1g"
	}
	return ""
}

func SSHAuthSock() string {
	if !isInitSSHAuthSock {
		sshAuthSock = os.Getenv("SSH_AUTH_SOCK")
		isInitSSHAuthSock = true
	}

	return sshAuthSock
}
