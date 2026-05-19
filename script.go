package main

import (
	"os"

	"github.com/honakac/nachos-console/runner"
)

func runScript(filename string) error {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := runner.ParseCommand(string(buffer)); err != nil {
		return err
	}

	return nil
}
