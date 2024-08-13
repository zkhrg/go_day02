package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: <program> <command> [args...]")
		os.Exit(1)
	}

	command := os.Args[1]
	cmdArgs := os.Args[2:]

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arg := strings.TrimSpace(scanner.Text())
		if arg == "" {
			continue
		}

		cmd := exec.Command(command, append(cmdArgs, arg)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "error executing command with arg '%s': %v\n", arg, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading from stdin: %v\n", err)
		os.Exit(1)
	}
}
