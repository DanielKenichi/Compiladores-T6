package visitor

import (
	"fmt"

	parser "github.com/DanielKenichi/Compiladores-T6/Antlr"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/relations"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/scope"
	"github.com/antlr4-go/antlr/v4"
)

type GraphGenVisitor struct {
	parser.BaseGraphGenVisitor
	Scopes    *scope.Scope
	Relations *relations.Relations
}

func New() *GraphGenVisitor {
	return &GraphGenVisitor{
		Scopes:    scope.New(),
		Relations: relations.NewRelations(),
	}
}

func SemanticError(token antlr.Token, message string) string {
	return fmt.Sprintf("Line %v: %s\n", token.GetLine(), message)
}

func (v *GraphGenVisitor) VisitProgram(ctx parser.IProgramContext) []string {
	var programResult = make([]string, 0)
	v.Scopes.NewScope()

	result := v.VisitDeclarations(ctx.AllDeclarations())
	programResult = append(programResult, result...)

	result = v.VisitRelationshipDefinitions(ctx.AllRelationship_definitions())
	programResult = append(programResult, result...)

	result = v.VisitDrawCommands(ctx.AllDraw_command())
	programResult = append(programResult, result...)

	return programResult
}

func (v *GraphGenVisitor) VisitDeclarations(ctxs []parser.IDeclarationsContext) []string {
	var declarationsResult = make([]string, 0)

	for _, ctx := range ctxs {

		result := v.AddVarsToSymbolTable(ctx)

		declarationsResult = append(declarationsResult, result...)
	}

	return declarationsResult
}

func (v *GraphGenVisitor) VisitRelationshipDefinitions(ctxs []parser.IRelationship_definitionsContext) []string {
	var relationshipDefinitionsResult = make([]string, 0)

	for _, ctx := range ctxs {
		result := v.CheckRelationShipsDefinitions(ctx)

		if len(result) > 0 {
			return result
		}

		v.BuildRelations(ctx)

		relationshipDefinitionsResult = append(relationshipDefinitionsResult, result...)
	}

	return relationshipDefinitionsResult
}

func (v *GraphGenVisitor) VisitDrawCommands(ctxs []parser.IDraw_commandContext) []string {
	var drawCommandResult = make([]string, 0)

	for _, ctx := range ctxs {
		result := v.CheckDrawCommandsCall(ctx)

		drawCommandResult = append(drawCommandResult, result...)
	}

	return drawCommandResult
}
