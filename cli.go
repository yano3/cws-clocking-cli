package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		yes    bool
		out    bool
		status bool

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&yes, "yes", false, "Skip y/n prompt")
	flags.BoolVar(&yes, "y", false, "Skip y/n prompt(Short)")
	flags.BoolVar(&out, "out", false, "Clocking out")
	flags.BoolVar(&out, "o", false, "Clocking out(Short)")
	flags.BoolVar(&status, "status", false, "Just show clock in/out time and exit")
	flags.BoolVar(&status, "s", false, "Just show clock in/out time and exit(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = yes

	_ = out

	_ = status

	return ExitCodeOK
}
