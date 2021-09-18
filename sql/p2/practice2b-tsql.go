// example2.go
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/go-sqlparser/antlr4-grammars/tsql"
)

type tsqlListener struct {
	*tsql.BaseTSqlParserListener
}

const sql = `
SELECT LASTNAME, FIRSTNAME
FROM PERSON.PERSON C
INNER JOIN HUMANRESOURCES.EMPLOYEE E
ON C.BUSINESSENTITYID = E.BUSINESSENTITYID
JOIN SALES.SALESPERSON S 
ON E.BUSINESSENTITYID = S.BUSINESSENTITYID;

SELECT LASTNAME, FIRSTNAME
FROM PERSON.PERSON C
INNER JOIN HUMANRESOURCES.EMPLOYEE E
ON C.BUSINESSENTITYID = E.BUSINESSENTITYID
JOIN SALES.SALESPERSON S 
ON E.BUSINESSENTITYID = S.BUSINESSENTITYID

SELECT LastName, FirstName
FROM Person.Person c
INNER JOIN HumanResources.Employee e
ON c.BusinessEntityID = e.BusinessEntityID
JOIN Sales.SalesPerson s 
ON e.BusinessEntityID = s.BusinessEntityID

SELECT [Name]
FROM Production.Product
WHERE ListPrice =
    (SELECT ListPrice
     FROM Production.Product
     WHERE [Name] = 'Chainring Bolts' )
`

// ExitTable_name is called when production table_name is exited.
func (s *tsqlListener) ExitTable_name(ctx *tsql.Table_nameContext) {
	fmt.Printf("%s\n", ctx.GetText())
}

// ExitSelect_statement is called when production select_statement is exited.
func (s *tsqlListener) ExitSelect_statement(ctx *tsql.Select_statementContext) {
	fmt.Println()
}

func main() {
	// Setup the input
	is := antlr.NewInputStream(sql)

	// Create the Lexer
	lexer := tsql.NewTSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := tsql.NewTSqlParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := tsqlListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Dml_clause())
}

/*

$ go run practice2.go

*/
