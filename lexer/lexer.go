// Package lexer implements a lexer
package lexer

import (
	"fmt"
	"strings"

	"github.com/honakac/nachos-console/config"
)

const EOF = 0

type Lexer struct {
	Tokens []Token

	input        []byte
	inputLength  int
	position     int
	readPosition int
	char         byte

	buffer []byte
	line   uint
}

func New(input string) *Lexer {
	return &Lexer{
		input:       []byte(input),
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

func (l *Lexer) SkipChars(chars string) {
	for strings.ContainsRune(chars, rune(l.char)) {
		l.ReadChar()
	}
}

func (l *Lexer) SkipWhitespace() {
	l.SkipChars(" \t")
}

// Add token by exist information and tokenType.
// After allocate new buffer after adding token
func (l *Lexer) addToken(tokenType TokenType) {
	if len(l.buffer) > 0 {
		l.Tokens = append(l.Tokens, Token{
			Type:    tokenType,
			Literal: l.buffer,
			Line:    l.line,
			Column:  uint(l.position) - uint(len(l.buffer)) + 1,
		})

		l.buffer = make([]byte, 0)
	}
}

func (l *Lexer) appendTokenLiteral(tokenType TokenType, literal []byte) {
	l.Tokens = append(l.Tokens, Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  uint(l.position) - uint(len(l.buffer)) + 1,
	})
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
		case ' ', EOF, '\n', '\\', ';':
			if l.char != '\\' {
				l.addToken(Word)
			}

			switch l.char {
			case EOF:
				break exitFor
			case '\\':
				l.ReadChar()
				if l.char != '\n' { // \ \n is just removed
					l.appendChar()
				}
			case '\n', ';':
				l.appendTokenLiteral(Newline, nil)
			}
		case '#':
			if l.position == 0 || l.input[l.position-1] == ' ' || l.input[l.position-1] == '\n' {
				for l.char != '\n' && l.char != EOF {
					l.ReadChar()
				}
			} else {
				l.appendChar()
			}
		case '"':
			l.handleString()
			continue
		default:
			l.appendChar()
		}

		l.ReadChar()
	}
	l.appendTokenLiteral(Newline, nil)

	if config.DebugLexer {
		fmt.Println("Lexer result:")
		for i, t := range l.Tokens {
			fmt.Printf("%d: %s\n", i, t.String())
		}
	}

	return nil
}
