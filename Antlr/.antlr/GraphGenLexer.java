// Generated from /home/lucky/UFSCAR/Compiladores/Compiladores-T6/Antlr/GraphGen.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class GraphGenLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMMENT=1, PERSON=2, GROUP=3, RELATIONSHIP=4, SUBGROUP_OF=5, DRAW=6, OPENPAR=7, 
		CLOSEPAR=8, VIRGULA=9, IDENT=10, WS=11, ERROR_OPEN_COMMENT=12, ERROR=13;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW", 
			"OPENPAR", "CLOSEPAR", "VIRGULA", "IDENT", "WS", "ERROR_OPEN_COMMENT", 
			"ERROR"
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


	public GraphGenLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "GraphGen.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\rm\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0001\u0000\u0001\u0000\u0005\u0000\u001e"+
		"\b\u0000\n\u0000\f\u0000!\t\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0005\tZ\b\t\n\t\f\t]\t\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0005\u000be\b\u000b\n\u000b"+
		"\f\u000bh\t\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0002\u001f"+
		"f\u0000\r\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b"+
		"\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u0001"+
		"\u0000\u0005\u0001\u0000\n\n\u0002\u0000AZaz\u0004\u000009AZ__az\u0003"+
		"\u0000\t\n\r\r  \u0002\u0000\n\n}}o\u0000\u0001\u0001\u0000\u0000\u0000"+
		"\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000"+
		"\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000"+
		"\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f"+
		"\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013"+
		"\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017"+
		"\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0001\u001b"+
		"\u0001\u0000\u0000\u0000\u0003&\u0001\u0000\u0000\u0000\u0005-\u0001\u0000"+
		"\u0000\u0000\u00073\u0001\u0000\u0000\u0000\t@\u0001\u0000\u0000\u0000"+
		"\u000bL\u0001\u0000\u0000\u0000\rQ\u0001\u0000\u0000\u0000\u000fS\u0001"+
		"\u0000\u0000\u0000\u0011U\u0001\u0000\u0000\u0000\u0013W\u0001\u0000\u0000"+
		"\u0000\u0015^\u0001\u0000\u0000\u0000\u0017b\u0001\u0000\u0000\u0000\u0019"+
		"k\u0001\u0000\u0000\u0000\u001b\u001f\u0005{\u0000\u0000\u001c\u001e\b"+
		"\u0000\u0000\u0000\u001d\u001c\u0001\u0000\u0000\u0000\u001e!\u0001\u0000"+
		"\u0000\u0000\u001f \u0001\u0000\u0000\u0000\u001f\u001d\u0001\u0000\u0000"+
		"\u0000 \"\u0001\u0000\u0000\u0000!\u001f\u0001\u0000\u0000\u0000\"#\u0005"+
		"}\u0000\u0000#$\u0001\u0000\u0000\u0000$%\u0006\u0000\u0000\u0000%\u0002"+
		"\u0001\u0000\u0000\u0000&\'\u0005p\u0000\u0000\'(\u0005e\u0000\u0000("+
		")\u0005r\u0000\u0000)*\u0005s\u0000\u0000*+\u0005o\u0000\u0000+,\u0005"+
		"n\u0000\u0000,\u0004\u0001\u0000\u0000\u0000-.\u0005g\u0000\u0000./\u0005"+
		"r\u0000\u0000/0\u0005o\u0000\u000001\u0005u\u0000\u000012\u0005p\u0000"+
		"\u00002\u0006\u0001\u0000\u0000\u000034\u0005r\u0000\u000045\u0005e\u0000"+
		"\u000056\u0005l\u0000\u000067\u0005a\u0000\u000078\u0005t\u0000\u0000"+
		"89\u0005i\u0000\u00009:\u0005o\u0000\u0000:;\u0005n\u0000\u0000;<\u0005"+
		"s\u0000\u0000<=\u0005h\u0000\u0000=>\u0005i\u0000\u0000>?\u0005p\u0000"+
		"\u0000?\b\u0001\u0000\u0000\u0000@A\u0005s\u0000\u0000AB\u0005u\u0000"+
		"\u0000BC\u0005b\u0000\u0000CD\u0005g\u0000\u0000DE\u0005r\u0000\u0000"+
		"EF\u0005o\u0000\u0000FG\u0005u\u0000\u0000GH\u0005p\u0000\u0000HI\u0005"+
		" \u0000\u0000IJ\u0005o\u0000\u0000JK\u0005f\u0000\u0000K\n\u0001\u0000"+
		"\u0000\u0000LM\u0005d\u0000\u0000MN\u0005r\u0000\u0000NO\u0005a\u0000"+
		"\u0000OP\u0005w\u0000\u0000P\f\u0001\u0000\u0000\u0000QR\u0005(\u0000"+
		"\u0000R\u000e\u0001\u0000\u0000\u0000ST\u0005)\u0000\u0000T\u0010\u0001"+
		"\u0000\u0000\u0000UV\u0005,\u0000\u0000V\u0012\u0001\u0000\u0000\u0000"+
		"W[\u0007\u0001\u0000\u0000XZ\u0007\u0002\u0000\u0000YX\u0001\u0000\u0000"+
		"\u0000Z]\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000[\\\u0001\u0000"+
		"\u0000\u0000\\\u0014\u0001\u0000\u0000\u0000][\u0001\u0000\u0000\u0000"+
		"^_\u0007\u0003\u0000\u0000_`\u0001\u0000\u0000\u0000`a\u0006\n\u0000\u0000"+
		"a\u0016\u0001\u0000\u0000\u0000bf\u0005{\u0000\u0000ce\b\u0004\u0000\u0000"+
		"dc\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000\u0000fg\u0001\u0000\u0000"+
		"\u0000fd\u0001\u0000\u0000\u0000gi\u0001\u0000\u0000\u0000hf\u0001\u0000"+
		"\u0000\u0000ij\u0005\n\u0000\u0000j\u0018\u0001\u0000\u0000\u0000kl\t"+
		"\u0000\u0000\u0000l\u001a\u0001\u0000\u0000\u0000\u0004\u0000\u001f[f"+
		"\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}