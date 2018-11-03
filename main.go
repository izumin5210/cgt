package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var exitCode int

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
	}

	os.Exit(exitCode)
}

func run() error {
	if terminal.IsTerminal(0) {
		return errors.New("require stdin from pipe")
	}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		parse(s.Text())
	}

	return nil
}

var (
	success = aurora.Green
	fail    = aurora.Red
	skipped = aurora.Brown
	info    = aurora.Gray

	c func(interface{}) aurora.Value
)

// ref: https://github.com/rakyll/gotest/blob/86f0749cd8ccdc08f2edb69f170ad6f06393455d/main.go#L77-L108
func parse(line string) {
	trimmed := strings.TrimSpace(line)

	switch {
	// start
	case strings.HasPrefix(trimmed, "=== RUN"):
		c = aurora.Bold
	case strings.HasPrefix(trimmed, "?"):
		c = nil

	// info
	case strings.HasPrefix(trimmed, "=== CONT"):
		fallthrough
	case strings.HasPrefix(trimmed, "=== PAUSE"):
		c = info

	// success
	case strings.HasPrefix(trimmed, "--- PASS"):
		fallthrough
	case strings.HasPrefix(trimmed, "ok"):
		fallthrough
	case strings.HasPrefix(trimmed, "PASS"):
		c = success

	// failure
	case strings.HasPrefix(trimmed, "--- FAIL"):
		fallthrough
	case strings.HasPrefix(trimmed, "FAIL"):
		c = fail

	// skipped
	case strings.HasPrefix(trimmed, "--- SKIP"):
		c = skipped
	}

	if c == nil {
		fmt.Println(line)
		return
	}
	fmt.Println(c(line))
}
