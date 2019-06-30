package config

import (
	"fmt"
	"io"
)

func PrintConfig(w io.Writer) {
	fmt.Fprintf(w, "BinName() = %v\n", BinName())
	fmt.Fprintf(w, "CmdAndArgs() = %v\n", CmdAndArgs())
	fmt.Fprintf(w, "Cmd() = %v\n", Cmd())
	fmt.Fprintf(w, "ToolArgs() = %v\n", Args())
	fmt.Fprintf(w, "WorkDir() = %v\n", WorkDir())
	fmt.Fprintf(w, "Docker() = %v\n", Docker())
	fmt.Fprintf(w, "Podman() = %v\n", Podman())
	fmt.Fprintf(w, "SELinuxEnabled() = %v\n", SELinuxEnabled())
	fmt.Fprintf(w, "UID() = %v\n", UID())
	fmt.Fprintf(w, "GID() = %v\n", GID())
	fmt.Fprintf(w, "User() = %v\n", User())
	fmt.Fprintf(w, "Group() = %v\n", Group())
	fmt.Fprintf(w, "DockerSocketIsWritable() = %v\n", DockerSocketIsWritable())
	fmt.Fprintf(w, "HomeDir() = %v\n", HomeDir())
	fmt.Fprintf(w, "Tool() = %v\n", Tool())
	fmt.Fprintf(w, "ToolName() = %v\n", ToolName())
	fmt.Fprintf(w, "ToolArgs() = %v\n", ToolArgs())
	fmt.Fprintf(w, "Image() = %v\n", Image())
	fmt.Fprintf(w, "Tmp() = %v\n", Tmp())
	fmt.Fprintf(w, "Network() = %v\n", Network())
	fmt.Fprintf(w, "TTY() = %v\n", TTY())
	fmt.Fprintf(w, "Interactive() = %v\n", Interactive())
	fmt.Fprintf(w, "Limit() = %v\n", Limit())
	fmt.Fprintf(w, "LimitString() = %v\n", LimitString())
	fmt.Fprintf(w, "SSHAuthSock() = %v\n", SSHAuthSock())
	fmt.Fprintf(w, "Shell() = %v\n", Shell())
	fmt.Fprintf(w, "CPUs() = %v\n", CPUs())
	fmt.Fprintf(w, "Memory() = %v\n", Memory())
	fmt.Fprintf(w, "MemoryReservation() = %v\n", MemoryReservation())
	fmt.Fprintf(w, "Debug() = %v\n", Debug())
	fmt.Fprintf(w, "DryRun() = %v\n", DryRun())
}
