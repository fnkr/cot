package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func boolFromEnv(key string, def bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	res, err := strconv.ParseBool(val)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: error: %s\n", BinName(), err.Error())
		os.Exit(1)
	}
	return res
}

func listFromEnv(key, sep string) (res []string) {
	for _, val := range strings.Split(os.Getenv(key), sep) {
		if val == "" {
			continue
		}
		res = append(res, val)
	}
	return
}

