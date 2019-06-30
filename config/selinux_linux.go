// +build linux,!cgo

package config

func SELinuxEnabled() bool {
	return false
}
