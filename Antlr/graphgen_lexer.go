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
		"", "", "'person'", "'group'", "'relationship'", "'subgroup of'", "'draw'",
		"'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW",
		"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT",
		"ERROR",
	}
	staticData.RuleNames = []string{
		"COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW",
		"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT",
		"ERROR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 109, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 5, 0, 30, 8, 0, 10,
		0, 12, 0, 33, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 90,
		8, 9, 10, 9, 12, 9, 93, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		5, 11, 101, 8, 11, 10, 11, 12, 11, 104, 9, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 2, 31, 102, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1, 0, 5, 1, 0, 10, 10, 2, 0,
		65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 2, 0, 10, 10, 125, 125, 111, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27,
		1, 0, 0, 0, 3, 38, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9,
		64, 1, 0, 0, 0, 11, 76, 1, 0, 0, 0, 13, 81, 1, 0, 0, 0, 15, 83, 1, 0, 0,
		0, 17, 85, 1, 0, 0, 0, 19, 87, 1, 0, 0, 0, 21, 94, 1, 0, 0, 0, 23, 98,
		1, 0, 0, 0, 25, 107, 1, 0, 0, 0, 27, 31, 5, 123, 0, 0, 28, 30, 8, 0, 0,
		0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 31, 29,
		1, 0, 0, 0, 32, 34, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5, 125, 0,
		0, 35, 36, 1, 0, 0, 0, 36, 37, 6, 0, 0, 0, 37, 2, 1, 0, 0, 0, 38, 39, 5,
		112, 0, 0, 39, 40, 5, 101, 0, 0, 40, 41, 5, 114, 0, 0, 41, 42, 5, 115,
		0, 0, 42, 43, 5, 111, 0, 0, 43, 44, 5, 110, 0, 0, 44, 4, 1, 0, 0, 0, 45,
		46, 5, 103, 0, 0, 46, 47, 5, 114, 0, 0, 47, 48, 5, 111, 0, 0, 48, 49, 5,
		117, 0, 0, 49, 50, 5, 112, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52, 5, 114, 0,
		0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 108, 0, 0, 54, 55, 5, 97, 0, 0, 55,
		56, 5, 116, 0, 0, 56, 57, 5, 105, 0, 0, 57, 58, 5, 111, 0, 0, 58, 59, 5,
		110, 0, 0, 59, 60, 5, 115, 0, 0, 60, 61, 5, 104, 0, 0, 61, 62, 5, 105,
		0, 0, 62, 63, 5, 112, 0, 0, 63, 8, 1, 0, 0, 0, 64, 65, 5, 115, 0, 0, 65,
		66, 5, 117, 0, 0, 66, 67, 5, 98, 0, 0, 67, 68, 5, 103, 0, 0, 68, 69, 5,
		114, 0, 0, 69, 70, 5, 111, 0, 0, 70, 71, 5, 117, 0, 0, 71, 72, 5, 112,
		0, 0, 72, 73, 5, 32, 0, 0, 73, 74, 5, 111, 0, 0, 74, 75, 5, 102, 0, 0,
		75, 10, 1, 0, 0, 0, 76, 77, 5, 100, 0, 0, 77, 78, 5, 114, 0, 0, 78, 79,
		5, 97, 0, 0, 79, 80, 5, 119, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5, 40, 0,
		0, 82, 14, 1, 0, 0, 0, 83, 84, 5, 41, 0, 0, 84, 16, 1, 0, 0, 0, 85, 86,
		5, 44, 0, 0, 86, 18, 1, 0, 0, 0, 87, 91, 7, 1, 0, 0, 88, 90, 7, 2, 0, 0,
		89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1,
		0, 0, 0, 92, 20, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 7, 3, 0, 0, 95,
		96, 1, 0, 0, 0, 96, 97, 6, 10, 0, 0, 97, 22, 1, 0, 0, 0, 98, 102, 5, 123,
		0, 0, 99, 101, 8, 4, 0, 0, 100, 99, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102,
		103, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 105, 106, 5, 10, 0, 0, 106, 24, 1, 0, 0, 0, 107, 108, 9, 0,
		0, 0, 108, 26, 1, 0, 0, 0, 4, 0, 31, 91, 102, 1, 6, 0, 0,
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
	GraphGenLexerSUBGROUP_OF        = 5
	GraphGenLexerDRAW               = 6
	GraphGenLexerOPENPAR            = 7
	GraphGenLexerCLOSEPAR           = 8
	GraphGenLexerVIRGULA            = 9
	GraphGenLexerIDENT              = 10
	GraphGenLexerWS                 = 11
	GraphGenLexerERROR_OPEN_COMMENT = 12
	GraphGenLexerERROR              = 13
)
