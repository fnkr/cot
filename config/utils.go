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

func listFromEnvs(prefix string) map[string]string {
	list := map[string]string{}

	for _, evar := range os.Environ() {
		if !strings.HasPrefix(evar, prefix) {
			continue
		}

		pair := strings.SplitN(evar, "=", 2)
		if len(pair) < 2 {
			pair[1] = ""
		}

		list[strings.TrimPrefix(pair[0], prefix)] = pair[1]
	}

	return list
}
