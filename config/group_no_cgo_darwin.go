// +build !cgo,darwin

package config

import (
	"os/user"
)

func Group() string {
	return "staff"
}

func InDockerGroup() bool {
	return true
}

func Groups() []user.Group {
	return []user.Group{}
}
