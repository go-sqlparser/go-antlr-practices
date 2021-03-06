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

// ExitEveryRule is called when any rule is exited.
func (s *tsqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Printf("%#v\n", ctx)
	rc := ctx.GetRuleContext()
	ri := rc.GetRuleIndex()
	//fmt.Printf("%#v\n", rc)
	fmt.Printf("%s (%s)\n", s.parser.RuleNames[ri], ctx.GetText())
}

func main() {
	// Setup the input
	is, _ := antlr.NewFileStream("../test-sample.sql")

	// Create the Lexer
	lexer := tsql.NewTSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := tsql.NewTSqlParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := tsqlListener{}
	listener.parser = p
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Tsql_file())
}

/*

$ go run practice2.go
id_ (LASTNAME)
full_column_name (LASTNAME)
column_elem (LASTNAME)
select_list_elem (LASTNAME)
id_ (FIRSTNAME)
full_column_name (FIRSTNAME)
column_elem (FIRSTNAME)
select_list_elem (FIRSTNAME)
select_list (LASTNAME,FIRSTNAME)
id_ (PERSON)
id_ (PERSON)
table_name (PERSON.PERSON)
table_name_with_hint (PERSON.PERSON)
id_ (C)
table_alias (C)
as_table_alias (C)
table_source_item (PERSON.PERSONC)
id_ (HUMANRESOURCES)
id_ (EMPLOYEE)
table_name (HUMANRESOURCES.EMPLOYEE)
table_name_with_hint (HUMANRESOURCES.EMPLOYEE)
id_ (E)
table_alias (E)
as_table_alias (E)
table_source_item (HUMANRESOURCES.EMPLOYEEE)
table_source_item_joined (HUMANRESOURCES.EMPLOYEEE)
table_source (HUMANRESOURCES.EMPLOYEEE)
id_ (C)
id_ (BUSINESSENTITYID)
full_column_name (C.BUSINESSENTITYID)
expression (C.BUSINESSENTITYID)
comparison_operator (=)
id_ (E)
id_ (BUSINESSENTITYID)
full_column_name (E.BUSINESSENTITYID)
expression (E.BUSINESSENTITYID)
predicate (C.BUSINESSENTITYID=E.BUSINESSENTITYID)
search_condition (C.BUSINESSENTITYID=E.BUSINESSENTITYID)
join_on (INNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYID)
join_part (INNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYID)
id_ (SALES)
id_ (SALESPERSON)
table_name (SALES.SALESPERSON)
table_name_with_hint (SALES.SALESPERSON)
id_ (S)
table_alias (S)
as_table_alias (S)
table_source_item (SALES.SALESPERSONS)
table_source_item_joined (SALES.SALESPERSONS)
table_source (SALES.SALESPERSONS)
id_ (E)
id_ (BUSINESSENTITYID)
full_column_name (E.BUSINESSENTITYID)
expression (E.BUSINESSENTITYID)
comparison_operator (=)
id_ (S)
id_ (BUSINESSENTITYID)
full_column_name (S.BUSINESSENTITYID)
expression (S.BUSINESSENTITYID)
predicate (E.BUSINESSENTITYID=S.BUSINESSENTITYID)
search_condition (E.BUSINESSENTITYID=S.BUSINESSENTITYID)
join_on (JOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
join_part (JOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
table_source_item_joined (PERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
table_source (PERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
table_sources (PERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
query_specification (SELECTLASTNAME,FIRSTNAMEFROMPERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
query_expression (SELECTLASTNAME,FIRSTNAMEFROMPERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID)
select_statement (SELECTLASTNAME,FIRSTNAMEFROMPERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID;)
select_statement_standalone (SELECTLASTNAME,FIRSTNAMEFROMPERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID;)
dml_clause (SELECTLASTNAME,FIRSTNAMEFROMPERSON.PERSONCINNERJOINHUMANRESOURCES.EMPLOYEEEONC.BUSINESSENTITYID=E.BUSINESSENTITYIDJOINSALES.SALESPERSONSONE.BUSINESSENTITYID=S.BUSINESSENTITYID;)

*/
