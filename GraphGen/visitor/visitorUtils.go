package visitor

import (
	"fmt"
	"log"

	parser "github.com/DanielKenichi/Compiladores-T6/Antlr"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/relations"
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

		if len(result) > 0 {
			return result
		}

		if varType == symboltable.GROUP {
			v.Relations.AddNewGroup(varName.GetText())
		} else if varType == symboltable.PERSON {
			v.Relations.AddNewPerson(varName.GetText())
		}
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

func (v *GraphGenVisitor) CheckRelationShipsDefinitions(relantionship parser.IRelationship_definitionsContext) []string {
	result := make([]string, 0)

	alreadyDeclaredInRelationship := map[string]bool{}

	relation := relantionship.GetRelation()

	for _, ident := range relantionship.AllIDENT() {
		varDeclarationResult := v.CheckVarDeclaration(ident)

		if len(varDeclarationResult) > 0 {
			return varDeclarationResult
		}

		_, ok := alreadyDeclaredInRelationship[ident.GetText()]

		if ok {
			result = append(result,
				SemanticError(ident.GetSymbol(), fmt.Sprintf("variable %v is declared more than once in the relationship", ident.GetText())))

			return result
		}

		alreadyDeclaredInRelationship[ident.GetText()] = true

		if ident.GetText() == relation.GetText() {
			continue
		}

		varType := v.Scopes.CurrentScope().GetType(ident.GetText())

		if varType == symboltable.RELATIONSHIP {
			result = append(result,
				SemanticError(ident.GetSymbol(), fmt.Sprintf("variable %v must be a group or a person", ident.GetText())))

			return result
		}
	}

	return result
}

func (v *GraphGenVisitor) BuildRelations(relantionship parser.IRelationship_definitionsContext) {
	related := relantionship.GetRelated()
	relation := relantionship.GetRelation()

	for _, ident := range relantionship.AllIDENT() {
		if ident.GetText() == related.GetText() || ident.GetText() == relation.GetText() {
			continue
		}
		persons := make([]*relations.Person, 0)
		relatedPersons := make([]*relations.Person, 0)

		relatedType := v.Scopes.CurrentScope().GetType(related.GetText())
		varType := v.Scopes.CurrentScope().GetType(ident.GetText())

		if relatedType == symboltable.GROUP {
			relatedPersons = v.Relations.GetGroupMembers(related.GetText())
		} else {
			relatedPerson := v.Relations.GetPerson(ident.GetText())

			relatedPersons = append(relatedPersons, relatedPerson)
		}

		if varType == symboltable.GROUP {
			persons = v.Relations.GetGroupMembers(ident.GetText())
		} else {
			person := v.Relations.GetPerson(ident.GetText())

			persons = append(persons, person)
		}

		for _, person := range persons {
			for _, relatedPerson := range relatedPersons {
				person.AddPersonToRelationship(relatedPerson, relation.GetText())
				relatedPerson.AddPersonToRelationship(person, relation.GetText())
			}

			if relatedType == symboltable.GROUP {
				v.Relations.AddPersonToGroup(person, related.GetText())
			}
		}

	}
}

func (v *GraphGenVisitor) CheckDrawCommandsCall(drawCommand parser.IDraw_commandContext) []string {
	result := make([]string, 0)

	originNode := drawCommand.IDENT(0)

	varDeclarationResult := v.CheckVarDeclaration(originNode)

	if len(varDeclarationResult) > 0 {
		return varDeclarationResult
	}

	originNodeType := v.Scopes.CurrentScope().GetType(originNode.GetText())

	if originNodeType == symboltable.RELATIONSHIP {
		result = append(result,
			SemanticError(originNode.GetSymbol(), fmt.Sprintf("relationship %v is not a node to draw", originNode.GetText())))

		return result
	}

	if originNodeType == symboltable.GROUP && drawCommand.FILTER_BY() != nil {
		result = append(result,
			SemanticError(originNode.GetSymbol(), "only persons relationships can be filtered"))

		return result
	}

	if originNodeType == symboltable.PERSON && drawCommand.FILTER_BY() != nil {
		filter := drawCommand.IDENT(1)

		filterType := v.Scopes.CurrentScope().GetType(filter.GetText())

		if filterType != symboltable.RELATIONSHIP {
			result = append(result,
				SemanticError(filter.GetSymbol(), fmt.Sprintf("variable %v must be a filter", filter.GetText())))
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
