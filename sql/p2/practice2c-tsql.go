// example2.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/tsql"
)

type tsqlListener struct {
	*tsql.BaseTSqlParserListener
	parser *tsql.TSqlParser
}

// ExitTable_name is called when production table_name is exited.
func (s *tsqlListener) ExitTable_name(ctx *tsql.Table_nameContext) {
	fmt.Printf("%s\n", ctx.GetText())
}

// ExitSelect_statement is called when production select_statement is exited.
func (s *tsqlListener) ExitSelect_statement(ctx *tsql.Select_statementContext) {
	fmt.Println()
}

func main() {
	// Setup the case insensitive input streams
	// https://github.com/antlr/antlr4/blob/master/doc/case-insensitive-lexing.md
	is, _ := antlr.NewFileStream("../test-sample.sql")
	upper := NewCaseChangingStream(is, true)
	// Create the Lexer
	lexer := tsql.NewTSqlLexer(upper)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := tsql.NewTSqlParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := tsqlListener{}
	listener.parser = p
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Tsql_file())
}

/*

$ go run practice2c-tsql.go case_changing_stream.go
PERSON.PERSON
HUMANRESOURCES.EMPLOYEE
SALES.SALESPERSON

Person.Person
HumanResources.Employee
Sales.SalesPerson

Sales.SalesOrderDetail

Sales.SalesOrderHeader

Production.Product
Production.Product


Person.Person
HumanResources.Employee
Sales.SalesPerson



Person.Person
HumanResources.Employee
Sales.SalesPerson


production.products
production.products
production.brands



[AdventureWorks2014].[Production].[ProductCategory]
[AdventureWorks2014].[Production].[ProductSubcategory]
[AdventureWorks2014].[Production].[ProductCategory]


*/
