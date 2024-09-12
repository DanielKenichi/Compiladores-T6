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
			SemanticError(identifier.GetSymbol(), fmt.Sprintf("variable %v already declared", identifier.GetText())),
		)
		return result
	}

	v.Scopes.CurrentScope().AddSymbol(identifier.GetText(), varType)

	return result
}

func (v *GraphGenVisitor) CheckSubGroupDefinitions(ctx parser.ISubgroups_definitionsContext) []string {
	result := make([]string, 0)

	for _, ident := range ctx.AllIDENT() {

		varDeclarationResult := v.CheckVarDeclaration(ident)

		if len(varDeclarationResult) > 0 {
			return varDeclarationResult
		}

		identType := v.Scopes.CurrentScope().GetType(ident.GetText())

		if identType != symboltable.GROUP {
			result = append(result,
				SemanticError(ident.GetSymbol(), fmt.Sprintf("variable %v is not a group", ident.GetText())))
			return result
		}
	}

	return result
}

func (v *GraphGenVisitor) CheckVarDeclaration(ident antlr.TerminalNode) []string {
	result := make([]string, 0)

	if !v.Scopes.CurrentScope().Exists(ident.GetText()) {
		result = append(result,
			SemanticError(ident.GetSymbol(), fmt.Sprintf("variable %v not declared", ident.GetText())))
		return result
	}

	return result
}
