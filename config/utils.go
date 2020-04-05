package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func boolFromStr(val string, def bool) bool {
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

func boolFromEnv(key string, def bool) bool {
	return boolFromStr(os.Getenv(key), def)
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

func stringInSlice(search string, slice []string) bool {
	for _, value := range slice {
		if search == value {
			return true
		}
	}
	return false
}

func removeStringFromSlice(search string, slice []string) []string {
	for i, value := range slice {
		if value == search {
			return removeStringFromSlice(search, append(slice[:i], slice[i+1:]...))
		}
	}
	return slice
}
