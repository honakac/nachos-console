package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/honakac/nachos-console/runner"
)

func main() {
	if len(os.Args) >= 2 {
		buffer, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "nachos-console: %v\n", err)
			os.Exit(1)
		}

		if err := runner.ParseCommand(string(buffer)); err != nil {
			fmt.Fprintf(os.Stderr, "nachos-console: %v\n", err)
			os.Exit(1)
		}

		return
	}
	completer := readline.NewPrefixCompleter(
		readline.PcItem("exit"),
	)

	rl, err := readline.NewEx(&readline.Config{
		AutoComplete: completer,
		Prompt:       "> ",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			continue
		}
		line = strings.TrimSpace(line)

		if line == "exit" {
			break
		}

		if err := runner.ParseCommand(line); err != nil {
			fmt.Fprintf(os.Stderr, "nachos-console: %v\n", err)
		}
	}
}
