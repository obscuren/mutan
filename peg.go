package mutan

// Mutatis mutandis - Changing [only] those things which need to be changed.
import ()

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
