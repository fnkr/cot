package config

import (
	"fmt"
	"os"
	userpkg "os/user"
)

var (
	isInitUser bool
	uid        string
	user       string
	gid        string
	group      string
	homeDir    string
)

func initUser() {
	if u, err := userpkg.Current(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: user.Current(): %s", BinName(), err.Error())
		os.Exit(1)
	} else {
		uid = u.Uid
		gid = u.Gid
		user = u.Username
		homeDir = u.HomeDir
	}
	isInitUser = true
}

func UID() string {
	if !isInitUser {
		initUser()
	}

	return uid
}

func GID() string {
	if !isInitUser {
		initUser()
	}

	return gid
}

func User() string {
	if !isInitUser {
		initUser()
	}

	return user
}

func Group() string {
	if !isInitUser {
		initUser()
	}

	return group
}

func HomeDir() string {
	if !isInitUser {
		initUser()
	}

	return homeDir
}
