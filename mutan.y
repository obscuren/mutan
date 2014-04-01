%{
package mutan

import (
	_"fmt"
)

var Tree *SyntaxTree

%}

%union {
	num int
	str string
	tnode *SyntaxTree
}

%token ASSIGN EQUAL IF LEFT_BRACES RIGHT_BRACES STORE LEFT_BRACKET RIGHT_BRACKET ASM LEFT_PAR RIGHT_PAR STOP
%token ADDR ORIGIN CALLER CALLVAL CALLDATALOAD CALLDATASIZE GASPRICE DOT THIS
%token <str> ID NUMBER INLINE_ASM OP
%type <tnode> program statement_list statement expression assign_expression simple_expression get_variable
%type <tnode> if_statement op_expression buildins closure_funcs

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
	| if_statement { $$ = $1 }
	| ASM LEFT_PAR INLINE_ASM RIGHT_PAR { $$ = NewNode(InlineAsmTy); $$.Constant = $3 }
	;


buildins
	: STOP LEFT_PAR RIGHT_PAR { $$ = NewNode(StopTy) }
	| THIS DOT closure_funcs { $$ = $3 }
	;

closure_funcs
	: ORIGIN LEFT_PAR RIGHT_PAR { $$ = NewNode(OriginTy) }
	| CALLER LEFT_PAR RIGHT_PAR { $$ = NewNode(CallerTy) }
	| CALLVAL LEFT_PAR RIGHT_PAR { $$ = NewNode(CallValTy) }
	| CALLDATALOAD LEFT_PAR RIGHT_PAR { $$ = NewNode(CallDataLoadTy) }
	| CALLDATASIZE LEFT_PAR RIGHT_PAR { $$ = NewNode(CallDataSizeTy) }
	| GASPRICE LEFT_PAR RIGHT_PAR { $$ = NewNode(GasPriceTy) }
	;

if_statement
	: IF expression LEFT_BRACES statement_list RIGHT_BRACES { $$ = NewNode(IfThenTy, $2, $4) }
	;

expression
	: op_expression { $$ = $1 }
	| assign_expression { $$ = $1 }
	;

op_expression
	: expression OP expression { $$ = NewNode(OpTy, $1, $3); $$.Constant = $2 }


assign_expression
	: ID ASSIGN assign_expression
	  {
	      node := NewNode(SetLocalTy)
	      node.Constant = $1
	      $$ = NewNode(AssignmentTy, $3, node)
	  }
	| STORE LEFT_BRACKET expression RIGHT_BRACKET ASSIGN assign_expression
	  {
	      node := NewNode(SetStoreTy, $3)
	      $$ = NewNode(AssignmentTy, $6, node)
	  }
	| simple_expression { $$ = $1 }
	;

simple_expression
	: get_variable { $$ = $1 }
	;

get_variable
	: ID { $$ = NewNode(IdentifierTy); $$.Constant = $1 }
	| NUMBER { $$ = NewNode(ConstantTy); $$.Constant = $1 }
	| STORE LEFT_BRACKET expression RIGHT_BRACKET { $$ = NewNode(StoreTy, $3) }
	| buildins { $$ = $1 }
	;

%%

