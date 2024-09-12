package visitor

import (
	"fmt"
	"log"

	parser "github.com/DanielKenichi/Compiladores-T6/Antlr"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/symboltable"
	"github.com/antlr4-go/antlr/v4"
)

func (v *GraphGenVisitor) AddVarsToSymbolTable(ctx parser.IDeclarationsContext) []string {
	var declarationsResult = make([]string, 0)

	var varType symboltable.Type = symboltable.INVALIDO

	if ctx.Var_type() != nil {
		if ctx.Var_type().PERSON() != nil {
			varType = symboltable.PERSON
		} else if ctx.Var_type().GROUP() != nil {
			varType = symboltable.GROUP
		} else if ctx.Var_type().RELATIONSHIP() != nil {
			varType = symboltable.RELATIONSHIP
		}
	}

	for _, varName := range ctx.AllIDENT() {
		result := v.AddIdentifierToSymbolTable(varName, varType)

		declarationsResult = append(declarationsResult, result...)
	}

	return declarationsResult
}

func (v *GraphGenVisitor) AddIdentifierToSymbolTable(identifier antlr.TerminalNode, varType symboltable.Type) []string {
	result := make([]string, 0)

	log.Printf("%v", v.Scopes.CurrentScope())
	log.Printf("Adding ident %v", identifier.GetText())
	if v.Scopes.CurrentScope().Exists(identifier.GetText()) {
		result = append(result,
			SemanticError(identifier.GetSymbol(), fmt.Sprintf("var %v already declared", identifier.GetText())),
		)
		return result
	}

	v.Scopes.CurrentScope().AddSymbol(identifier.GetText(), varType)

	return result
}
