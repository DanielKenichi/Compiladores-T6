// Generated from /home/lucky/UFSCAR/Compiladores/Compiladores-T6/Antlr/GraphGen.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GraphGenParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMMENT=1, PERSON=2, GROUP=3, RELATIONSHIP=4, SUBGROUP_OF=5, DRAW=6, OPENPAR=7, 
		CLOSEPAR=8, VIRGULA=9, IDENT=10, WS=11, ERROR_OPEN_COMMENT=12, ERROR=13;
	public static final int
		RULE_program = 0, RULE_declarations = 1, RULE_var_type = 2, RULE_subgroups_definitions = 3, 
		RULE_relationship_definitions = 4, RULE_draw_command = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "declarations", "var_type", "subgroups_definitions", "relationship_definitions", 
			"draw_command"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'person'", "'group'", "'relationship'", "'subgroup of'", 
			"'draw'", "'('", "')'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW", 
			"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT", 
			"ERROR"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "GraphGen.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GraphGenParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<DeclarationsContext> declarations() {
			return getRuleContexts(DeclarationsContext.class);
		}
		public DeclarationsContext declarations(int i) {
			return getRuleContext(DeclarationsContext.class,i);
		}
		public List<Subgroups_definitionsContext> subgroups_definitions() {
			return getRuleContexts(Subgroups_definitionsContext.class);
		}
		public Subgroups_definitionsContext subgroups_definitions(int i) {
			return getRuleContext(Subgroups_definitionsContext.class,i);
		}
		public List<Relationship_definitionsContext> relationship_definitions() {
			return getRuleContexts(Relationship_definitionsContext.class);
		}
		public Relationship_definitionsContext relationship_definitions(int i) {
			return getRuleContext(Relationship_definitionsContext.class,i);
		}
		public List<Draw_commandContext> draw_command() {
			return getRuleContexts(Draw_commandContext.class);
		}
		public Draw_commandContext draw_command(int i) {
			return getRuleContext(Draw_commandContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(15);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 28L) != 0)) {
				{
				{
				setState(12);
				declarations();
				}
				}
				setState(17);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(21);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(18);
					subgroups_definitions();
					}
					} 
				}
				setState(23);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			setState(27);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(24);
				relationship_definitions();
				}
				}
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(33);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DRAW) {
				{
				{
				setState(30);
				draw_command();
				}
				}
				setState(35);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationsContext extends ParserRuleContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public List<TerminalNode> IDENT() { return getTokens(GraphGenParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GraphGenParser.IDENT, i);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(GraphGenParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(GraphGenParser.VIRGULA, i);
		}
		public DeclarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarations; }
	}

	public final DeclarationsContext declarations() throws RecognitionException {
		DeclarationsContext _localctx = new DeclarationsContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			var_type();
			setState(37);
			match(IDENT);
			setState(42);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(38);
				match(VIRGULA);
				setState(39);
				match(IDENT);
				}
				}
				setState(44);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Var_typeContext extends ParserRuleContext {
		public TerminalNode PERSON() { return getToken(GraphGenParser.PERSON, 0); }
		public TerminalNode GROUP() { return getToken(GraphGenParser.GROUP, 0); }
		public TerminalNode RELATIONSHIP() { return getToken(GraphGenParser.RELATIONSHIP, 0); }
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_var_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 28L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Subgroups_definitionsContext extends ParserRuleContext {
		public List<TerminalNode> IDENT() { return getTokens(GraphGenParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GraphGenParser.IDENT, i);
		}
		public TerminalNode SUBGROUP_OF() { return getToken(GraphGenParser.SUBGROUP_OF, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(GraphGenParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(GraphGenParser.VIRGULA, i);
		}
		public Subgroups_definitionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subgroups_definitions; }
	}

	public final Subgroups_definitionsContext subgroups_definitions() throws RecognitionException {
		Subgroups_definitionsContext _localctx = new Subgroups_definitionsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_subgroups_definitions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(IDENT);
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(48);
				match(VIRGULA);
				setState(49);
				match(IDENT);
				}
				}
				setState(54);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(55);
			match(SUBGROUP_OF);
			setState(56);
			match(IDENT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Relationship_definitionsContext extends ParserRuleContext {
		public Token relation;
		public Token related;
		public List<TerminalNode> IDENT() { return getTokens(GraphGenParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GraphGenParser.IDENT, i);
		}
		public List<TerminalNode> VIRGULA() { return getTokens(GraphGenParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(GraphGenParser.VIRGULA, i);
		}
		public Relationship_definitionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationship_definitions; }
	}

	public final Relationship_definitionsContext relationship_definitions() throws RecognitionException {
		Relationship_definitionsContext _localctx = new Relationship_definitionsContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_relationship_definitions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			match(IDENT);
			setState(63);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(59);
				match(VIRGULA);
				setState(60);
				match(IDENT);
				}
				}
				setState(65);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(66);
			((Relationship_definitionsContext)_localctx).relation = match(IDENT);
			setState(67);
			((Relationship_definitionsContext)_localctx).related = match(IDENT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Draw_commandContext extends ParserRuleContext {
		public Token relation;
		public TerminalNode DRAW() { return getToken(GraphGenParser.DRAW, 0); }
		public TerminalNode OPENPAR() { return getToken(GraphGenParser.OPENPAR, 0); }
		public List<TerminalNode> IDENT() { return getTokens(GraphGenParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GraphGenParser.IDENT, i);
		}
		public TerminalNode CLOSEPAR() { return getToken(GraphGenParser.CLOSEPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(GraphGenParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(GraphGenParser.VIRGULA, i);
		}
		public Draw_commandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_draw_command; }
	}

	public final Draw_commandContext draw_command() throws RecognitionException {
		Draw_commandContext _localctx = new Draw_commandContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_draw_command);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(DRAW);
			setState(70);
			match(OPENPAR);
			setState(71);
			match(IDENT);
			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(72);
				match(VIRGULA);
				setState(73);
				((Draw_commandContext)_localctx).relation = match(IDENT);
				}
				}
				setState(78);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(79);
			match(CLOSEPAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\rR\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0001\u0000\u0005\u0000\u000e\b\u0000\n\u0000\f\u0000"+
		"\u0011\t\u0000\u0001\u0000\u0005\u0000\u0014\b\u0000\n\u0000\f\u0000\u0017"+
		"\t\u0000\u0001\u0000\u0005\u0000\u001a\b\u0000\n\u0000\f\u0000\u001d\t"+
		"\u0000\u0001\u0000\u0005\u0000 \b\u0000\n\u0000\f\u0000#\t\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0005\u0001)\b\u0001\n\u0001"+
		"\f\u0001,\t\u0001\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0005\u00033\b\u0003\n\u0003\f\u00036\t\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004>\b"+
		"\u0004\n\u0004\f\u0004A\t\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005K\b"+
		"\u0005\n\u0005\f\u0005N\t\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0000"+
		"\u0000\u0006\u0000\u0002\u0004\u0006\b\n\u0000\u0001\u0001\u0000\u0002"+
		"\u0004S\u0000\u000f\u0001\u0000\u0000\u0000\u0002$\u0001\u0000\u0000\u0000"+
		"\u0004-\u0001\u0000\u0000\u0000\u0006/\u0001\u0000\u0000\u0000\b:\u0001"+
		"\u0000\u0000\u0000\nE\u0001\u0000\u0000\u0000\f\u000e\u0003\u0002\u0001"+
		"\u0000\r\f\u0001\u0000\u0000\u0000\u000e\u0011\u0001\u0000\u0000\u0000"+
		"\u000f\r\u0001\u0000\u0000\u0000\u000f\u0010\u0001\u0000\u0000\u0000\u0010"+
		"\u0015\u0001\u0000\u0000\u0000\u0011\u000f\u0001\u0000\u0000\u0000\u0012"+
		"\u0014\u0003\u0006\u0003\u0000\u0013\u0012\u0001\u0000\u0000\u0000\u0014"+
		"\u0017\u0001\u0000\u0000\u0000\u0015\u0013\u0001\u0000\u0000\u0000\u0015"+
		"\u0016\u0001\u0000\u0000\u0000\u0016\u001b\u0001\u0000\u0000\u0000\u0017"+
		"\u0015\u0001\u0000\u0000\u0000\u0018\u001a\u0003\b\u0004\u0000\u0019\u0018"+
		"\u0001\u0000\u0000\u0000\u001a\u001d\u0001\u0000\u0000\u0000\u001b\u0019"+
		"\u0001\u0000\u0000\u0000\u001b\u001c\u0001\u0000\u0000\u0000\u001c!\u0001"+
		"\u0000\u0000\u0000\u001d\u001b\u0001\u0000\u0000\u0000\u001e \u0003\n"+
		"\u0005\u0000\u001f\u001e\u0001\u0000\u0000\u0000 #\u0001\u0000\u0000\u0000"+
		"!\u001f\u0001\u0000\u0000\u0000!\"\u0001\u0000\u0000\u0000\"\u0001\u0001"+
		"\u0000\u0000\u0000#!\u0001\u0000\u0000\u0000$%\u0003\u0004\u0002\u0000"+
		"%*\u0005\n\u0000\u0000&\'\u0005\t\u0000\u0000\')\u0005\n\u0000\u0000("+
		"&\u0001\u0000\u0000\u0000),\u0001\u0000\u0000\u0000*(\u0001\u0000\u0000"+
		"\u0000*+\u0001\u0000\u0000\u0000+\u0003\u0001\u0000\u0000\u0000,*\u0001"+
		"\u0000\u0000\u0000-.\u0007\u0000\u0000\u0000.\u0005\u0001\u0000\u0000"+
		"\u0000/4\u0005\n\u0000\u000001\u0005\t\u0000\u000013\u0005\n\u0000\u0000"+
		"20\u0001\u0000\u0000\u000036\u0001\u0000\u0000\u000042\u0001\u0000\u0000"+
		"\u000045\u0001\u0000\u0000\u000057\u0001\u0000\u0000\u000064\u0001\u0000"+
		"\u0000\u000078\u0005\u0005\u0000\u000089\u0005\n\u0000\u00009\u0007\u0001"+
		"\u0000\u0000\u0000:?\u0005\n\u0000\u0000;<\u0005\t\u0000\u0000<>\u0005"+
		"\n\u0000\u0000=;\u0001\u0000\u0000\u0000>A\u0001\u0000\u0000\u0000?=\u0001"+
		"\u0000\u0000\u0000?@\u0001\u0000\u0000\u0000@B\u0001\u0000\u0000\u0000"+
		"A?\u0001\u0000\u0000\u0000BC\u0005\n\u0000\u0000CD\u0005\n\u0000\u0000"+
		"D\t\u0001\u0000\u0000\u0000EF\u0005\u0006\u0000\u0000FG\u0005\u0007\u0000"+
		"\u0000GL\u0005\n\u0000\u0000HI\u0005\t\u0000\u0000IK\u0005\n\u0000\u0000"+
		"JH\u0001\u0000\u0000\u0000KN\u0001\u0000\u0000\u0000LJ\u0001\u0000\u0000"+
		"\u0000LM\u0001\u0000\u0000\u0000MO\u0001\u0000\u0000\u0000NL\u0001\u0000"+
		"\u0000\u0000OP\u0005\b\u0000\u0000P\u000b\u0001\u0000\u0000\u0000\b\u000f"+
		"\u0015\u001b!*4?L";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}