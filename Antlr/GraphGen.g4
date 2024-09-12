grammar GraphGen;

/* 
---------------
    Léxico
--------------- 
*/

COMMENT: '{' ~('\n')*? '}' -> skip;

//type definition words
PERSON: 'person'; 
GROUP: 'group';
RELATIONSHIP: 'relationship';

//filter of a person's relationships
FILTER_BY: 'filter by';

//key word to build a graph
DRAW: 'draw';

VIRGULA: ',';

//identificadores
IDENT: ('a'..'z'|'A'..'Z')('a'..'z'|'A'..'Z'|'0'..'9'|'_')*;

//whitespaces
WS: ( ' ' | '\t' | '\r' | '\n' ) -> skip;

ERROR_OPEN_COMMENT : '{' (~('\n'|'}'))*? '\n';

//some lexer error
ERROR: .;


/*
-----------------
    Sintatico
-----------------
*/

program
    : (declarations)* (relationship_definitions)* (draw_command)*;

declarations
    : var_type IDENT (VIRGULA IDENT)*;

var_type: PERSON | GROUP | RELATIONSHIP;

relationship_definitions
    : IDENT (VIRGULA IDENT)* relation=IDENT related=IDENT;

draw_command
    : DRAW IDENT (FILTER_BY filter=IDENT)?;
