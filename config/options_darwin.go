// +build darwin

package config

func CPUsDefault() string {
	return ""
}

func MemoryReservationDefault() string {
	return ""
}

func SSHAuthSock() string {
	return "/run/host-services/ssh-auth.sock"
}
