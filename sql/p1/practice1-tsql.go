// example1.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/tsql"
)

func main() {
	// Setup the case insensitive input streams
	// https://github.com/antlr/antlr4/blob/master/doc/case-insensitive-lexing.md
	is, _ := antlr.NewFileStream("../test-sample.sql")
	upper := NewCaseChangingStream(is, true)
	// Create the Lexer
	lexer := tsql.NewTSqlLexer(upper)

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

$ go run practice1.go
SELECT ("SELECT")
ID ("LASTNAME")
COMMA (",")
ID ("FIRSTNAME")
FROM ("FROM")
ID ("PERSON")
DOT (".")
ID ("PERSON")
ID ("C")
INNER ("INNER")
JOIN ("JOIN")
ID ("HUMANRESOURCES")
DOT (".")
ID ("EMPLOYEE")
ID ("E")
ON ("ON")
ID ("C")
DOT (".")
ID ("BUSINESSENTITYID")
EQUAL ("=")
ID ("E")
DOT (".")
ID ("BUSINESSENTITYID")
JOIN ("JOIN")
ID ("SALES")
DOT (".")
ID ("SALESPERSON")
ID ("S")
ON ("ON")
ID ("E")
DOT (".")
ID ("BUSINESSENTITYID")
EQUAL ("=")
ID ("S")
DOT (".")
ID ("BUSINESSENTITYID")
SEMI (";")

*/
