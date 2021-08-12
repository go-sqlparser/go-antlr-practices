// example1.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/goadvent-antlr/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("22 * (34 + 56)")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

/*

$ go run example1.go
OPP ("(")
NUMBER ("1")
ADD ("+")
NUMBER ("2")
CLP (")")
MUL ("*")
NUMBER ("3")

*/
