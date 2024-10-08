package main

import (
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type customErrorListener struct {
	output     *os.File
	syntax_err bool
}

func NewCustomErrorListener(output *os.File) *customErrorListener {
	return &customErrorListener{
		output: output,
	}
}

func (l *customErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *customErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (l *customErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

func (l *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, charPositionInLine int, msg string, e antlr.RecognitionException) {
	// Reportar apenas o primeiro erro, ignorar os outros
	if !l.syntax_err {
		t := offendingSymbol.(antlr.Token)

		// Modificação da string "EOF" para conformidade com os casos de testes
		tokenText := t.GetText()
		if tokenText == "<EOF>" {
			tokenText = "EOF"
		}

		// Impressão do erro
		l.output.WriteString("Line " + strconv.Itoa(line) + ": syntax error next to " + tokenText + "\n")
		l.syntax_err = true
	}
}
