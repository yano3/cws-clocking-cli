package main

import (
	"runtime/debug"
)

const Name string = "cws-clocking-cli"

var version = ""

func getVersion() string {
	if version != "" {
		return version
	}

	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "(unknown)"
	}
	return i.Main.Version
}
