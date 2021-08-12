// example2.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/dot"
)

type dotListener struct {
	*dot.BaseDOTListener
	//lexer  *dot.DOTLexer
	parser *dot.DOTParser
}

// ExitEveryRule is called when any rule is exited.
func (s *dotListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Printf("%#v\n", ctx)
	rc := ctx.GetRuleContext()
	ri := rc.GetRuleIndex()
	//fmt.Printf("%#v\n", rc)
	fmt.Printf("%s (%s)\n", s.parser.RuleNames[ri], ctx.GetText())
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
	is, _ := antlr.NewFileStream("../sales.xo.dot")

	// Create the Lexer
	lexer := dot.NewDOTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := dot.NewDOTParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := dotListener{}
	//listener.lexer = lexer
	listener.parser = p
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Graph())
}

/*

$ p2
id (dbo)
id (shape)
id (none)
id (margin)
id (0)
a_list (shape=none,margin=0)
attr_list ([shape=none,margin=0])
attr_stmt (node[shape=none,margin=0])
stmt (node[shape=none,margin=0])
id ("dbo.Customer")
node_id ("dbo.Customer")
id (label)
id (<
                <table border="0" cellborder="1" cellspacing="0" cellpadding="4">
                <tr><td bgcolor="lightblue">"dbo.Customer"</td></tr>
                <tr><td align="left" PORT="CustomerID">CustomerID: int</td></tr>
                <tr><td align="left" PORT="Name">Name: varchar</td></tr>
                <tr><td align="left" PORT="Address1">Address1: varchar</td></tr>
                <tr><td align="left" PORT="Address2">Address2: varchar</td></tr>
                <tr><td align="left" PORT="Address3">Address3: varchar</td></tr>
                </table>>)
a_list (label=<
                <table border="0" cellborder="1" cellspacing="0" cellpadding="4">
                <tr><td bgcolor="lightblue">"dbo.Customer"</td></tr>
                <tr><td align="left" PORT="CustomerID">CustomerID: int</td></tr>
                <tr><td align="left" PORT="Name">Name: varchar</td></tr>
                <tr><td align="left" PORT="Address1">Address1: varchar</td></tr>
                <tr><td align="left" PORT="Address2">Address2: varchar</td></tr>
                <tr><td align="left" PORT="Address3">Address3: varchar</td></tr>
                </table>>)
. . .

*/
