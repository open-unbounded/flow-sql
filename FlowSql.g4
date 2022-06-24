grammar FlowSql;
sql
    :
    SELECT field FROM ID (where)? (orderby)? ';'
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
    | alias (',' alias)*
    ;

alias
    : expr AS ID
    ;

orderby
    : ORDER_BY ID
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

ORDER_BY
    : [Oo] [Rr] [Dd] [Ee] [Rr] (' ')+  [Bb] [Yy]
    ;
AND
    : [Aa] [Nn] [Dd]
    ;
AS
    : [Aa] [Ss]
    ;

fragment DIGIT
    : [0-9]
    ;

fragment LETTER
    : [a-zA-Z_]
    ;

ID
    : LETTER+ DIGIT* LETTER*;

NUMBER
    : [0-9]+('.'[0-9]+)?;

CP
    : '='
    | '>'
    | '<'
    | '>='
    | '<='
    ;

WS
    : [ \t\n\r]+ -> skip
    ;
