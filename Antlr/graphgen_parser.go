// Code generated from Antlr/GraphGen.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GraphGen

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GraphGenParser struct {
	*antlr.BaseParser
}

var GraphGenParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func graphgenParserInit() {
	staticData := &GraphGenParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'person'", "'group'", "'relationship'", "'filter by'", "'draw'",
		"','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "FILTER_BY", "DRAW",
		"VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT", "ERROR",
	}
	staticData.RuleNames = []string{
		"program", "declarations", "var_type", "relationship_definitions", "draw_command",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 57, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 1, 0, 5, 0, 18, 8, 0,
		10, 0, 12, 0, 21, 9, 0, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 5, 1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 5, 3, 43, 8, 3, 10, 3, 12, 3, 46, 9, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 55, 8, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8,
		0, 1, 1, 0, 2, 4, 57, 0, 13, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 37, 1, 0,
		0, 0, 6, 39, 1, 0, 0, 0, 8, 50, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 10,
		1, 0, 0, 0, 12, 15, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0,
		14, 19, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 16, 18, 3, 6, 3, 0, 17, 16, 1,
		0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20,
		25, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 24, 3, 8, 4, 0, 23, 22, 1, 0, 0,
		0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 1, 1,
		0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 3, 4, 2, 0, 29, 34, 5, 8, 0, 0, 30,
		31, 5, 7, 0, 0, 31, 33, 5, 8, 0, 0, 32, 30, 1, 0, 0, 0, 33, 36, 1, 0, 0,
		0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 3, 1, 0, 0, 0, 36, 34, 1,
		0, 0, 0, 37, 38, 7, 0, 0, 0, 38, 5, 1, 0, 0, 0, 39, 44, 5, 8, 0, 0, 40,
		41, 5, 7, 0, 0, 41, 43, 5, 8, 0, 0, 42, 40, 1, 0, 0, 0, 43, 46, 1, 0, 0,
		0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44,
		1, 0, 0, 0, 47, 48, 5, 8, 0, 0, 48, 49, 5, 8, 0, 0, 49, 7, 1, 0, 0, 0,
		50, 51, 5, 6, 0, 0, 51, 54, 5, 8, 0, 0, 52, 53, 5, 5, 0, 0, 53, 55, 5,
		8, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 9, 1, 0, 0, 0, 6,
		13, 19, 25, 34, 44, 54,
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

// GraphGenParserInit initializes any static state used to implement GraphGenParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGraphGenParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GraphGenParserInit() {
	staticData := &GraphGenParserStaticData
	staticData.once.Do(graphgenParserInit)
}

// NewGraphGenParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGraphGenParser(input antlr.TokenStream) *GraphGenParser {
	GraphGenParserInit()
	this := new(GraphGenParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GraphGenParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GraphGen.g4"

	return this
}

// GraphGenParser tokens.
const (
	GraphGenParserEOF                = antlr.TokenEOF
	GraphGenParserCOMMENT            = 1
	GraphGenParserPERSON             = 2
	GraphGenParserGROUP              = 3
	GraphGenParserRELATIONSHIP       = 4
	GraphGenParserFILTER_BY          = 5
	GraphGenParserDRAW               = 6
	GraphGenParserVIRGULA            = 7
	GraphGenParserIDENT              = 8
	GraphGenParserWS                 = 9
	GraphGenParserERROR_OPEN_COMMENT = 10
	GraphGenParserERROR              = 11
)

