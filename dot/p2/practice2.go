// example2.go
package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/goadvent-antlr/parser"
)

type calcListener struct {
	*parser.BaseCalcListener
	lexer *parser.CalcLexer
}

// ExitParentheses is called when exiting the Parentheses pair.
func (l *calcListener) ExitParentheses(c *parser.ParenthesesContext) {
	//fmt.Printf("%#v\n", c)
	fmt.Printf("%#v {%q}\n",
		c.Expression(), c.GetText())
}

// ExitMulDiv is called when exiting the MulDiv production.
func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	//fmt.Printf("%#v\n", c)
	fmt.Printf("%s {%q}\n",
		l.lexer.SymbolicNames[c.GetOp().GetTokenType()], c.GetText())
}

// ExitAddSub is called when exiting the AddSub production.
func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	//fmt.Printf("%#v\n", c)
	fmt.Printf("%s {%q}\n",
		l.lexer.SymbolicNames[c.GetOp().GetTokenType()], c.GetText())
}

// ExitNumber is called when exiting the Number production.
func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	//fmt.Printf("%#v\n", c)
	i, _ := strconv.Atoi(c.GetText())
	fmt.Printf("NUMBER {%d}\n", i)
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("12 * (34 - 56)")
	//is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := calcListener{}
	listener.lexer = lexer
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
}

/*

$ go run example2.go
NUMBER {12}
NUMBER {34}
NUMBER {56}
SUB {"34-56"}
&parser.AddSubContext{ExpressionContext:(*parser.ExpressionContext)(...), op:(*antlr.CommonToken)(...)} {"(34-56)"}
MUL {"12*(34-56)"}

*/
