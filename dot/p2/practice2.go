// example2.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/dot"
	"github.com/go-sqlparser/goadvent-antlr/parser"
)

type dotListener struct {
	*dot.BaseDOTListener
	lexer *dot.DOTLexer
}

// ExitMulDiv is called when exiting the MulDiv production.
func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	//fmt.Printf("%#v\n", c)
	fmt.Printf("%s {%q}\n",
		l.lexer.SymbolicNames[c.GetOp().GetTokenType()], c.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (s *dotListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

// ExitGraph is called when production graph is exited.
func (s *dotListener) ExitGraph(ctx *dot.GraphContext) {
}

// ExitStmt_list is called when production stmt_list is exited.
func (s *dotListener) ExitStmt_list(ctx *dot.Stmt_listContext) {
}

// ExitStmt is called when production stmt is exited.
func (s *dotListener) ExitStmt(ctx *dot.StmtContext) {
}

// ExitAttr_stmt is called when production attr_stmt is exited.
func (s *dotListener) ExitAttr_stmt(ctx *dot.Attr_stmtContext) {
}

// ExitAttr_list is called when production attr_list is exited.
func (s *dotListener) ExitAttr_list(ctx *dot.Attr_listContext) {
}

// ExitA_list is called when production a_list is exited.
func (s *dotListener) ExitA_list(ctx *dot.A_listContext) {
}

// ExitEdge_stmt is called when production edge_stmt is exited.
func (s *dotListener) ExitEdge_stmt(ctx *dot.Edge_stmtContext) {
}

// ExitEdgeRHS is called when production edgeRHS is exited.
func (s *dotListener) ExitEdgeRHS(ctx *dot.EdgeRHSContext) {
}

// ExitEdgeop is called when production edgeop is exited.
func (s *dotListener) ExitEdgeop(ctx *dot.EdgeopContext) {
}

// ExitNode_stmt is called when production node_stmt is exited.
func (s *dotListener) ExitNode_stmt(ctx *dot.Node_stmtContext) {
}

// ExitNode_id is called when production node_id is exited.
func (s *dotListener) ExitNode_id(ctx *dot.Node_idContext) {
}

// ExitPort is called when production port is exited.
func (s *dotListener) ExitPort(ctx *dot.PortContext) {
}

// ExitSubgraph is called when production subgraph is exited.
func (s *dotListener) ExitSubgraph(ctx *dot.SubgraphContext) {
}

// ExitId is called when production id is exited.
func (s *dotListener) ExitId(ctx *dot.IdContext) {
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
