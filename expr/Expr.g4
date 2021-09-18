grammar Expr;

@parser::header {
import (
    "os"
)
}

@parser::members {

func eval(left int, op antlr.Token, right int) int {
    if   (op.GetText() == "*") {
        return left * right
    } else if (op.GetText() == "+") {
        return left + right
    } else if (op.GetText() == "-") {
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
    : a=e op=('+'|'-') b=e  {
                $v = eval($a.v, $op, $b.v)
                fmt.Fprintf(os.Stdout, "got args=%d %d\n", $a.v, $b.v)
                }
    | INT                   {
                $v = $INT.int
                fmt.Fprintf(os.Stdout, "got number=%d\n", $v)
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