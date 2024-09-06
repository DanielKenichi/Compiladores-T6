grammar GraphGen;

/*
 Lexico
*/

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


/*
    Sintatico
*/

TYPE: PERSON | GROUP | RELATIONSHIP;
RELATION: IDENT;

program
    : (declarations)* (subgroups_definitions)* (relationship_definitions)* (draw_command)*;

declarations
    : TYPE IDENT (VIRGULA IDENT)*;

subgroups_definitions
    : IDENT (VIRGULA IDENT)* SUBGROUP_OF IDENT;

relationship_definitions
    : IDENT (VIRGULA IDENT)* RELATION IDENT;

draw_command
    : DRAW OPENPAR IDENT (VIRGULA RELATION)* CLOSEPAR;
