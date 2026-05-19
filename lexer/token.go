package lexer

import "fmt"

const (
	None = iota
	Newline
	Word
	Pipe
)

type TokenType uint8

func (t TokenType) String() string {
	switch t {
	case Word:
		return "Word"
	case Pipe:
		return "Pipe"
	case Newline:
		return "Newline"
	default:
		return "Unknown"
	}
}

type Token struct {
	Type    TokenType // Token type (example None, Word...)
	Literal []rune    // Token
	Line    uint      // Line number
	Column  uint      // Position number
}

func (t *Token) String() string {
	return fmt.Sprintf("Type=%v, Line=%d, Column=%d, Literal=\"%s\"", t.Type, t.Line, t.Column, string(t.Literal))
}
