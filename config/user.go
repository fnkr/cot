package config

import (
	"fmt"
	"os"
	"os/user"
)

var (
	isInitUser bool
	uid        string
	gid        string
	username   string
	homeDir    string
)

func initUser() {
	u, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: user.Current(): %s", BinName(), err.Error())
		os.Exit(1)
	}

	uid = u.Uid
	gid = u.Gid
	username = u.Username
	homeDir = u.HomeDir
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

	return username
}

func HomeDir() string {
	if !isInitUser {
		initUser()
	}

	return homeDir
}
