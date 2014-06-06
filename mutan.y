%{
package mutan

import (
_	"fmt"
)

var Tree *SyntaxTree

// Helper function to turn a tree in to a regular list.
// Especially handy when parsing argument lists
func makeSlice(tree *SyntaxTree) (ret []*SyntaxTree) {
	if tree != nil && tree.Type != EmptyTy {
		ret = append(ret, tree)
		for _, i := range tree.Children {
			ret = append(ret, makeSlice(i)...)
		}
		tree.Children = nil
	} 

	return
}

func makeArgs(tree *SyntaxTree, reverse bool) (ret []*SyntaxTree) {
	l := makeSlice(tree)
	// Quick reverse
	if reverse {
		for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
			l[i], l[j] = l[j], l[i]
		}
	}

	for _, s := range l {
		if s.Type != ArgTy {
			ret = append(ret, s)
		}
	}

	return
}

%}

%union {
	num int
	str string
	tnode *SyntaxTree
    	check bool
}

%token ASSIGN EQUAL IF ELSE FOR LEFT_BRACES RIGHT_BRACES STORE LEFT_BRACKET RIGHT_BRACKET ASM LEFT_PAR RIGHT_PAR STOP
%token ADDR ORIGIN CALLER CALLVAL CALLDATALOAD CALLDATASIZE GASPRICE DOT THIS ARRAY CALL COMMA SIZEOF QUOTE
%token END_STMT EXIT CREATE TRANSACT NIL BALANCE VAR_ASSIGN LAMBDA COLON ADDRESS RETURN PUSH POP
%token DIFFICULTY PREVHASH TIMESTAMP GASPRICE BLOCKNUM COINBASE GAS FOR VAR FUNC FUNC_CALL
%token <str> ID NUMBER INLINE_ASM OP DOP STR BOOLEAN CODE oper AND MUL
%type <tnode> program statement_list statement expression assign_expression simple_expression get_variable
%type <tnode> if_statement op_expression buildins closure_funcs new_var new_array arguments sep get_id string
%type <tnode> for_statement optional_else_statement ptr sub_expression opt_arg_def_list opt_arg_call_list
%type <check> optional_type

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
	| LAMBDA LEFT_BRACKET CODE RIGHT_BRACKET { $$ = NewNode(LambdaTy); $$.Constant = $3 }
	| if_statement { $$ = $1 }
	| for_statement { $$ = $1 }
	| FUNC ID LEFT_PAR opt_arg_def_list RIGHT_PAR optional_type LEFT_BRACES statement_list RIGHT_BRACES
		{
			$$ = NewNode(FuncDefTy, $8);
			$$.Constant = $2
			$$.HasRet = $6
			$$.ArgList = makeArgs($4, true)
		}
	| ASM LEFT_PAR INLINE_ASM RIGHT_PAR { $$ = NewNode(InlineAsmTy); $$.Constant = $3 }
	| END_STMT { $$ = NewNode(EmptyTy); }
	;

opt_arg_def_list
	: opt_arg_def_list VAR ID sep { $$ = NewNode(NewVarTy, $1); $$.Constant = $3 }
	| opt_arg_def_list VAR MUL ID sep { $$ = NewNode(NewVarTy, $1); $$.Constant = $4; $$.Ptr = true }
	| /* Empty */ { $$ = nil }
	;


optional_type
    :  VAR { $$ = true }
    | /* Empty */ { $$ = false }
    ;


buildins
	: STOP LEFT_PAR RIGHT_PAR { $$ = NewNode(StopTy) }
	/*| CALL LEFT_PAR arguments RIGHT_PAR { $$ = NewNode(CallTy, $3) }*/
	| CALL LEFT_PAR get_variable COMMA get_variable COMMA get_variable COMMA ptr COMMA ptr RIGHT_PAR
	  {
		  $$ = NewNode(CallTy, $3, $5, $7, $9, $11)
	  }
	| TRANSACT LEFT_PAR get_variable COMMA get_variable COMMA ptr RIGHT_PAR
	  {
		  $$ = NewNode(TransactTy, $3, $5, $7)
	  }
	| CREATE LEFT_PAR get_variable COMMA ptr RIGHT_PAR { $$ = NewNode(CreateTy, $3, $5) }
	| SIZEOF LEFT_PAR ID RIGHT_PAR { $$ = NewNode(SizeofTy); $$.Constant = $3 }
	| PUSH LEFT_PAR expression RIGHT_PAR { $$ = NewNode(PushTy, $3) }
	| POP LEFT_PAR RIGHT_PAR { $$ = NewNode(PopTy) }
	| THIS DOT closure_funcs { $$ = $3 }
	;

