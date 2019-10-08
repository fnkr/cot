// +build linux

package config

import (
	"runtime"
	"strconv"
)

func CPUsDefault() string {
	if ToolName() == DOCKER {
		return strconv.FormatFloat(float64(runtime.NumCPU())/1.25, 'f', 6, 64) // 80%
	}
	return ""
}

func MemoryReservationDefault() string {
	if ToolName() == DOCKER {
		return "1g"
	}
	return ""
}
