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
		CLOSEPAR=8, VIRGULA=9, IDENT=10, WS=11, ERROR_OPEN_COMMENT=12, ERROR=13, 
		TYPE=14, RELATION=15;
	public static final int
		RULE_program = 0, RULE_declarations = 1, RULE_subgroups_definitions = 2, 
		RULE_relationship_definitions = 3, RULE_draw_command = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "declarations", "subgroups_definitions", "relationship_definitions", 
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
			"ERROR", "TYPE", "RELATION"
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
			setState(13);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(10);
				declarations();
				}
				}
				setState(15);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(19);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(16);
					subgroups_definitions();
					}
					} 
				}
				setState(21);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			setState(25);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(22);
				relationship_definitions();
				}
				}
				setState(27);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(31);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DRAW) {
				{
				{
				setState(28);
				draw_command();
				}
				}
				setState(33);
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
		public TerminalNode TYPE() { return getToken(GraphGenParser.TYPE, 0); }
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
			setState(34);
			match(TYPE);
			setState(35);
			match(IDENT);
			setState(40);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(36);
				match(VIRGULA);
				setState(37);
				match(IDENT);
				}
				}
				setState(42);
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
		enterRule(_localctx, 4, RULE_subgroups_definitions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			match(IDENT);
			setState(48);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(44);
				match(VIRGULA);
				setState(45);
				match(IDENT);
				}
				}
				setState(50);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(51);
			match(SUBGROUP_OF);
			setState(52);
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
		public List<TerminalNode> IDENT() { return getTokens(GraphGenParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(GraphGenParser.IDENT, i);
		}
		public TerminalNode RELATION() { return getToken(GraphGenParser.RELATION, 0); }
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
		enterRule(_localctx, 6, RULE_relationship_definitions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(IDENT);
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(55);
				match(VIRGULA);
				setState(56);
				match(IDENT);
				}
				}
				setState(61);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(62);
			match(RELATION);
			setState(63);
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
	public static class Draw_commandContext extends ParserRuleContext {
		public TerminalNode DRAW() { return getToken(GraphGenParser.DRAW, 0); }
		public TerminalNode OPENPAR() { return getToken(GraphGenParser.OPENPAR, 0); }
		public TerminalNode IDENT() { return getToken(GraphGenParser.IDENT, 0); }
		public TerminalNode CLOSEPAR() { return getToken(GraphGenParser.CLOSEPAR, 0); }
		public List<TerminalNode> VIRGULA() { return getTokens(GraphGenParser.VIRGULA); }
		public TerminalNode VIRGULA(int i) {
			return getToken(GraphGenParser.VIRGULA, i);
		}
		public List<TerminalNode> RELATION() { return getTokens(GraphGenParser.RELATION); }
		public TerminalNode RELATION(int i) {
			return getToken(GraphGenParser.RELATION, i);
		}
		public Draw_commandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_draw_command; }
	}

	public final Draw_commandContext draw_command() throws RecognitionException {
		Draw_commandContext _localctx = new Draw_commandContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_draw_command);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			match(DRAW);
			setState(66);
			match(OPENPAR);
			setState(67);
			match(IDENT);
			setState(72);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VIRGULA) {
				{
				{
				setState(68);
				match(VIRGULA);
				setState(69);
				match(RELATION);
				}
				}
				setState(74);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(75);
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
		"\u0004\u0001\u000fN\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0001"+
		"\u0000\u0005\u0000\f\b\u0000\n\u0000\f\u0000\u000f\t\u0000\u0001\u0000"+
		"\u0005\u0000\u0012\b\u0000\n\u0000\f\u0000\u0015\t\u0000\u0001\u0000\u0005"+
		"\u0000\u0018\b\u0000\n\u0000\f\u0000\u001b\t\u0000\u0001\u0000\u0005\u0000"+
		"\u001e\b\u0000\n\u0000\f\u0000!\t\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0005\u0001\'\b\u0001\n\u0001\f\u0001*\t\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0005\u0002/\b\u0002\n\u0002\f\u00022\t"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0005\u0003:\b\u0003\n\u0003\f\u0003=\t\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0005\u0004G\b\u0004\n\u0004\f\u0004J\t\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0000\u0000\u0005\u0000\u0002\u0004\u0006\b\u0000\u0000"+
		"P\u0000\r\u0001\u0000\u0000\u0000\u0002\"\u0001\u0000\u0000\u0000\u0004"+
		"+\u0001\u0000\u0000\u0000\u00066\u0001\u0000\u0000\u0000\bA\u0001\u0000"+
		"\u0000\u0000\n\f\u0003\u0002\u0001\u0000\u000b\n\u0001\u0000\u0000\u0000"+
		"\f\u000f\u0001\u0000\u0000\u0000\r\u000b\u0001\u0000\u0000\u0000\r\u000e"+
		"\u0001\u0000\u0000\u0000\u000e\u0013\u0001\u0000\u0000\u0000\u000f\r\u0001"+
		"\u0000\u0000\u0000\u0010\u0012\u0003\u0004\u0002\u0000\u0011\u0010\u0001"+
		"\u0000\u0000\u0000\u0012\u0015\u0001\u0000\u0000\u0000\u0013\u0011\u0001"+
		"\u0000\u0000\u0000\u0013\u0014\u0001\u0000\u0000\u0000\u0014\u0019\u0001"+
		"\u0000\u0000\u0000\u0015\u0013\u0001\u0000\u0000\u0000\u0016\u0018\u0003"+
		"\u0006\u0003\u0000\u0017\u0016\u0001\u0000\u0000\u0000\u0018\u001b\u0001"+
		"\u0000\u0000\u0000\u0019\u0017\u0001\u0000\u0000\u0000\u0019\u001a\u0001"+
		"\u0000\u0000\u0000\u001a\u001f\u0001\u0000\u0000\u0000\u001b\u0019\u0001"+
		"\u0000\u0000\u0000\u001c\u001e\u0003\b\u0004\u0000\u001d\u001c\u0001\u0000"+
		"\u0000\u0000\u001e!\u0001\u0000\u0000\u0000\u001f\u001d\u0001\u0000\u0000"+
		"\u0000\u001f \u0001\u0000\u0000\u0000 \u0001\u0001\u0000\u0000\u0000!"+
		"\u001f\u0001\u0000\u0000\u0000\"#\u0005\u000e\u0000\u0000#(\u0005\n\u0000"+
		"\u0000$%\u0005\t\u0000\u0000%\'\u0005\n\u0000\u0000&$\u0001\u0000\u0000"+
		"\u0000\'*\u0001\u0000\u0000\u0000(&\u0001\u0000\u0000\u0000()\u0001\u0000"+
		"\u0000\u0000)\u0003\u0001\u0000\u0000\u0000*(\u0001\u0000\u0000\u0000"+
		"+0\u0005\n\u0000\u0000,-\u0005\t\u0000\u0000-/\u0005\n\u0000\u0000.,\u0001"+
		"\u0000\u0000\u0000/2\u0001\u0000\u0000\u00000.\u0001\u0000\u0000\u0000"+
		"01\u0001\u0000\u0000\u000013\u0001\u0000\u0000\u000020\u0001\u0000\u0000"+
		"\u000034\u0005\u0005\u0000\u000045\u0005\n\u0000\u00005\u0005\u0001\u0000"+
		"\u0000\u00006;\u0005\n\u0000\u000078\u0005\t\u0000\u00008:\u0005\n\u0000"+
		"\u000097\u0001\u0000\u0000\u0000:=\u0001\u0000\u0000\u0000;9\u0001\u0000"+
		"\u0000\u0000;<\u0001\u0000\u0000\u0000<>\u0001\u0000\u0000\u0000=;\u0001"+
		"\u0000\u0000\u0000>?\u0005\u000f\u0000\u0000?@\u0005\n\u0000\u0000@\u0007"+
		"\u0001\u0000\u0000\u0000AB\u0005\u0006\u0000\u0000BC\u0005\u0007\u0000"+
		"\u0000CH\u0005\n\u0000\u0000DE\u0005\t\u0000\u0000EG\u0005\u000f\u0000"+
		"\u0000FD\u0001\u0000\u0000\u0000GJ\u0001\u0000\u0000\u0000HF\u0001\u0000"+
		"\u0000\u0000HI\u0001\u0000\u0000\u0000IK\u0001\u0000\u0000\u0000JH\u0001"+
		"\u0000\u0000\u0000KL\u0005\b\u0000\u0000L\t\u0001\u0000\u0000\u0000\b"+
		"\r\u0013\u0019\u001f(0;H";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}