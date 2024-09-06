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
		"", "", "'person'", "'group'", "'relationship'", "'subgroup of'", "'draw'",
		"'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW",
		"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT",
		"ERROR", "TYPE", "RELATION",
	}
	staticData.RuleNames = []string{
		"program", "declarations", "subgroups_definitions", "relationship_definitions",
		"draw_command",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 59, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 20, 8, 1,
		10, 1, 12, 1, 23, 9, 1, 1, 2, 1, 2, 1, 2, 5, 2, 28, 8, 2, 10, 2, 12, 2,
		31, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 39, 8, 3, 10, 3, 12,
		3, 42, 9, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 52,
		8, 4, 10, 4, 12, 4, 55, 9, 4, 1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8,
		0, 0, 57, 0, 10, 1, 0, 0, 0, 2, 15, 1, 0, 0, 0, 4, 24, 1, 0, 0, 0, 6, 35,
		1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 12, 3, 4, 2, 0,
		12, 13, 3, 6, 3, 0, 13, 14, 3, 8, 4, 0, 14, 1, 1, 0, 0, 0, 15, 16, 5, 14,
		0, 0, 16, 21, 5, 10, 0, 0, 17, 18, 5, 9, 0, 0, 18, 20, 5, 10, 0, 0, 19,
		17, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0,
		0, 22, 3, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24, 29, 5, 10, 0, 0, 25, 26,
		5, 9, 0, 0, 26, 28, 5, 10, 0, 0, 27, 25, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0,
		29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1,
		0, 0, 0, 32, 33, 5, 5, 0, 0, 33, 34, 5, 10, 0, 0, 34, 5, 1, 0, 0, 0, 35,
		40, 5, 10, 0, 0, 36, 37, 5, 9, 0, 0, 37, 39, 5, 10, 0, 0, 38, 36, 1, 0,
		0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43,
		1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 44, 5, 15, 0, 0, 44, 45, 5, 10, 0,
		0, 45, 7, 1, 0, 0, 0, 46, 47, 5, 6, 0, 0, 47, 48, 5, 7, 0, 0, 48, 53, 5,
		10, 0, 0, 49, 50, 5, 9, 0, 0, 50, 52, 5, 15, 0, 0, 51, 49, 1, 0, 0, 0,
		52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1,
		0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 8, 0, 0, 57, 9, 1, 0, 0, 0, 4,
		21, 29, 40, 53,
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
	GraphGenParserSUBGROUP_OF        = 5
	GraphGenParserDRAW               = 6
	GraphGenParserOPENPAR            = 7
	GraphGenParserCLOSEPAR           = 8
	GraphGenParserVIRGULA            = 9
	GraphGenParserIDENT              = 10
	GraphGenParserWS                 = 11
	GraphGenParserERROR_OPEN_COMMENT = 12
	GraphGenParserERROR              = 13
	GraphGenParserTYPE               = 14
	GraphGenParserRELATION           = 15
)

// GraphGenParser rules.
const (
	GraphGenParserRULE_program                  = 0
	GraphGenParserRULE_declarations             = 1
	GraphGenParserRULE_subgroups_definitions    = 2
	GraphGenParserRULE_relationship_definitions = 3
	GraphGenParserRULE_draw_command             = 4
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declarations() IDeclarationsContext
	Subgroups_definitions() ISubgroups_definitionsContext
	Relationship_definitions() IRelationship_definitionsContext
	Draw_command() IDraw_commandContext

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

func (s *ProgramContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) Subgroups_definitions() ISubgroups_definitionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubgroups_definitionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubgroups_definitionsContext)
}

func (s *ProgramContext) Relationship_definitions() IRelationship_definitionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationship_definitionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationship_definitionsContext)
}

func (s *ProgramContext) Draw_command() IDraw_commandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDraw_commandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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

