// +build cgo

package config

import (
	"fmt"
	"os"
	"os/user"
)

var (
	isInitGroup   bool
	group         string
	inDockerGroup bool
)

func initGroup() {
	u, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: user.Current(): %s", BinName(), err.Error())
		os.Exit(1)
	}

	gids, err := u.GroupIds()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: u.GroupIds(): %s", BinName(), err.Error())
		os.Exit(1)
	}

	for _, gid := range gids {
		g, err := user.LookupGroupId(gid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: user.LookupGroupId(gid): %s", BinName(), err.Error())
			os.Exit(1)
		}

		if u.Gid == g.Gid {
			group = g.Name
		}

		if g.Name == "docker" {
			inDockerGroup = true
		}
	}

	isInitGroup = true
}

func Group() string {
	if !isInitGroup {
		initGroup()
	}

	return group
}

func InDockerGroup() bool {
	if !isInitGroup {
		initGroup()
	}

	return inDockerGroup
}
