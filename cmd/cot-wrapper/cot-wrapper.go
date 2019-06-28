package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	container := os.Getenv("COT_CONTAINER")

	var cmd []string

	if container != "podman" && container != "docker" {
		if cot, err := exec.LookPath("cot"); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		} else {
			cmd = append(cmd, cot)
		}
	}

	cmd = append(cmd, "foo")

	cmd = append(cmd, os.Args[1:]...)

	fmt.Println(cmd)
	os.Exit(2)

	if err := syscall.Exec(cmd[0], cmd, os.Environ()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
