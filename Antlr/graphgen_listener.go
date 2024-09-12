// Code generated from Antlr/GraphGen.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GraphGen

import "github.com/antlr4-go/antlr/v4"

// GraphGenListener is a complete listener for a parse tree produced by GraphGenParser.
type GraphGenListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterRelationship_definitions is called when entering the relationship_definitions production.
	EnterRelationship_definitions(c *Relationship_definitionsContext)

	// EnterDraw_command is called when entering the draw_command production.
	EnterDraw_command(c *Draw_commandContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitRelationship_definitions is called when exiting the relationship_definitions production.
	ExitRelationship_definitions(c *Relationship_definitionsContext)

	// ExitDraw_command is called when exiting the draw_command production.
	ExitDraw_command(c *Draw_commandContext)
}
