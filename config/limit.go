package config

import (
	"strings"
)

func IsInLimit(workDir string) bool {
	workDir = strings.TrimSuffix(workDir, "/") + "/"
	for _, dir := range Limit() {
		dir = strings.TrimSuffix(dir, "/") + "/"
		if strings.HasPrefix(workDir, dir) {
			return true
		}
	}
	return false
}
