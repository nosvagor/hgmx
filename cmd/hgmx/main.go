package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"

	"github.com/nosvagor/hgmx/cmd/hgmx/sloghandler"
)

var Version = "dev"

func main() {
	code := run(os.Stdin, os.Stdout, os.Stderr, os.Args)
	if code != 0 {
		os.Exit(code)
	}
}

const usageText = `usage: hgmx <command> [<args>...]

hgmx - A component management tool for Go Templ and HTMX projects.

commands:
  info      Displays information about the hgmx environmen
  version   Prints the version
`

func run(stdin io.Reader, stdout, stderr io.Writer, args []string) (code int) {
	if len(args) < 2 {
		fmt.Fprint(stderr, usageText)
		return 64
	}

	switch args[1] {
	case "info":
		return infoCmd(stdout, stderr, args[2:])
	// TODO: Add 'add' command
	// TODO: Add 'init' command
	case "version", "--version", "-v":
		fmt.Fprintln(stdout, Version)
		return 0
	case "help", "-help", "--help", "-h":
		fmt.Fprint(stdout, usageText)
		return 0
	}

	fmt.Fprintf(stderr, "Unknown command: %q\n", args[1])
	fmt.Fprint(stderr, usageText)
	return 64
}

func newLogger(logLevel string, verbose bool, stderr io.Writer) *slog.Logger {
	if verbose {
		logLevel = "debug"
	}

	level := slog.LevelInfo
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		if logLevel != "info" {
			fmt.Fprintf(stderr, "Invalid log level %q, defaulting to info\n", logLevel)
		}
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	h := sloghandler.NewHandler(stderr, opts)
	return slog.New(h)
}

// --- info command ---

const infoUsageText = `usage: hgmx info [<args>]

Displays information about the hgmx environment.

Args:
  -v           Set log verbosity level to "debug". (default "info")
  -log-level   Set log verbosity level. (default "info", options: "debug", "success", "info", "warn", "error")
  -help        Print help and exit.
`

func infoCmd(stdout, stderr io.Writer, args []string) (code int) {
	cmd := flag.NewFlagSet("info", flag.ExitOnError)
	verboseFlag := cmd.Bool("v", false, "")
	logLevelFlag := cmd.String("log-level", "info", "")
	helpFlag := cmd.Bool("help", false, "")

	cmd.Usage = func() {
		fmt.Fprint(stderr, infoUsageText)
	}

	err := cmd.Parse(args)
	if err != nil {
		return 64
	}

	if *helpFlag {
		fmt.Fprint(stdout, infoUsageText)
		return 0
	}

	lg := newLogger(*logLevelFlag, *verboseFlag, stderr)

	lg.Info("Environment:",
		slog.Group("version",
			slog.String("hgmx", Version),
			slog.String("go", runtime.Version()),
		),
	)

	return 0
}
