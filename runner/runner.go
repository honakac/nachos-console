// Package runner contains function for running commands
package runner

import (
	"github.com/honakac/nachos-console/lexer"
)

func ParseCommand(input string) error {
	l := lexer.New(input)
	if err := l.Handle(); err != nil {
		return err
	}

	return nil
}