arguments
	: arguments get_variable sep { $$ = NewNode(ArgTy, $1, $2) }
	| /* Empty */ { $$ = NewNode(EmptyTy) }
	;

sep
	: COMMA { $$ = NewNode(EmptyTy) }
	| /* Empty */ { $$ = NewNode(EmptyTy) }
	;


closure_funcs
	: ORIGIN LEFT_PAR RIGHT_PAR { $$ = NewNode(OriginTy) }
	| ADDRESS LEFT_PAR RIGHT_PAR { $$ = NewNode(AddressTy) }
	| CALLER LEFT_PAR RIGHT_PAR { $$ = NewNode(CallerTy) }
	| CALLVAL LEFT_PAR RIGHT_PAR { $$ = NewNode(CallValTy) }
	| CALLDATALOAD LEFT_BRACKET expression RIGHT_BRACKET { $$ = NewNode(CallDataLoadTy, $3) }
	| CALLDATASIZE LEFT_PAR RIGHT_PAR { $$ = NewNode(CallDataSizeTy) }
	| DIFFICULTY LEFT_PAR RIGHT_PAR { $$ = NewNode(DiffTy) }
	| PREVHASH LEFT_PAR RIGHT_PAR { $$ = NewNode(PrevHashTy) }
	| TIMESTAMP LEFT_PAR RIGHT_PAR { $$ = NewNode(TimestampTy) }
	| GASPRICE LEFT_PAR RIGHT_PAR { $$ = NewNode(GasPriceTy) }
	| BLOCKNUM LEFT_PAR RIGHT_PAR { $$ = NewNode(BlockNumTy) }
	| COINBASE LEFT_PAR RIGHT_PAR { $$ = NewNode(CoinbaseTy) }
	| BALANCE LEFT_PAR RIGHT_PAR { $$ = NewNode(BalanceTy) }
	| GAS LEFT_PAR RIGHT_PAR { $$ = NewNode(GasTy) }
	| STORE LEFT_BRACKET expression RIGHT_BRACKET { $$ = NewNode(StoreTy, $3) }
	| STORE LEFT_BRACKET expression RIGHT_BRACKET ASSIGN expression
	  {
	      node := NewNode(SetStoreTy, $3)
	      $$ = NewNode(AssignmentTy, $6, node)
	  }
	;

if_statement
	: IF expression LEFT_BRACES statement_list RIGHT_BRACES optional_else_statement
	  {
	      if $6 == nil {
		    $$ = NewNode(IfThenTy, $2, $4)
	      } else {
		    $$ = NewNode(IfThenElseTy, $2, $4, $6)
	      }
	  }
	;
optional_else_statement
	: ELSE LEFT_BRACES statement_list RIGHT_BRACES
	  {
	      $$ = $3
	  }
	| /* Empty */ { $$ = nil }
	;

for_statement
	: FOR expression END_STMT expression END_STMT expression LEFT_BRACES statement_list RIGHT_BRACES
	  {
		  $$ = NewNode(ForThenTy, $2, $4, $6, $8)
	  }
	/* TODO */
	| FOR expression END_STMT expression LEFT_BRACES statement_list RIGHT_BRACES
	  {
		  $$ = NewNode(ForThenTy, $2, $4, $6)
	  }
	/* TODO */
	| FOR expression LEFT_BRACES statement_list RIGHT_BRACES
	  {
		  $$ = NewNode(ForThenTy, $2, $4)
	  }
	;

