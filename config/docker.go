package config

import (
	"golang.org/x/sys/unix"
)

const (
	DOCKER_SOCKET_PATH = "/var/run/docker.sock"
)

var (
	isInitDockerSocketIsWritable bool
	dockerSocketIsWritable       bool
)

func DockerSocketIsWritable() bool {
	if !isInitDockerSocketIsWritable {
		dockerSocketIsWritable = unix.Access(DOCKER_SOCKET_PATH, unix.W_OK) == nil
		isInitDockerSocketIsWritable = true
	}

	return dockerSocketIsWritable
}
