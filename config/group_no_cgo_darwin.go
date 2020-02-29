// +build !cgo,darwin

package config

func Group() string {
	return "staff"
}

func InDockerGroup() bool {
	return true
}
