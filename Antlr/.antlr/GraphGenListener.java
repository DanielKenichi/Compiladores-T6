// Generated from /home/rodrigo/UFSCAR/Compiladores/Compiladores-T6/Antlr/GraphGen.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GraphGenParser}.
 */
public interface GraphGenListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GraphGenParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(GraphGenParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link GraphGenParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(GraphGenParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link GraphGenParser#declarations}.
	 * @param ctx the parse tree
	 */
	void enterDeclarations(GraphGenParser.DeclarationsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GraphGenParser#declarations}.
	 * @param ctx the parse tree
	 */
	void exitDeclarations(GraphGenParser.DeclarationsContext ctx);
	/**
	 * Enter a parse tree produced by {@link GraphGenParser#subgroups_definitions}.
	 * @param ctx the parse tree
	 */
	void enterSubgroups_definitions(GraphGenParser.Subgroups_definitionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GraphGenParser#subgroups_definitions}.
	 * @param ctx the parse tree
	 */
	void exitSubgroups_definitions(GraphGenParser.Subgroups_definitionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link GraphGenParser#relationship_definitions}.
	 * @param ctx the parse tree
	 */
	void enterRelationship_definitions(GraphGenParser.Relationship_definitionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GraphGenParser#relationship_definitions}.
	 * @param ctx the parse tree
	 */
	void exitRelationship_definitions(GraphGenParser.Relationship_definitionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link GraphGenParser#draw_command}.
	 * @param ctx the parse tree
	 */
	void enterDraw_command(GraphGenParser.Draw_commandContext ctx);
	/**
	 * Exit a parse tree produced by {@link GraphGenParser#draw_command}.
	 * @param ctx the parse tree
	 */
	void exitDraw_command(GraphGenParser.Draw_commandContext ctx);
}