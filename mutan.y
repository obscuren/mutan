%{
package main

import (
	"fmt"
)

var Tree *SyntaxTree

%}

%union {
	num int
	str string
	tnode *SyntaxTree
}

%token ASSIGN
%token <str> ID
%token <num> NUMBER
%type <tnode> program statement_list statement expression assign_expression simple_expression get_variable

%%

program
	: statement_list { Tree = $1 }
	;

statement_list
	: statement_list statement { $$ = NewNode(StatementListTy, $1, $2) }
	| /* Empty */ { $$ = NewNode(EmptyTy) }
	;

statement
	: expression { $$ = $1 }
	;

expression
	: assign_expression { $$ = $1 }
	;

assign_expression
	: ID ASSIGN assign_expression
	  {
	      node := NewNode(SetLocalTy)
	      node.Constant = $1
	      $$ = NewNode(AssignmentTy, $3, node)
	  }
	| simple_expression { $$ = $1 }
	;

simple_expression
	: get_variable { $$ = $1 }
	;

get_variable
	: ID { $$ = NewNode(ConstantTy); $$.Constant = $1 }
	;

%%

func main() {
	yyParse(lexer("example", "a = 0"))
	fmt.Println(Tree)
}

