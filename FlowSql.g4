grammar FlowSql;
sql
    :
    SELECT fields FROM ID (where)? ';'
    ;

where
    : WHERE ID CP expr (AND expr)*
    ;

expr
    : expr op=('*'|'/') expr  # MulDiv
    | expr op=('-'|'+') expr  # AddSub
    | NUMBER                  # number
    | ID                      # id
    | '(' expr ')'            # parens
    ;

field
    : ID
    | alias
    |
    ;

fields
    : field (',' field)*
    ;


alias
    : expr AS ID
    ;


SELECT
    : [Ss] [Ee] [Ll] [Ee] [Cc] [Tt]
    ;

FROM
    : [Ff] [Rr] [Oo] [Mm]
    ;

WHERE
    : [Ww] [Hh] [Ee] [Rr] [Ee]
    ;

AND
    : [Aa] [Nn] [Dd]
    ;

AS
    : [Aa] [Ss]
    ;


fragment LETTER
    : [a-zA-Z_]
    ;

ID
    : [a-zA-Z] [0-9a-zA-Z]*
    ;

fragment INT
    : '0'
    | [1-9] [0-9]*
    ;

NUMBER
    : '-'? INT ('.'INT+)?
    ;

CP
    : '='
    | '>'
    | '<'
    | '>='
    | '<='
    | '!='
    ;

WS
    : [ \t\n\r]+ -> skip
    ;
