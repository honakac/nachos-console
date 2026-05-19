// Package runner contains function for running commands
package runner

import (
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

func handleCommandNode(node *parser.CommandNode) {
	RunCommand(node.Command, node.Args)
}

func handleRootNode(node *parser.RootNode) {
	for _, node := range node.Nodes {
		switch n := node.(type) {
		case *parser.CommandNode:
			handleCommandNode(n)
		}
	}
}

func ParseCommand(input string) error {
	p, err := handleInput(input)
	if err != nil {
		return err
	}

	handleRootNode(p.AST)

	return nil
}
