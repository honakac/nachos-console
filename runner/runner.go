// Package runner contains function for running commands
package runner

import (
	"fmt"
	"os"

	"github.com/honakac/nachos-console/lexer"
	"github.com/honakac/nachos-console/parser"
)

func handleInput(input string) (*parser.Parser, error) {
	l := lexer.New(input)
	if err := l.Handle(); err != nil {
		return nil, err
	}

	p := parser.New(l)
	if err := p.Handle(); err != nil {
		return nil, err
	}

	return p, nil
}

func handleCommandNode(node *parser.CommandNode) error {
	return RunCommand(node.Command, node.Args)
}

func handleRootNode(node *parser.RootNode) error {
	for _, node := range node.Nodes {
		switch n := node.(type) {
		case *parser.CommandNode:
			if err := handleCommandNode(n); err != nil {
				fmt.Fprintf(os.Stderr, "nachos-console: %s\n", err) // Just print error
			}
		}
	}

	return nil
}

func ParseCommand(input string) error {
	p, err := handleInput(input)
	if err != nil {
		return err
	}

	if err := handleRootNode(p.AST); err != nil {
		return err
	}

	return nil
}