func (p *GraphGenParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphGenParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Declarations()
	}
	{
		p.SetState(11)
		p.Subgroups_definitions()
	}
	{
		p.SetState(12)
		p.Relationship_definitions()
	}
	{
		p.SetState(13)
		p.Draw_command()
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
	TYPE() antlr.TerminalNode
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

func (s *DeclarationsContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GraphGenParserTYPE, 0)
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

func (p *GraphGenParser) Declarations() (localctx IDeclarationsContext) {
	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphGenParserRULE_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.Match(GraphGenParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(16)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserVIRGULA {
		{
			p.SetState(17)
			p.Match(GraphGenParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)
			p.Match(GraphGenParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(23)
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

// ISubgroups_definitionsContext is an interface to support dynamic dispatch.
type ISubgroups_definitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	SUBGROUP_OF() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsSubgroups_definitionsContext differentiates from other interfaces.
	IsSubgroups_definitionsContext()
}

type Subgroups_definitionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubgroups_definitionsContext() *Subgroups_definitionsContext {
	var p = new(Subgroups_definitionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_subgroups_definitions
	return p
}

func InitEmptySubgroups_definitionsContext(p *Subgroups_definitionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphGenParserRULE_subgroups_definitions
}

func (*Subgroups_definitionsContext) IsSubgroups_definitionsContext() {}

func NewSubgroups_definitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subgroups_definitionsContext {
	var p = new(Subgroups_definitionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphGenParserRULE_subgroups_definitions

	return p
}

func (s *Subgroups_definitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Subgroups_definitionsContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserIDENT)
}

func (s *Subgroups_definitionsContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserIDENT, i)
}

func (s *Subgroups_definitionsContext) SUBGROUP_OF() antlr.TerminalNode {
	return s.GetToken(GraphGenParserSUBGROUP_OF, 0)
}

func (s *Subgroups_definitionsContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserVIRGULA)
}

func (s *Subgroups_definitionsContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserVIRGULA, i)
}

func (s *Subgroups_definitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subgroups_definitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subgroups_definitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.EnterSubgroups_definitions(s)
	}
}

func (s *Subgroups_definitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphGenListener); ok {
		listenerT.ExitSubgroups_definitions(s)
	}
}

func (p *GraphGenParser) Subgroups_definitions() (localctx ISubgroups_definitionsContext) {
	localctx = NewSubgroups_definitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphGenParserRULE_subgroups_definitions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserVIRGULA {
		{
			p.SetState(25)
			p.Match(GraphGenParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(GraphGenParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(GraphGenParserSUBGROUP_OF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(GraphGenParserIDENT)
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

// IRelationship_definitionsContext is an interface to support dynamic dispatch.
type IRelationship_definitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	RELATION() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode

	// IsRelationship_definitionsContext differentiates from other interfaces.
	IsRelationship_definitionsContext()
}

type Relationship_definitionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Relationship_definitionsContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserIDENT)
}

func (s *Relationship_definitionsContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserIDENT, i)
}

func (s *Relationship_definitionsContext) RELATION() antlr.TerminalNode {
	return s.GetToken(GraphGenParserRELATION, 0)
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

func (p *GraphGenParser) Relationship_definitions() (localctx IRelationship_definitionsContext) {
	localctx = NewRelationship_definitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphGenParserRULE_relationship_definitions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserVIRGULA {
		{
			p.SetState(36)
			p.Match(GraphGenParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(GraphGenParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(GraphGenParserRELATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(GraphGenParserIDENT)
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

	// Getter signatures
	DRAW() antlr.TerminalNode
	OPENPAR() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	CLOSEPAR() antlr.TerminalNode
	AllVIRGULA() []antlr.TerminalNode
	VIRGULA(i int) antlr.TerminalNode
	AllRELATION() []antlr.TerminalNode
	RELATION(i int) antlr.TerminalNode

	// IsDraw_commandContext differentiates from other interfaces.
	IsDraw_commandContext()
}

type Draw_commandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Draw_commandContext) DRAW() antlr.TerminalNode {
	return s.GetToken(GraphGenParserDRAW, 0)
}

func (s *Draw_commandContext) OPENPAR() antlr.TerminalNode {
	return s.GetToken(GraphGenParserOPENPAR, 0)
}

func (s *Draw_commandContext) IDENT() antlr.TerminalNode {
	return s.GetToken(GraphGenParserIDENT, 0)
}

func (s *Draw_commandContext) CLOSEPAR() antlr.TerminalNode {
	return s.GetToken(GraphGenParserCLOSEPAR, 0)
}

func (s *Draw_commandContext) AllVIRGULA() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserVIRGULA)
}

func (s *Draw_commandContext) VIRGULA(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserVIRGULA, i)
}

func (s *Draw_commandContext) AllRELATION() []antlr.TerminalNode {
	return s.GetTokens(GraphGenParserRELATION)
}

func (s *Draw_commandContext) RELATION(i int) antlr.TerminalNode {
	return s.GetToken(GraphGenParserRELATION, i)
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

func (p *GraphGenParser) Draw_command() (localctx IDraw_commandContext) {
	localctx = NewDraw_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphGenParserRULE_draw_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(GraphGenParserDRAW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(GraphGenParserOPENPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(GraphGenParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphGenParserVIRGULA {
		{
			p.SetState(49)
			p.Match(GraphGenParserVIRGULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(GraphGenParserRELATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(GraphGenParserCLOSEPAR)
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
