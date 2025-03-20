package main

import (
	"io"
	"log"
	"os"
)

// Log are global logger variable.
var (
	Log *Logger
)

const logFlags = 0

// const logFlags = log.LstdFlags

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
}

func setLogger(level string) {
	switch level {
	case "debug":
		setLoggerDebug()
	case "silent":
		setLoggerSilent()
	default:
		setLoggerDefault()
	}
}

func setLoggerDefault() {
	Log = &Logger{
		Info:  log.New(os.Stdout, "INFO: ", logFlags),
		Error: log.New(os.Stderr, "ERROR: ", logFlags),
		Debug: log.New(io.Discard, "DEBUG: ", logFlags),
	}
}

func setLoggerSilent() {
	Log = &Logger{
		Info:  log.New(io.Discard, "INFO: ", logFlags),
		Error: log.New(os.Stderr, "ERROR: ", logFlags),
		Debug: log.New(io.Discard, "DEBUG: ", logFlags),
	}
}

func setLoggerDebug() {
	Log = &Logger{
		Info:  log.New(os.Stdout, "INFO: ", logFlags),
		Error: log.New(os.Stderr, "ERROR: ", logFlags),
		Debug: log.New(os.Stdout, "DEBUG: ", logFlags),
	}
}
