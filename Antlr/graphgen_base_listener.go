// Code generated from Antlr/GraphGen.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GraphGen

import "github.com/antlr4-go/antlr/v4"

// BaseGraphGenListener is a complete listener for a parse tree produced by GraphGenParser.
type BaseGraphGenListener struct{}

var _ GraphGenListener = &BaseGraphGenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphGenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphGenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphGenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphGenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGraphGenListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGraphGenListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseGraphGenListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseGraphGenListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterSubgroups_definitions is called when production subgroups_definitions is entered.
func (s *BaseGraphGenListener) EnterSubgroups_definitions(ctx *Subgroups_definitionsContext) {}

// ExitSubgroups_definitions is called when production subgroups_definitions is exited.
func (s *BaseGraphGenListener) ExitSubgroups_definitions(ctx *Subgroups_definitionsContext) {}

// EnterRelationship_definitions is called when production relationship_definitions is entered.
func (s *BaseGraphGenListener) EnterRelationship_definitions(ctx *Relationship_definitionsContext) {}

// ExitRelationship_definitions is called when production relationship_definitions is exited.
func (s *BaseGraphGenListener) ExitRelationship_definitions(ctx *Relationship_definitionsContext) {}

// EnterDraw_command is called when production draw_command is entered.
func (s *BaseGraphGenListener) EnterDraw_command(ctx *Draw_commandContext) {}

// ExitDraw_command is called when production draw_command is exited.
func (s *BaseGraphGenListener) ExitDraw_command(ctx *Draw_commandContext) {}