expression
	: op_expression { $$ = $1 }
	| AND ID { $$ = NewNode(RefTy); $$.Constant = $2 }
	| assign_expression { $$ = $1 }
	| ID LEFT_PAR opt_arg_call_list RIGHT_PAR
		{
			$$ = NewNode(FuncCallTy, $3)
			$$.Constant = $1
			$$.ArgList = makeArgs($3, false)
		}
	| RETURN statement { $$ = NewNode(ReturnTy, $2) }
	| EXIT statement { $$ = NewNode(ExitTy, $2) }
	| /* Empty */  { $$ = NewNode(EmptyTy) }
	;

opt_arg_call_list
	: opt_arg_call_list expression sep { $$ = NewNode(ArgTy, $1, $2);}
	| /* Empty */ { $$ = nil }
	;

op_expression
    /* ++, -- */
	: get_id DOP { $$ = NewNode(OpTy, $1); $$.Constant = $2 } 
    /* Everything else */
	| sub_expression OP sub_expression { $$ = NewNode(OpTy, $1, $3); $$.Constant = $2 }
	| sub_expression AND sub_expression { $$ = NewNode(OpTy, $1, $3); $$.Constant = $2 }
	| sub_expression MUL sub_expression { $$ = NewNode(OpTy, $1, $3); $$.Constant = $2 }
	;


sub_expression
    : simple_expression { $$ = $1; }
    | op_expression { $$ = $1; }
    ;

assign_expression
	: MUL ID ASSIGN expression
		{
			node := NewNode(SetPtrTy)
			node.Constant = $2
	      		$$ = NewNode(AssignmentTy, $4, node)
		}
	| ID ASSIGN expression
	  {
	      node := NewNode(SetLocalTy)
	      node.Constant = $1
	      $$ = NewNode(AssignmentTy, $3, node)
	  }
	| ID LEFT_BRACKET expression RIGHT_BRACKET ASSIGN assign_expression
	  {
	      $$ = NewNode(AssignArrayTy, $3, $6); $$.Constant = $1
	  }
	| new_var ASSIGN expression
	  {
	      node := NewNode(SetLocalTy)
	      node.Constant = $1.Constant
	      $$ = NewNode(AssignmentTy, $3, $1, node)
	  }
	| ID COLON ASSIGN expression
	  {
		node := NewNode(SetLocalTy)
		node.Constant = $1
		varNode := NewNode(NewVarTy); varNode.Constant = $1
		$$ = NewNode(AssignmentTy, $4, varNode, node)
	  }
	| new_var { $$ = $1 }
	| new_array { $$ = $1 }
	| simple_expression { $$ = $1 }
	;

new_var
	: VAR ID
	  {
	      $$ = NewNode(NewVarTy)
	      $$.Constant = $2
	  }
	| VAR MUL ID
	  {
	      $$ = NewNode(NewVarTy)
	      $$.Constant = $3
	      $$.Ptr = true
	  }
	;

new_array
	: VAR LEFT_BRACKET NUMBER RIGHT_BRACKET ID
	  {
	      $$ = NewNode(NewArrayTy)
	      //$$.VarType = $1
	      $$.Size = $3
	      $$.Constant = $5

}
	;

simple_expression
	: get_variable { $$ = $1 }
	;

get_variable
	: ptr { $$ = $1 }
	| NUMBER { $$ = NewNode(ConstantTy); $$.Constant = $1 }
	| ID LEFT_BRACKET expression RIGHT_BRACKET { $$ = NewNode(ArrayTy, $3); $$.Constant = $1 }
	| BOOLEAN { $$ = NewNode(BoolTy); $$.Constant = $1 }
	| string { $$ = $1 }
	| buildins { $$ = $1 }
	;

ptr
	: get_id { $$ = $1 }
	| NIL { $$ = NewNode(NilTy) }
	;

get_id
	: ID { $$ = NewNode(IdentifierTy); $$.Constant = $1 }
	;

string 
	: QUOTE STR QUOTE { $$ = NewNode(StringTy); $$.Constant = $2 }
	;

%%

