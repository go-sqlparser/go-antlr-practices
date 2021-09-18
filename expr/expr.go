// https://stackoverflow.com/questions/64842665/

package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/go-antlr-practices/expr/parser"
)

func calc(inputfile string) {
	input, _ := antlr.NewFileStream(inputfile) // Setup the input
	lexer := parser.NewExprLexer(input)        // Create the Lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(stream) // Create the Parser
	p.Stat()
}

func main() {
	calc(os.Args[1])
}
