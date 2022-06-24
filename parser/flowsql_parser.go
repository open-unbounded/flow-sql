// Code generated from D:/nextgen/golang/flow-sql\FlowSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FlowSql

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FlowSqlParser struct {
	*antlr.BaseParser
}

var flowsqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func flowsqlParserInit() {
	staticData := &flowsqlParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'*'", "'/'", "'-'", "'+'", "'('", "')'", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "SELECT", "FROM", "WHERE", "ORDER_BY",
		"AND", "AS", "ID", "NUMBER", "CP", "WS",
	}
	staticData.ruleNames = []string{
		"sql", "where", "expr", "field", "fields", "alias", "orderby",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 77, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 20, 8, 0,
		1, 0, 3, 0, 23, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 45, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 53, 8,
		2, 10, 2, 12, 2, 56, 9, 2, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 4, 1, 4, 1, 4,
		5, 4, 65, 8, 4, 10, 4, 12, 4, 68, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 0, 1, 4, 7, 0, 2, 4, 6, 8, 10, 12, 0, 2, 1, 0, 2, 3, 1,
		0, 4, 5, 78, 0, 14, 1, 0, 0, 0, 2, 26, 1, 0, 0, 0, 4, 44, 1, 0, 0, 0, 6,
		59, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 69, 1, 0, 0, 0, 12, 73, 1, 0, 0,
		0, 14, 15, 5, 9, 0, 0, 15, 16, 3, 8, 4, 0, 16, 17, 5, 10, 0, 0, 17, 19,
		5, 15, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0,
		20, 22, 1, 0, 0, 0, 21, 23, 3, 12, 6, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1,
		0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 5, 1, 0, 0, 25, 1, 1, 0, 0, 0, 26,
		27, 5, 11, 0, 0, 27, 28, 5, 15, 0, 0, 28, 29, 5, 17, 0, 0, 29, 34, 3, 4,
		2, 0, 30, 31, 5, 13, 0, 0, 31, 33, 3, 4, 2, 0, 32, 30, 1, 0, 0, 0, 33,
		36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 3, 1, 0, 0,
		0, 36, 34, 1, 0, 0, 0, 37, 38, 6, 2, -1, 0, 38, 45, 5, 16, 0, 0, 39, 45,
		5, 15, 0, 0, 40, 41, 5, 6, 0, 0, 41, 42, 3, 4, 2, 0, 42, 43, 5, 7, 0, 0,
		43, 45, 1, 0, 0, 0, 44, 37, 1, 0, 0, 0, 44, 39, 1, 0, 0, 0, 44, 40, 1,
		0, 0, 0, 45, 54, 1, 0, 0, 0, 46, 47, 10, 5, 0, 0, 47, 48, 7, 0, 0, 0, 48,
		53, 3, 4, 2, 6, 49, 50, 10, 4, 0, 0, 50, 51, 7, 1, 0, 0, 51, 53, 3, 4,
		2, 5, 52, 46, 1, 0, 0, 0, 52, 49, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 5, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		57, 60, 5, 15, 0, 0, 58, 60, 3, 10, 5, 0, 59, 57, 1, 0, 0, 0, 59, 58, 1,
		0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 66, 3, 6, 3, 0, 62, 63, 5, 8, 0, 0, 63,
		65, 3, 6, 3, 0, 64, 62, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0,
		0, 66, 67, 1, 0, 0, 0, 67, 9, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 3,
		4, 2, 0, 70, 71, 5, 14, 0, 0, 71, 72, 5, 15, 0, 0, 72, 11, 1, 0, 0, 0,
		73, 74, 5, 12, 0, 0, 74, 75, 5, 15, 0, 0, 75, 13, 1, 0, 0, 0, 8, 19, 22,
		34, 44, 52, 54, 59, 66,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FlowSqlParserInit initializes any static state used to implement FlowSqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFlowSqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FlowSqlParserInit() {
	staticData := &flowsqlParserStaticData
	staticData.once.Do(flowsqlParserInit)
}

// NewFlowSqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFlowSqlParser(input antlr.TokenStream) *FlowSqlParser {
	FlowSqlParserInit()
	this := new(FlowSqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &flowsqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FlowSql.g4"

	return this
}

// FlowSqlParser tokens.
const (
	FlowSqlParserEOF      = antlr.TokenEOF
	FlowSqlParserT__0     = 1
	FlowSqlParserT__1     = 2
	FlowSqlParserT__2     = 3
	FlowSqlParserT__3     = 4
	FlowSqlParserT__4     = 5
	FlowSqlParserT__5     = 6
	FlowSqlParserT__6     = 7
	FlowSqlParserT__7     = 8
	FlowSqlParserSELECT   = 9
	FlowSqlParserFROM     = 10
	FlowSqlParserWHERE    = 11
	FlowSqlParserORDER_BY = 12
	FlowSqlParserAND      = 13
	FlowSqlParserAS       = 14
	FlowSqlParserID       = 15
	FlowSqlParserNUMBER   = 16
	FlowSqlParserCP       = 17
	FlowSqlParserWS       = 18
)

// FlowSqlParser rules.
const (
	FlowSqlParserRULE_sql     = 0
	FlowSqlParserRULE_where   = 1
	FlowSqlParserRULE_expr    = 2
	FlowSqlParserRULE_field   = 3
	FlowSqlParserRULE_fields  = 4
	FlowSqlParserRULE_alias   = 5
	FlowSqlParserRULE_orderby = 6
)

// ISqlContext is an interface to support dynamic dispatch.
type ISqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSqlContext differentiates from other interfaces.
	IsSqlContext()
}

type SqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlContext() *SqlContext {
	var p = new(SqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_sql
	return p
}

func (*SqlContext) IsSqlContext() {}

func NewSqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlContext {
	var p = new(SqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_sql

	return p
}

func (s *SqlContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlContext) SELECT() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserSELECT, 0)
}

func (s *SqlContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *SqlContext) FROM() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserFROM, 0)
}

func (s *SqlContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserID, 0)
}

func (s *SqlContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *SqlContext) Orderby() IOrderbyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderbyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderbyContext)
}

func (s *SqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterSql(s)
	}
}

func (s *SqlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitSql(s)
	}
}

func (s *SqlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitSql(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Sql() (localctx ISqlContext) {
	this := p
	_ = this

	localctx = NewSqlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FlowSqlParserRULE_sql)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Match(FlowSqlParserSELECT)
	}
	{
		p.SetState(15)
		p.Fields()
	}
	{
		p.SetState(16)
		p.Match(FlowSqlParserFROM)
	}
	{
		p.SetState(17)
		p.Match(FlowSqlParserID)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlowSqlParserWHERE {
		{
			p.SetState(18)
			p.Where()
		}

	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlowSqlParserORDER_BY {
		{
			p.SetState(21)
			p.Orderby()
		}

	}
	{
		p.SetState(24)
		p.Match(FlowSqlParserT__0)
	}

	return localctx
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_where
	return p
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserWHERE, 0)
}

func (s *WhereContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserID, 0)
}

func (s *WhereContext) CP() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserCP, 0)
}

func (s *WhereContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *WhereContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhereContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(FlowSqlParserAND)
}

func (s *WhereContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(FlowSqlParserAND, i)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterWhere(s)
	}
}

func (s *WhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitWhere(s)
	}
}

func (s *WhereContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitWhere(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Where() (localctx IWhereContext) {
	this := p
	_ = this

	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FlowSqlParserRULE_where)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(FlowSqlParserWHERE)
	}
	{
		p.SetState(27)
		p.Match(FlowSqlParserID)
	}
	{
		p.SetState(28)
		p.Match(FlowSqlParserCP)
	}
	{
		p.SetState(29)
		p.expr(0)
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FlowSqlParserAND {
		{
			p.SetState(30)
			p.Match(FlowSqlParserAND)
		}
		{
			p.SetState(31)
			p.expr(0)
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	*ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	*ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	*ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *FlowSqlParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, FlowSqlParserRULE_expr, _p)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FlowSqlParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(38)
			p.Match(FlowSqlParserNUMBER)
		}

	case FlowSqlParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(39)
			p.Match(FlowSqlParserID)
		}

	case FlowSqlParserT__5:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(FlowSqlParserT__5)
		}
		{
			p.SetState(41)
			p.expr(0)
		}
		{
			p.SetState(42)
			p.Match(FlowSqlParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FlowSqlParserRULE_expr)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(47)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == FlowSqlParserT__1 || _la == FlowSqlParserT__2) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(48)
					p.expr(6)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FlowSqlParserRULE_expr)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(50)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == FlowSqlParserT__3 || _la == FlowSqlParserT__4) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(51)
					p.expr(5)
				}

			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserID, 0)
}

func (s *FieldContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FlowSqlParserRULE_field)

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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(FlowSqlParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Alias()
		}

	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitFields(s)
	}
}

func (s *FieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitFields(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FlowSqlParserRULE_fields)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Field()
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FlowSqlParserT__7 {
		{
			p.SetState(62)
			p.Match(FlowSqlParserT__7)
		}
		{
			p.SetState(63)
			p.Field()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AliasContext) AS() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserAS, 0)
}

func (s *AliasContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserID, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Alias() (localctx IAliasContext) {
	this := p
	_ = this

	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FlowSqlParserRULE_alias)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.expr(0)
	}
	{
		p.SetState(70)
		p.Match(FlowSqlParserAS)
	}
	{
		p.SetState(71)
		p.Match(FlowSqlParserID)
	}

	return localctx
}

// IOrderbyContext is an interface to support dynamic dispatch.
type IOrderbyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderbyContext differentiates from other interfaces.
	IsOrderbyContext()
}

type OrderbyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderbyContext() *OrderbyContext {
	var p = new(OrderbyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlowSqlParserRULE_orderby
	return p
}

func (*OrderbyContext) IsOrderbyContext() {}

func NewOrderbyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderbyContext {
	var p = new(OrderbyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowSqlParserRULE_orderby

	return p
}

func (s *OrderbyContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderbyContext) ORDER_BY() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserORDER_BY, 0)
}

func (s *OrderbyContext) ID() antlr.TerminalNode {
	return s.GetToken(FlowSqlParserID, 0)
}

func (s *OrderbyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderbyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderbyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.EnterOrderby(s)
	}
}

func (s *OrderbyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowSqlListener); ok {
		listenerT.ExitOrderby(s)
	}
}

func (s *OrderbyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowSqlVisitor:
		return t.VisitOrderby(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowSqlParser) Orderby() (localctx IOrderbyContext) {
	this := p
	_ = this

	localctx = NewOrderbyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FlowSqlParserRULE_orderby)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(FlowSqlParserORDER_BY)
	}
	{
		p.SetState(74)
		p.Match(FlowSqlParserID)
	}

	return localctx
}

func (p *FlowSqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FlowSqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
