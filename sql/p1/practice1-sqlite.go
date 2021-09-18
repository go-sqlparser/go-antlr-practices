// example1.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/sqlite"
)

func main() {
	// Setup the input
	// https://pkg.go.dev/github.com/antlr/antlr4/runtime/Go/antlr#NewFileStream
	is, _ := antlr.NewFileStream("../test-sample.sql")

	// Create the Lexer
	// https://pkg.go.dev/github.com/go-sqlparser/antlr4-grammars
	lexer := sqlite.NewSQLiteLexer(is)

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

$ go run practice1-sqlite.go
SELECT_ ("SELECT")
SPACES (" ")
IDENTIFIER ("LASTNAME")
COMMA (",")
SPACES (" ")
IDENTIFIER ("FIRSTNAME")
SPACES ("\n")
FROM_ ("FROM")
SPACES (" ")
IDENTIFIER ("PERSON")
DOT (".")
IDENTIFIER ("PERSON")
SPACES (" ")
IDENTIFIER ("C")
SPACES ("\n")
INNER_ ("INNER")
SPACES (" ")
JOIN_ ("JOIN")
SPACES (" ")
IDENTIFIER ("HUMANRESOURCES")
DOT (".")
IDENTIFIER ("EMPLOYEE")
SPACES (" ")
IDENTIFIER ("E")
SPACES ("\n")
ON_ ("ON")
SPACES (" ")
IDENTIFIER ("C")
DOT (".")
IDENTIFIER ("BUSINESSENTITYID")
SPACES (" ")
ASSIGN ("=")
SPACES (" ")
IDENTIFIER ("E")
DOT (".")
IDENTIFIER ("BUSINESSENTITYID")
SPACES ("\n")
JOIN_ ("JOIN")
SPACES (" ")
IDENTIFIER ("SALES")
DOT (".")
IDENTIFIER ("SALESPERSON")
SPACES (" ")
IDENTIFIER ("S")
SPACES (" ")
SPACES ("\n")
ON_ ("ON")
SPACES (" ")
IDENTIFIER ("E")
DOT (".")
IDENTIFIER ("BUSINESSENTITYID")
SPACES (" ")
ASSIGN ("=")
SPACES (" ")
IDENTIFIER ("S")
DOT (".")
IDENTIFIER ("BUSINESSENTITYID")
SCOL (";")

*/
