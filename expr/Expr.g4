grammar Expr;

@parser::header {
import (
    "os"
)
}

@parser::members {

func eval(left int, op antlr.Token, right int) int {
    if (op.GetTokenType() == ExprParserMUL) {
        return left * right
    } else if (op.GetTokenType() == ExprParserDIV) {
        return left / right
    } else if (op.GetTokenType() == ExprParserADD) {
        return left + right
    } else if (op.GetTokenType() == ExprParserSUB) {
        return left - right
    } else {
        return 0
    }
}
}

stat:   e NEWLINE 
    |   NEWLINE                   
    ;

e returns [int v]
    : a=e op=('*'|'/') b=e  {
                $v = eval($a.v, $op, $b.v)
                // fmt.Fprintf(os.Stdout, "%d %s %d = %d\n", $a.v, $op.GetText(), $b.v, $v)
                }
    | a=e op=('+'|'-') b=e  {
                $v = eval($a.v, $op, $b.v)
                // fmt.Fprintf(os.Stdout, "%d %s %d = %d\n", $a.v, $op.GetText(), $b.v, $v)
                }
    | INT                   {
                $v = $INT.int
                // fmt.Fprintf(os.Stdout, "got number=%d\n", $v)
                }    
    ; 

MUL : '*' ;
DIV : '/' ;
ADD : '+' ;
SUB : '-' ;

ID  :   [a-zA-Z]+ ;      // match identifiers
INT :   [0-9]+ ;         // match integers
NEWLINE:'\r'? '\n' ;     // return newlines to parser (is end-statement signal)
WS  :   [ \t]+ -> skip ; // toss out whitespace
