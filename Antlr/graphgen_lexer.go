// Code generated from Antlr/GraphGen.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GraphGen struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GraphGenLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func graphgenLexerInit() {
	staticData := &GraphGenLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'person'", "'group'", "'relationship'", "'subgroup of'", "'draw'",
		"'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW",
		"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERRO",
	}
	staticData.RuleNames = []string{
		"COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW",
		"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERRO",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 98, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 88, 8, 9, 10, 9,
		12, 9, 91, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 29, 0, 12,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 1, 0, 4, 1, 0, 10, 10, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 99, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0,
		0, 0, 3, 36, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 49, 1, 0, 0, 0, 9, 62, 1,
		0, 0, 0, 11, 74, 1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 81, 1, 0, 0, 0, 17,
		83, 1, 0, 0, 0, 19, 85, 1, 0, 0, 0, 21, 92, 1, 0, 0, 0, 23, 96, 1, 0, 0,
		0, 25, 29, 5, 123, 0, 0, 26, 28, 8, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 31,
		1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0,
		31, 29, 1, 0, 0, 0, 32, 33, 5, 125, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 6,
		0, 0, 0, 35, 2, 1, 0, 0, 0, 36, 37, 5, 112, 0, 0, 37, 38, 5, 101, 0, 0,
		38, 39, 5, 114, 0, 0, 39, 40, 5, 115, 0, 0, 40, 41, 5, 111, 0, 0, 41, 42,
		5, 110, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 103, 0, 0, 44, 45, 5, 114,
		0, 0, 45, 46, 5, 111, 0, 0, 46, 47, 5, 117, 0, 0, 47, 48, 5, 112, 0, 0,
		48, 6, 1, 0, 0, 0, 49, 50, 5, 114, 0, 0, 50, 51, 5, 101, 0, 0, 51, 52,
		5, 108, 0, 0, 52, 53, 5, 97, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5, 105,
		0, 0, 55, 56, 5, 111, 0, 0, 56, 57, 5, 110, 0, 0, 57, 58, 5, 115, 0, 0,
		58, 59, 5, 104, 0, 0, 59, 60, 5, 105, 0, 0, 60, 61, 5, 112, 0, 0, 61, 8,
		1, 0, 0, 0, 62, 63, 5, 115, 0, 0, 63, 64, 5, 117, 0, 0, 64, 65, 5, 98,
		0, 0, 65, 66, 5, 103, 0, 0, 66, 67, 5, 114, 0, 0, 67, 68, 5, 111, 0, 0,
		68, 69, 5, 117, 0, 0, 69, 70, 5, 112, 0, 0, 70, 71, 5, 32, 0, 0, 71, 72,
		5, 111, 0, 0, 72, 73, 5, 102, 0, 0, 73, 10, 1, 0, 0, 0, 74, 75, 5, 100,
		0, 0, 75, 76, 5, 114, 0, 0, 76, 77, 5, 97, 0, 0, 77, 78, 5, 119, 0, 0,
		78, 12, 1, 0, 0, 0, 79, 80, 5, 40, 0, 0, 80, 14, 1, 0, 0, 0, 81, 82, 5,
		41, 0, 0, 82, 16, 1, 0, 0, 0, 83, 84, 5, 44, 0, 0, 84, 18, 1, 0, 0, 0,
		85, 89, 7, 1, 0, 0, 86, 88, 7, 2, 0, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1,
		0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 20, 1, 0, 0, 0, 91,
		89, 1, 0, 0, 0, 92, 93, 7, 3, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 6, 10,
		0, 0, 95, 22, 1, 0, 0, 0, 96, 97, 9, 0, 0, 0, 97, 24, 1, 0, 0, 0, 3, 0,
		29, 89, 1, 6, 0, 0,
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

// GraphGenInit initializes any static state used to implement GraphGen. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGraphGen(). You can call this function if you wish to initialize the static state ahead
// of time.
func GraphGenInit() {
	staticData := &GraphGenLexerStaticData
	staticData.once.Do(graphgenLexerInit)
}

// NewGraphGen produces a new lexer instance for the optional input antlr.CharStream.
func NewGraphGen(input antlr.CharStream) *GraphGen {
	GraphGenInit()
	l := new(GraphGen)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GraphGenLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "GraphGen.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GraphGen tokens.
const (
	GraphGenCOMMENT      = 1
	GraphGenPERSON       = 2
	GraphGenGROUP        = 3
	GraphGenRELATIONSHIP = 4
	GraphGenSUBGROUP_OF  = 5
	GraphGenDRAW         = 6
	GraphGenOPENPAR      = 7
	GraphGenCLOSEPAR     = 8
	GraphGenVIRGULA      = 9
	GraphGenIDENT        = 10
	GraphGenWS           = 11
	GraphGenERRO         = 12
)
