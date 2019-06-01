package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fnkr/cot/config"
	"github.com/fnkr/cot/template/group"
	"github.com/fnkr/cot/template/passwd"
)

func writePasswdFile() {
	path := filepath.Join(config.Tmp(), "etc", "passwd")

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", config.BinName(), err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: warning: %s\n", config.BinName(), err.Error())
		}
	}()

	if err := passwd.Write(passwd.File{
		Users: []passwd.User{
			passwd.User{
				Name:  "root",
				UID:   "0",
				GID:   "0",
				Home:  "/root",
				Shell: "/bin/sh",
			},
			passwd.User{
				Name:  "nobody",
				UID:   "65534",
				GID:   "65534",
				Home:  "/",
				Shell: "/sbin/nologin",
			},
			passwd.User{
				Name:  config.User(),
				UID:   config.UID(),
				GID:   config.GID(),
				Home:  config.HomeDir(),
				Shell: config.Shell(),
			},
		},
	}, file); err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", config.BinName(), err.Error())
		os.Exit(1)
	}
}

func writeGroupFile() {
	path := filepath.Join(config.Tmp(), "etc", "group")

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", config.BinName(), err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "%s: warning: %s\n", config.BinName(), err.Error())
		}
	}()

	if err := group.Write(group.File{
		Groups: []group.Group{
			group.Group{
				Name: "root",
				GID:  "0",
			},
			group.Group{
				Name: "nobody",
				GID:  "65534",
			},
			group.Group{
				Name: config.Group(),
				GID:  config.GID(),
			},
		},
	}, file); err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", config.BinName(), err.Error())
		os.Exit(1)
	}
}
