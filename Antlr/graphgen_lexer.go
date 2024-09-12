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

type GraphGenLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GraphGenLexerLexerStaticData struct {
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

func graphgenlexerLexerInit() {
	staticData := &GraphGenLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'person'", "'group'", "'relationship'", "'filter by'", "'draw'",
		"','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "FILTER_BY", "DRAW",
		"VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT", "ERROR",
	}
	staticData.RuleNames = []string{
		"COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "FILTER_BY", "DRAW", "VIRGULA",
		"IDENT", "WS", "ERROR_OPEN_COMMENT", "ERROR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 99, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 5,
		7, 80, 8, 7, 10, 7, 12, 7, 83, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		5, 9, 91, 8, 9, 10, 9, 12, 9, 94, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 2, 27,
		92, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 1, 0, 5, 1, 0, 10, 10, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 125,
		125, 101, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 1,
		23, 1, 0, 0, 0, 3, 34, 1, 0, 0, 0, 5, 41, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0,
		9, 60, 1, 0, 0, 0, 11, 70, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 77, 1, 0,
		0, 0, 17, 84, 1, 0, 0, 0, 19, 88, 1, 0, 0, 0, 21, 97, 1, 0, 0, 0, 23, 27,
		5, 123, 0, 0, 24, 26, 8, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0,
		0, 27, 28, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 27,
		1, 0, 0, 0, 30, 31, 5, 125, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 6, 0, 0,
		0, 33, 2, 1, 0, 0, 0, 34, 35, 5, 112, 0, 0, 35, 36, 5, 101, 0, 0, 36, 37,
		5, 114, 0, 0, 37, 38, 5, 115, 0, 0, 38, 39, 5, 111, 0, 0, 39, 40, 5, 110,
		0, 0, 40, 4, 1, 0, 0, 0, 41, 42, 5, 103, 0, 0, 42, 43, 5, 114, 0, 0, 43,
		44, 5, 111, 0, 0, 44, 45, 5, 117, 0, 0, 45, 46, 5, 112, 0, 0, 46, 6, 1,
		0, 0, 0, 47, 48, 5, 114, 0, 0, 48, 49, 5, 101, 0, 0, 49, 50, 5, 108, 0,
		0, 50, 51, 5, 97, 0, 0, 51, 52, 5, 116, 0, 0, 52, 53, 5, 105, 0, 0, 53,
		54, 5, 111, 0, 0, 54, 55, 5, 110, 0, 0, 55, 56, 5, 115, 0, 0, 56, 57, 5,
		104, 0, 0, 57, 58, 5, 105, 0, 0, 58, 59, 5, 112, 0, 0, 59, 8, 1, 0, 0,
		0, 60, 61, 5, 102, 0, 0, 61, 62, 5, 105, 0, 0, 62, 63, 5, 108, 0, 0, 63,
		64, 5, 116, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66, 5, 114, 0, 0, 66, 67, 5,
		32, 0, 0, 67, 68, 5, 98, 0, 0, 68, 69, 5, 121, 0, 0, 69, 10, 1, 0, 0, 0,
		70, 71, 5, 100, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73, 5, 97, 0, 0, 73, 74,
		5, 119, 0, 0, 74, 12, 1, 0, 0, 0, 75, 76, 5, 44, 0, 0, 76, 14, 1, 0, 0,
		0, 77, 81, 7, 1, 0, 0, 78, 80, 7, 2, 0, 0, 79, 78, 1, 0, 0, 0, 80, 83,
		1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 16, 1, 0, 0, 0,
		83, 81, 1, 0, 0, 0, 84, 85, 7, 3, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 6,
		8, 0, 0, 87, 18, 1, 0, 0, 0, 88, 92, 5, 123, 0, 0, 89, 91, 8, 4, 0, 0,
		90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 92, 90, 1,
		0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 10, 0, 0, 96,
		20, 1, 0, 0, 0, 97, 98, 9, 0, 0, 0, 98, 22, 1, 0, 0, 0, 4, 0, 27, 81, 92,
		1, 6, 0, 0,
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

// GraphGenLexerInit initializes any static state used to implement GraphGenLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGraphGenLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GraphGenLexerInit() {
	staticData := &GraphGenLexerLexerStaticData
	staticData.once.Do(graphgenlexerLexerInit)
}

// NewGraphGenLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGraphGenLexer(input antlr.CharStream) *GraphGenLexer {
	GraphGenLexerInit()
	l := new(GraphGenLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GraphGenLexerLexerStaticData
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

// GraphGenLexer tokens.
const (
	GraphGenLexerCOMMENT            = 1
	GraphGenLexerPERSON             = 2
	GraphGenLexerGROUP              = 3
	GraphGenLexerRELATIONSHIP       = 4
	GraphGenLexerFILTER_BY          = 5
	GraphGenLexerDRAW               = 6
	GraphGenLexerVIRGULA            = 7
	GraphGenLexerIDENT              = 8
	GraphGenLexerWS                 = 9
	GraphGenLexerERROR_OPEN_COMMENT = 10
	GraphGenLexerERROR              = 11
)
