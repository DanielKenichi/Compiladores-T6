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
		CLOSEPAR=8, IDENT=9, WS=10, ERRO=11;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW", 
			"OPENPAR", "CLOSEPAR", "IDENT", "WS", "ERRO"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'person'", "'group'", "'relationship'", "'subgroup of'", 
			"'draw'", "'('", "')'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMMENT", "PERSON", "GROUP", "RELATIONSHIP", "SUBGROUP_OF", "DRAW", 
			"OPENPAR", "CLOSEPAR", "IDENT", "WS", "ERRO"
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
		"\u0004\u0000\u000b^\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0001\u0000"+
		"\u0001\u0000\u0005\u0000\u001a\b\u0000\n\u0000\f\u0000\u001d\t\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0005\bT\b\b\n\b\f\bW\t"+
		"\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u001b\u0000\u000b"+
		"\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r"+
		"\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0001\u0000\u0004\u0001\u0000"+
		"\n\n\u0002\u0000AZaz\u0004\u000009AZ__az\u0003\u0000\t\n\r\r  _\u0000"+
		"\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000"+
		"\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000"+
		"\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r"+
		"\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011"+
		"\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015"+
		"\u0001\u0000\u0000\u0000\u0001\u0017\u0001\u0000\u0000\u0000\u0003\"\u0001"+
		"\u0000\u0000\u0000\u0005)\u0001\u0000\u0000\u0000\u0007/\u0001\u0000\u0000"+
		"\u0000\t<\u0001\u0000\u0000\u0000\u000bH\u0001\u0000\u0000\u0000\rM\u0001"+
		"\u0000\u0000\u0000\u000fO\u0001\u0000\u0000\u0000\u0011Q\u0001\u0000\u0000"+
		"\u0000\u0013X\u0001\u0000\u0000\u0000\u0015\\\u0001\u0000\u0000\u0000"+
		"\u0017\u001b\u0005{\u0000\u0000\u0018\u001a\b\u0000\u0000\u0000\u0019"+
		"\u0018\u0001\u0000\u0000\u0000\u001a\u001d\u0001\u0000\u0000\u0000\u001b"+
		"\u001c\u0001\u0000\u0000\u0000\u001b\u0019\u0001\u0000\u0000\u0000\u001c"+
		"\u001e\u0001\u0000\u0000\u0000\u001d\u001b\u0001\u0000\u0000\u0000\u001e"+
		"\u001f\u0005}\u0000\u0000\u001f \u0001\u0000\u0000\u0000 !\u0006\u0000"+
		"\u0000\u0000!\u0002\u0001\u0000\u0000\u0000\"#\u0005p\u0000\u0000#$\u0005"+
		"e\u0000\u0000$%\u0005r\u0000\u0000%&\u0005s\u0000\u0000&\'\u0005o\u0000"+
		"\u0000\'(\u0005n\u0000\u0000(\u0004\u0001\u0000\u0000\u0000)*\u0005g\u0000"+
		"\u0000*+\u0005r\u0000\u0000+,\u0005o\u0000\u0000,-\u0005u\u0000\u0000"+
		"-.\u0005p\u0000\u0000.\u0006\u0001\u0000\u0000\u0000/0\u0005r\u0000\u0000"+
		"01\u0005e\u0000\u000012\u0005l\u0000\u000023\u0005a\u0000\u000034\u0005"+
		"t\u0000\u000045\u0005i\u0000\u000056\u0005o\u0000\u000067\u0005n\u0000"+
		"\u000078\u0005s\u0000\u000089\u0005h\u0000\u00009:\u0005i\u0000\u0000"+
		":;\u0005p\u0000\u0000;\b\u0001\u0000\u0000\u0000<=\u0005s\u0000\u0000"+
		"=>\u0005u\u0000\u0000>?\u0005b\u0000\u0000?@\u0005g\u0000\u0000@A\u0005"+
		"r\u0000\u0000AB\u0005o\u0000\u0000BC\u0005u\u0000\u0000CD\u0005p\u0000"+
		"\u0000DE\u0005 \u0000\u0000EF\u0005o\u0000\u0000FG\u0005f\u0000\u0000"+
		"G\n\u0001\u0000\u0000\u0000HI\u0005d\u0000\u0000IJ\u0005r\u0000\u0000"+
		"JK\u0005a\u0000\u0000KL\u0005w\u0000\u0000L\f\u0001\u0000\u0000\u0000"+
		"MN\u0005(\u0000\u0000N\u000e\u0001\u0000\u0000\u0000OP\u0005)\u0000\u0000"+
		"P\u0010\u0001\u0000\u0000\u0000QU\u0007\u0001\u0000\u0000RT\u0007\u0002"+
		"\u0000\u0000SR\u0001\u0000\u0000\u0000TW\u0001\u0000\u0000\u0000US\u0001"+
		"\u0000\u0000\u0000UV\u0001\u0000\u0000\u0000V\u0012\u0001\u0000\u0000"+
		"\u0000WU\u0001\u0000\u0000\u0000XY\u0007\u0003\u0000\u0000YZ\u0001\u0000"+
		"\u0000\u0000Z[\u0006\t\u0000\u0000[\u0014\u0001\u0000\u0000\u0000\\]\t"+
		"\u0000\u0000\u0000]\u0016\u0001\u0000\u0000\u0000\u0003\u0000\u001bU\u0001"+
		"\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}