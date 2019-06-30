package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	isInitTool     bool
	tool           []string
	isInitToolName bool
	toolName       string
	isInitToolArgs bool
	toolArgs       []string
)

func Tool() []string {
	if !isInitTool {
		t := os.Getenv(EnvPrefix() + "_TOOL")
		if t == "" {
			if Podman() != "" {
				tool = []string{Podman()}
			} else if Docker() != "" {
				if DockerSocketIsWritable() {
					tool = []string{Docker()}
				} else {
					if Sudo() != "" {
						tool = []string{Sudo(), Docker()}
					} else {
						fmt.Fprintf(os.Stderr, "%s: error: user not in docker group and sudo not found in $PATH", BinName())
						os.Exit(1)
					}
				}
			} else {
				fmt.Fprintf(os.Stderr, "%s: error: neither podman nor docker found in $PATH", BinName())
				os.Exit(1)
			}
		} else {
			tool = strings.Split(t, " ")
			if resolvedTool, err := exec.LookPath(tool[0]); resolvedTool == "" {
				fmt.Fprintf(os.Stderr, "%s: could not find %s in $PATH", BinName(), tool[0])
				os.Exit(1)
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "%s: exec.LookPath(tool[0]): %s", BinName(), err.Error())
				os.Exit(1)
			} else {
				tool[0] = resolvedTool
			}
		}
		isInitTool = true
	}

	return tool
}

func ToolName() string {
	if !isInitToolName {
		for _, toolPart := range Tool() {
			if strings.HasSuffix(toolPart, "/podman") || toolPart == "podman" {
				toolName = "podman"
			} else if strings.HasSuffix(toolPart, "/docker") || toolPart == "docker" {
				toolName = "docker"
			}
			if toolName != "" {
				break
			}
		}
		if toolName == "" {
			fmt.Fprintf(os.Stderr, "%s: could not detect tool name: %s", BinName(), Tool())
			os.Exit(1)
		}
		isInitToolName = true
	}

	return toolName

}

func ToolArgs() []string {
	if !isInitToolArgs {
		toolArgs = listFromEnv(EnvPrefix()+"_ARGS", " ")
		isInitToolArgs = true
	}

	return toolArgs
}
