// Package config implements working with configs and env.
package config

import "os"

var DebugLexer bool
var DebugParser bool

func init() {
	if os.Getenv("NC_DEBUG_LEXER") == "1" {
		DebugLexer = true
	}
	if os.Getenv("NC_DEBUG_PARSER") == "1" {
		DebugParser = true
	}
}
