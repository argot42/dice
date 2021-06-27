grammar Dice;

// Tokens
DIC: 'd';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUM: [0-9]+;
WS: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

expression
    : expression DIC expression             # Dic
    | expression op=('*'|'/') expression    # MulDiv
    | expression op=('+'|'-') expression    # AddSub
    | NUM                                   # Number 
    ;
