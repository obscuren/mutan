package mutan

// Mutatis mutandis - Changing [only] those things which need to be changed.
import ()

const grammer = `prgm <- statement+
statement <- _?^ compound_statement
compound_statement <- _?^ expression _?^
expression <- assign_expression / arithmetic_expression / simple_expression
assign_expression <- simple_expression _?^ '=' _?^ expression
arithmetic_expression <- simple_expression _?^ operator _?^ simple_expression
definition <- array
simple_expression <- call / store / create_array / array / number / name
array <- name open_bracket _^? number _^? close_bracket
call <- 'call' open arg+ close
arg <- simple_expression ~','?^ _?^
number <- ~'-?\d+\.?\d*'
name <- ~'[a-zA-Z/\-\*\+=><]+'
store <- 'store' open_bracket number close_bracket
create_array <- 'array' open _^? number _^? close
operator <- ~'[\-+/*]'
open_bracket <- '['
close_bracket <- ']'
open <- '('
close <- ')'
_ <- ~'\s+'`

/*
const grammer = `prgm <- expr+
expr <- _?^ assignment / arithmetic / atom
assignment <- assignees _?^ '=' _?^ expr
arithmetic <- atom _?^ arith _?^ atom
arith <- ~'[\-+/*]'
assignees <- name / store
atom <- number / name / store
store <- open number close
name <- ~'[a-zA-Z/\-\*\+=><]+'
number <- ~'-?\d+\.?\d*'
open <- '['
close <- ']'
_ <- ~'\s+'`
*/
