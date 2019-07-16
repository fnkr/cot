package config

import (
	"fmt"
	"os"
	userpkg "os/user"
)

var (
	isInitUser    bool
	uid           string
	user          string
	gid           string
	group         string
	inDockerGroup bool
	homeDir       string
)

func initUser() {
	if u, err := userpkg.Current(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: user.Current(): %s", BinName(), err.Error())
		os.Exit(1)
	} else {
		if gids, err := u.GroupIds(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: u.GroupIds(): %s", BinName(), err.Error())
			os.Exit(1)
		} else {
			for _, gid := range gids {
				if g, err := userpkg.LookupGroupId(gid); err != nil {
					fmt.Fprintf(os.Stderr, "%s: user.LookupGroupId(gid): %s", BinName(), err.Error())
					os.Exit(1)
				} else {
					if u.Gid == g.Gid {
						group = g.Name
					}

					if g.Name == "docker" {
						inDockerGroup = true
					}
				}
			}
		}

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

func InDockerGroup() bool {
	if !isInitUser {
		initUser()
	}

	return inDockerGroup
}

func HomeDir() string {
	if !isInitUser {
		initUser()
	}

	return homeDir
}
