lexer grammar GraphGen;

COMMENT: '{' ~('\n')*? '}' -> skip;

//type definition words
PERSON: 'person'; 
GROUP: 'group';
RELATIONSHIP: 'relationship';

//subgroup definition
SUBGROUP_OF: 'subgroup of';

//key word to build a graph
DRAW: 'draw';

OPENPAR: '(';
CLOSEPAR: ')';
VIRGULA: ',';

//identificadores
IDENT: ('a'..'z'|'A'..'Z')('a'..'z'|'A'..'Z'|'0'..'9'|'_')*;

//whitespaces
WS: ( ' ' | '\t' | '\r' | '\n' ) -> skip;


ERROR_OPEN_COMMENT : '{' (~('\n'|'}'))*? '\n';

//some lexer error
ERROR: .;

// TYPE: PERSON | GROUP | RELATIONSHIP;




