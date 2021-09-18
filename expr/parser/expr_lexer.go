// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 49, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 6, 6, 29, 10, 6, 13, 6, 14, 6, 30, 3, 7, 6, 7, 34, 10, 7, 13,
	7, 14, 7, 35, 3, 8, 5, 8, 39, 10, 8, 3, 8, 3, 8, 3, 9, 6, 9, 44, 10, 9,
	13, 9, 14, 9, 45, 3, 9, 3, 9, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 3, 2, 5, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 4,
	2, 11, 11, 34, 34, 2, 52, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 21, 3, 2, 2, 2, 7,
	23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 28, 3, 2, 2, 2, 13, 33, 3, 2, 2,
	2, 15, 38, 3, 2, 2, 2, 17, 43, 3, 2, 2, 2, 19, 20, 7, 44, 2, 2, 20, 4,
	3, 2, 2, 2, 21, 22, 7, 49, 2, 2, 22, 6, 3, 2, 2, 2, 23, 24, 7, 45, 2, 2,
	24, 8, 3, 2, 2, 2, 25, 26, 7, 47, 2, 2, 26, 10, 3, 2, 2, 2, 27, 29, 9,
	2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30,
	31, 3, 2, 2, 2, 31, 12, 3, 2, 2, 2, 32, 34, 9, 3, 2, 2, 33, 32, 3, 2, 2,
	2, 34, 35, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 14,
	3, 2, 2, 2, 37, 39, 7, 15, 2, 2, 38, 37, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2,
	39, 40, 3, 2, 2, 2, 40, 41, 7, 12, 2, 2, 41, 16, 3, 2, 2, 2, 42, 44, 9,
	4, 2, 2, 43, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45,
	46, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 8, 9, 2, 2, 48, 18, 3, 2, 2,
	2, 7, 2, 30, 35, 38, 45, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'/'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "MUL", "DIV", "ADD", "SUB", "ID", "INT", "NEWLINE", "WS",
}

var lexerRuleNames = []string{
	"MUL", "DIV", "ADD", "SUB", "ID", "INT", "NEWLINE", "WS",
}

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewExprLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ExprLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewExprLexer(input antlr.CharStream) *ExprLexer {
	l := new(ExprLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerMUL     = 1
	ExprLexerDIV     = 2
	ExprLexerADD     = 3
	ExprLexerSUB     = 4
	ExprLexerID      = 5
	ExprLexerINT     = 6
	ExprLexerNEWLINE = 7
	ExprLexerWS      = 8
)
