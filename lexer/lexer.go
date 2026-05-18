// Package lexer implements a lexer
package lexer

import "fmt"

const EOF = 0

type Lexer struct {
	tokens []Token

	input        []rune
	inputLength  int
	position     int
	readPosition int
	char         rune

	buffer []rune
	line   uint
}

func New(input string) *Lexer {
	return &Lexer{
		input:       []rune(input),
		inputLength: len(input),
	}
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= l.inputLength {
		l.char = EOF
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// Add token by exist information and tokenType.
// After allocate new buffer after adding token
func (l *Lexer) addToken(tokenType TokenType) {
	if len(l.buffer) > 0 {
		l.tokens = append(l.tokens, Token{
			Type:    tokenType,
			Literal: l.buffer,
			Line:    l.line,
			Column:  uint(l.position) - uint(len(l.buffer)) + 1,
		})

		l.buffer = make([]rune, 0)
	}
}

func (l *Lexer) appendChar() {
	l.buffer = append(l.buffer, l.char)
}

func (l *Lexer) handleString() {
	// Skip " symbol
	l.ReadChar()

exitFor:
	for {
		switch l.char {
		case '"':
			l.addToken(Word)
			break exitFor
		default:
			l.appendChar()
		}

		l.ReadChar()
	}

	// Again skip " symbol
	l.ReadChar()
}

func (l *Lexer) Handle() error {
	l.ReadChar()

exitFor:
	for {
		switch l.char {
		// Parse simple words
		case ' ', EOF:
			l.addToken(Word)

			if l.char == EOF {
				break exitFor
			}
		case '"':
			l.handleString()
		default:
			l.appendChar()
		}

		l.ReadChar()
	}

	for i, t := range l.tokens {
		fmt.Printf("%d: %s\n", i, t.String())
	}

	return nil
}
