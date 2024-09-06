// Code generated from Antlr/GraphGen.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GraphGen

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GraphGenParser.
type GraphGenVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GraphGenParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GraphGenParser#declarations.
	VisitDeclarations(ctx *DeclarationsContext) interface{}

	// Visit a parse tree produced by GraphGenParser#subgroups_definitions.
	VisitSubgroups_definitions(ctx *Subgroups_definitionsContext) interface{}

	// Visit a parse tree produced by GraphGenParser#relationship_definitions.
	VisitRelationship_definitions(ctx *Relationship_definitionsContext) interface{}

	// Visit a parse tree produced by GraphGenParser#draw_command.
	VisitDraw_command(ctx *Draw_commandContext) interface{}
}
