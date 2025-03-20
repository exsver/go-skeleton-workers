package main

import "flag"

// Flags defines configuration passed by flags.
type Flags struct {
	ConfigFilePath *string
	LogLevel       *string
}

func parseFlags() *Flags {
	var f Flags

	f.ConfigFilePath = flag.String("config", "./config.json", "path to config file")
	f.LogLevel = flag.String("log-level", "default", "log-level: default|silent|debug")

	flag.Parse()

	return &f
}
