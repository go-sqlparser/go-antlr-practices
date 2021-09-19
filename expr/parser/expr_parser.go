// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 34, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 13, 10,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3, 2, 3, 4,
	4, 2, 4, 2, 4, 3, 2, 3, 4, 3, 2, 5, 6, 2, 34, 2, 12, 3, 2, 2, 2, 4, 14,
	3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 9, 2, 2, 8, 9, 8, 2, 1, 2, 9, 13,
	3, 2, 2, 2, 10, 11, 7, 9, 2, 2, 11, 13, 8, 2, 1, 2, 12, 6, 3, 2, 2, 2,
	12, 10, 3, 2, 2, 2, 13, 3, 3, 2, 2, 2, 14, 15, 8, 3, 1, 2, 15, 16, 7, 8,
	2, 2, 16, 17, 8, 3, 1, 2, 17, 30, 3, 2, 2, 2, 18, 19, 12, 5, 2, 2, 19,
	20, 9, 2, 2, 2, 20, 21, 5, 4, 3, 6, 21, 22, 8, 3, 1, 2, 22, 29, 3, 2, 2,
	2, 23, 24, 12, 4, 2, 2, 24, 25, 9, 3, 2, 2, 25, 26, 5, 4, 3, 5, 26, 27,
	8, 3, 1, 2, 27, 29, 3, 2, 2, 2, 28, 18, 3, 2, 2, 2, 28, 23, 3, 2, 2, 2,
	29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2,
	2, 2, 32, 30, 3, 2, 2, 2, 5, 12, 28, 30,
}
var literalNames = []string{
	"", "'*'", "'/'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "MUL", "DIV", "ADD", "SUB", "ID", "INT", "NEWLINE", "WS",
}

var ruleNames = []string{
	"stat", "e",
}

type ExprParser struct {
	*antlr.BaseParser
}

// NewExprParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ExprParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewExprParser(input antlr.TokenStream) *ExprParser {
	this := new(ExprParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expr.g4"

	return this
}

func eval(left int, op antlr.Token, right int) int {
	if op.GetTokenType() == ExprParserMUL {
		return left * right
	} else if op.GetTokenType() == ExprParserDIV {
		return left / right
	} else if op.GetTokenType() == ExprParserADD {
		return left + right
	} else if op.GetTokenType() == ExprParserSUB {
		return left - right
	} else {
		return 0
	}
}

// ExprParser tokens.
const (
	ExprParserEOF     = antlr.TokenEOF
	ExprParserMUL     = 1
	ExprParserDIV     = 2
	ExprParserADD     = 3
	ExprParserSUB     = 4
	ExprParserID      = 5
	ExprParserINT     = 6
	ExprParserNEWLINE = 7
	ExprParserWS      = 8
)

// ExprParser rules.
const (
	ExprParserRULE_stat = 0
	ExprParserRULE_e    = 1
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetA returns the a rule contexts.
	GetA() IEContext

	// SetA sets the a rule contexts.
	SetA(IEContext)

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
	a      IEContext
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) GetA() IEContext { return s.a }

func (s *StatContext) SetA(v IEContext) { s.a = v }

func (s *StatContext) GetV() int { return s.v }

func (s *StatContext) SetV(v int) { s.v = v }

func (s *StatContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ExprParserNEWLINE, 0)
}

func (s *StatContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ExprParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(10)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(4)

			var _x = p.e(0)

			localctx.(*StatContext).a = _x
		}
		{
			p.SetState(5)
			p.Match(ExprParserNEWLINE)
		}
		localctx.(*StatContext).v = localctx.(*StatContext).GetA().GetV()

	case ExprParserNEWLINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(8)
			p.Match(ExprParserNEWLINE)
		}
		localctx.(*StatContext).v = 0

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetA returns the a rule contexts.
	GetA() IEContext

	// GetB returns the b rule contexts.
	GetB() IEContext

	// SetA sets the a rule contexts.
	SetA(IEContext)

	// SetB sets the b rule contexts.
	SetB(IEContext)

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
	a      IEContext
	_INT   antlr.Token
	op     antlr.Token
	b      IEContext
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) Get_INT() antlr.Token { return s._INT }

func (s *EContext) GetOp() antlr.Token { return s.op }

func (s *EContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *EContext) SetOp(v antlr.Token) { s.op = v }

func (s *EContext) GetA() IEContext { return s.a }

func (s *EContext) GetB() IEContext { return s.b }

func (s *EContext) SetA(v IEContext) { s.a = v }

func (s *EContext) SetB(v IEContext) { s.b = v }

func (s *EContext) GetV() int { return s.v }

func (s *EContext) SetV(v int) { s.v = v }

func (s *EContext) INT() antlr.TerminalNode {
	return s.GetToken(ExprParserINT, 0)
}

func (s *EContext) AllE() []IEContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEContext)(nil)).Elem())
	var tst = make([]IEContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEContext)
		}
	}

	return tst
}

func (s *EContext) E(i int) IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *EContext) MUL() antlr.TerminalNode {
	return s.GetToken(ExprParserMUL, 0)
}

func (s *EContext) DIV() antlr.TerminalNode {
	return s.GetToken(ExprParserDIV, 0)
}

func (s *EContext) ADD() antlr.TerminalNode {
	return s.GetToken(ExprParserADD, 0)
}

func (s *EContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExprParserSUB, 0)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *ExprParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *ExprParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ExprParserRULE_e, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)

		var _m = p.Match(ExprParserINT)

		localctx.(*EContext)._INT = _m
	}

	localctx.(*EContext).v = (func() int {
		if localctx.(*EContext).Get_INT() == nil {
			return 0
		} else {
			i, _ := strconv.Atoi(localctx.(*EContext).Get_INT().GetText())
			return i
		}
	}())
	// fmt.Fprintf(os.Stdout, "got number=%d\n", localctx.(*EContext).v)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEContext(p, _parentctx, _parentState)
				localctx.(*EContext).a = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_e)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserMUL || _la == ExprParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)

					var _x = p.e(4)

					localctx.(*EContext).b = _x
				}

				localctx.(*EContext).v = eval(localctx.(*EContext).GetA().GetV(), localctx.(*EContext).GetOp(), localctx.(*EContext).GetB().GetV())
				// fmt.Fprintf(os.Stdout, "%d %s %d = %d\n", localctx.(*EContext).GetA().GetV(), localctx.(*EContext).GetOp().GetText(), localctx.(*EContext).GetB().GetV(), localctx.(*EContext).v)

			case 2:
				localctx = NewEContext(p, _parentctx, _parentState)
				localctx.(*EContext).a = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_e)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(22)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserADD || _la == ExprParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(23)

					var _x = p.e(3)

					localctx.(*EContext).b = _x
				}

				localctx.(*EContext).v = eval(localctx.(*EContext).GetA().GetV(), localctx.(*EContext).GetOp(), localctx.(*EContext).GetB().GetV())
				// fmt.Fprintf(os.Stdout, "%d %s %d = %d\n", localctx.(*EContext).GetA().GetV(), localctx.(*EContext).GetOp().GetText(), localctx.(*EContext).GetB().GetV(), localctx.(*EContext).v)

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExprParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
