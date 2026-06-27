package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: issues [subcommand] [options]")
		os.Exit(1)
	}

	switch args[1] {
	case "create":
		createIssue(args[2:])
	case "close":
	case "list":
	case "read":
	case "update":
	}
}

func createIssue(args []string) {
	flag.NewFlagSet("create", flag.ExitOnError)
}
