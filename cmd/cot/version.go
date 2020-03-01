package main

var (
	ref       string
	sha       string
	buildDate string
)

type VersionInfo struct {
	Ref       string
	SHA       string
	BuildDate string
}

func Version() VersionInfo {
	return VersionInfo{
		Ref:       ref,
		SHA:       sha,
		BuildDate: buildDate,
	}
}
