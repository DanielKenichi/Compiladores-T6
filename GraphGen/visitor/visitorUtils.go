package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/Compiladores-T6/Antlr"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

func (v *GraphGenVisitor) AddVarsToSymbolTable(ctx parser.IDeclarationsContext) []string {
	var declarationsResult = make([]string, 0)

	var varType symboltable.Type = symboltable.INVALIDO

	if ctx.TYPE() != nil {
		switch ctx.TYPE().GetText() {
		case "person":
			varType = symboltable.PERSON
		case "group":
			varType = symboltable.GROUP
		case "relationship":
			varType = symboltable.RELATIONSHIP
		}
	}

	for _, varName := range ctx.AllIDENT() {
		v.AddIdentifierToSymbolTable(varName, varType)
	}

	return declarationsResult
}

func (v *GraphGenVisitor) AddIdentifierToSymbolTable(identifier antlr.TerminalNode, varType symboltable.Type) []string {
	result := make([]string, 0)

	if v.Scopes.CurrentScope().Exists(identifier.GetText()) {
		result = append(result,
			SemanticError(identifier.GetSymbol(), fmt.Sprintf("identificador %v ja declarado anteriormente", identifier.GetText())),
		)

		return result
	}

	v.Scopes.CurrentScope().AddSymbol(identifier.GetText(), varType)

	return result
}
