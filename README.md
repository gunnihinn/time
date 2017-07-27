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

The lexer must be context-sensitive: 2006-02-01 is not the same as -0700.
