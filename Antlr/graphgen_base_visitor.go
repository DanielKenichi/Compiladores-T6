// Code generated from Antlr/GraphGen.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GraphGen

import "github.com/antlr4-go/antlr/v4"

type BaseGraphGenVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGraphGenVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphGenVisitor) VisitDeclarations(ctx *DeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphGenVisitor) VisitVar_type(ctx *Var_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphGenVisitor) VisitSubgroups_definitions(ctx *Subgroups_definitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphGenVisitor) VisitRelationship_definitions(ctx *Relationship_definitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraphGenVisitor) VisitDraw_command(ctx *Draw_commandContext) interface{} {
	return v.VisitChildren(ctx)
}
