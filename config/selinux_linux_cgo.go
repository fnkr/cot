// +build linux,cgo

package config

// #cgo linux CFLAGS: -Iinclude -I.
// #cgo pkg-config: libselinux
// #include <selinux/selinux.h>
// #include <selinux/label.h>
// #include <stdlib.h>
// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
import "C"

var (
	isInitSELinuxEnabled bool
	selinuxEnabled       bool
)

func SELinuxEnabled() bool {
	if !isInitSELinuxEnabled {
		selinuxEnabled = C.is_selinux_enabled() > 0
		isInitSELinuxEnabled = true
	}
	return selinuxEnabled
}
