package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fnkr/cot/config"
)

func createTmpDirs() {
	// TODO: Add debug output

	if err := os.MkdirAll(config.Tmp(), os.FileMode(0750)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Join(config.Tmp(), "etc"), os.FileMode(0777)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Join(config.Tmp(), "home"), os.FileMode(0777)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Join(config.Tmp(), "tmp"), os.FileMode(0777)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
