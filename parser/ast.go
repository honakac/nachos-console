package parser

import (
	"strings"
)

type ASTNode interface {
	String() string
}

type RootNode struct {
	Nodes []ASTNode
}

type CommandNode struct {
	Command string
	Args    []string
}

// Implementing structure conversion to a string:

func (n *RootNode) String() string {
	var builder strings.Builder

	for i, node := range n.Nodes {
		if node == nil {
			continue
		}
		if i > 0 {
			builder.WriteString("\n")
		}

		builder.WriteString((node).String())
	}

	return builder.String()
}

func (n *CommandNode) String() string {
	var builder strings.Builder

	builder.WriteString("Command: ")
	builder.WriteString(n.Command)

	builder.WriteString(" Args: [")
	for i, arg := range n.Args {
		if i > 0 {
			builder.WriteString(", ")
		}

		builder.WriteString(arg)
	}
	builder.WriteString("]")

	return builder.String()
}
