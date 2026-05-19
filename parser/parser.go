// Package parser implements a parser, which builds an AST tree
package parser

import (
	"fmt"

	"github.com/honakac/nachos-console/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
	AST   *RootNode

	token        *lexer.Token
	position     int
	readPosition int
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
		AST: &RootNode{
			Nodes: make([]ASTNode, 0),
		},
	}
}

func (p *Parser) NextToken() {
	if p.readPosition >= len(p.lexer.Tokens) {
		p.token = nil
	} else {
		p.token = &p.lexer.Tokens[p.readPosition]
	}

	p.position = p.readPosition
	p.readPosition++
}

func (p *Parser) Handle() error {
	p.NextToken()

	for p.token != nil {
		if p.token.Type == lexer.Word {
			commandNode := &CommandNode{
				Command: string(p.token.Literal),
				Args:    make([]string, 0),
			}

			// Read arguments
			p.NextToken()
		argsFor:
			for p.token != nil {
				switch p.token.Type {
				case lexer.Word:
					commandNode.Args = append(commandNode.Args, string(p.token.Literal))
				case lexer.Newline: // If is end
					break argsFor
				}

				p.NextToken()
			}

			p.AST.Nodes = append(p.AST.Nodes, commandNode)
		}

		p.NextToken()
	}

	fmt.Printf("Parser result:\n%s\n", p.AST)

	return nil
}