// GraphGenParser rules.
const (
	GraphGenParserRULE_program                  = 0
	GraphGenParserRULE_declarations             = 1
	GraphGenParserRULE_var_type                 = 2
	GraphGenParserRULE_relationship_definitions = 3
	GraphGenParserRULE_draw_command             = 4
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclarations() []IDeclarationsContext
	Declarations(i int) IDeclarationsContext
	AllRelationship_definitions() []IRelationship_definitionsContext
	Relationship_definitions(i int) IRelationship_definitionsContext
	AllDraw_command() []IDraw_commandContext
	Draw_command(i int) IDraw_commandContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphGenParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDeclarations() []IDeclarationsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationsContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationsContext); ok {
			tst[i] = t.(IDeclarationsContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declarations(i int) IDeclarationsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
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

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) AllRelationship_definitions() []IRelationship_definitionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationship_definitionsContext); ok {
			len++
		}
	}

	tst := make([]IRelationship_definitionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationship_definitionsContext); ok {
			tst[i] = t.(IRelationship_definitionsContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Relationship_definitions(i int) IRelationship_definitionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationship_definitionsContext); ok {
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

	return t.(IRelationship_definitionsContext)
}

func (s *ProgramContext) AllDraw_command() []IDraw_commandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDraw_commandContext); ok {
			len++
		}
	}

	tst := make([]IDraw_commandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDraw_commandContext); ok {
			tst[i] = t.(IDraw_commandContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Draw_command(i int) IDraw_commandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDraw_commandContext); ok {
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

	return t.(IDraw_commandContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphGenVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphGenParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphGenParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28) != 0 {
		{
			p.SetState(10)
			p.Declarations()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserIDENT {
		{
			p.SetState(16)
			p.Relationship_definitions()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserDRAW {
		{
			p.SetState(22)
			p.Draw_command()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_declarations
	return p
}

func InitEmptyDeclarationsContext(p *DeclarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_declarations
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphGenParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *DeclarationsContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserIDENT)
}

func (s *DeclarationsContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserIDENT, i)
}

func (s *DeclarationsContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserVIRGULA)
}

func (s *DeclarationsContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserVIRGULA, i)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.EnterDeclarations(s)
	}
}

func (s *DeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.ExitDeclarations(s)
	}
}

func (s *DeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphGenVisitor:
		return t.VisitDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphGenParser) Declarations() (localctx IDeclarationsContext) {
	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphGenParserRULE_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Var_type()
	}
	{
		p.SetState(29)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserVIRGULA {
		{
			p.SetState(30)
			p.Match(GraphGenParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(GraphGenParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PERSON() antlr.TerminalNode
	GROUP() antlr.TerminalNode
	RELATIONSHIP() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphGenParserRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) PERSON() antlr.TerminalNode {
	return s.GetToken(GraphGenParserPERSON, 0)
}

func (s *Var_typeContext) GROUP() antlr.TerminalNode {
	return s.GetToken(GraphGenParserGROUP, 0)
}

func (s *Var_typeContext) RELATIONSHIP() antlr.TerminalNode {
	return s.GetToken(GraphGenParserRELATIONSHIP, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphGenVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphGenParser) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphGenParserRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationship_definitionsContext is an interface to support dynamic dispatch.
type IRelationship_definitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRelation returns the relation token.
	GetRelation() antlr.Token

	// GetRelated returns the related token.
	GetRelated() antlr.Token

	// SetRelation sets the relation token.
	SetRelation(antlr.Token)

	// SetRelated sets the related token.
	SetRelated(antlr.Token)

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsRelationship_definitionsContext differentiates from other interfaces.
	IsRelationship_definitionsContext()
}

type Relationship_definitionsContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	relation antlr.Token
	related  antlr.Token
}

func NewEmptyRelationship_definitionsContext() *Relationship_definitionsContext {
	var p = new(Relationship_definitionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_relationship_definitions
	return p
}

func InitEmptyRelationship_definitionsContext(p *Relationship_definitionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_relationship_definitions
}

func (*Relationship_definitionsContext) IsRelationship_definitionsContext() {}

func NewRelationship_definitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relationship_definitionsContext {
	var p = new(Relationship_definitionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphGenParserRULE_relationship_definitions

	return p
}

func (s *Relationship_definitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Relationship_definitionsContext) GetRelation() antlr.Token { return s.relation }

func (s *Relationship_definitionsContext) GetRelated() antlr.Token { return s.related }

func (s *Relationship_definitionsContext) SetRelation(v antlr.Token) { s.relation = v }

func (s *Relationship_definitionsContext) SetRelated(v antlr.Token) { s.related = v }

func (s *Relationship_definitionsContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserIDENT)
}

func (s *Relationship_definitionsContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserIDENT, i)
}

func (s *Relationship_definitionsContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserVIRGULA)
}

func (s *Relationship_definitionsContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserVIRGULA, i)
}

func (s *Relationship_definitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relationship_definitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relationship_definitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.EnterRelationship_definitions(s)
	}
}

func (s *Relationship_definitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.ExitRelationship_definitions(s)
	}
}

func (s *Relationship_definitionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphGenVisitor:
		return t.VisitRelationship_definitions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphGenParser) Relationship_definitions() (localctx IRelationship_definitionsContext) {
	localctx = NewRelationship_definitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphGenParserRULE_relationship_definitions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserVIRGULA {
		{
			p.SetState(40)
			p.Match(GraphGenParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(GraphGenParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)

		var _m = p.Match(GraphGenParserIDENT)

		localctx.(*Relationship_definitionsContext).relation = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)

		var _m = p.Match(GraphGenParserIDENT)

		localctx.(*Relationship_definitionsContext).related = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDraw_commandContext is an interface to support dynamic dispatch.
type IDraw_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilter returns the filter token.
	GetFilter() antlr.Token

	// SetFilter sets the filter token.
	SetFilter(antlr.Token)

	// Getter signatures
	DRAW() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	FILTER_BY() antlr.TerminalNode

	// IsDraw_commandContext differentiates from other interfaces.
	IsDraw_commandContext()
}

type Draw_commandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	filter antlr.Token
}

func NewEmptyDraw_commandContext() *Draw_commandContext {
	var p = new(Draw_commandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_draw_command
	return p
}

func InitEmptyDraw_commandContext(p *Draw_commandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_draw_command
}

func (*Draw_commandContext) IsDraw_commandContext() {}

func NewDraw_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Draw_commandContext {
	var p = new(Draw_commandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphGenParserRULE_draw_command

	return p
}

func (s *Draw_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Draw_commandContext) GetFilter() antlr.Token { return s.filter }

func (s *Draw_commandContext) SetFilter(v antlr.Token) { s.filter = v }

func (s *Draw_commandContext) DRAW() antlr.TerminalNode {
	return s.GetToken(GraphGenParserDRAW, 0)
}

func (s *Draw_commandContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserIDENT)
}

func (s *Draw_commandContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserIDENT, i)
}

func (s *Draw_commandContext) FILTER_BY() antlr.TerminalNode {
	return s.GetToken(GraphGenParserFILTER_BY, 0)
}

func (s *Draw_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Draw_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Draw_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.EnterDraw_command(s)
	}
}

func (s *Draw_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.ExitDraw_command(s)
	}
}

func (s *Draw_commandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphGenVisitor:
		return t.VisitDraw_command(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphGenParser) Draw_command() (localctx IDraw_commandContext) {
	localctx = NewDraw_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphGenParserRULE_draw_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(GraphGenParserDRAW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphGenParserFILTER_BY {
		{
			p.SetState(52)
			p.Match(GraphGenParserFILTER_BY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)

			var _m = p.Match(GraphGenParserIDENT)

			localctx.(*Draw_commandContext).filter = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
