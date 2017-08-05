# Time

Lexemes:

    - digit     [0-9]+
    - letter    [A-Za-z]+
    - slash     /
    - colon     :
    - dash      -
    - space     ' '
    - dot       . 
    - other

Post-processed things:

    - weekday
    - monthName
    - timezone
    - am/pm

Use lexeme "labels" for these? Mark lexemes as making up clock info, timezone,
weekday, month, day, year?


The lexer must be context-sensitive: 2006-02-01 is not the same as -0700.

Recognizable patterns:

    - A colon infects all numbers near to it with a "clock" attribute
    - a +(digit) or (not digit)-(digit) signals timezone "drift"
    - weekday lexemes are worthless

Proposal for parsing dates:

    - Lex the input string
    - Search for any digits connected by colons, and delete them
    - Search for digits preceeded by +, or by (not-digit)-, and delete them
    - Any remaining numbers should be year, month and day
    - Use hints in remaining lexemes to figure out which it is
