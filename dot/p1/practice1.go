// example1.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/dot"
)

func main() {
	// Setup the input
	// https://pkg.go.dev/github.com/antlr/antlr4/runtime/Go/antlr#NewFileStream
	is, _ := antlr.NewFileStream("../sales.xo.dot")

	// Create the Lexer
	// https://pkg.go.dev/github.com/go-sqlparser/antlr4-grammars
	lexer := dot.NewDOTLexer(is)

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
