package mutan

/*
int code generator takes care of memory allocation
generate all opcode and instructions
sets pc for each instruction afterwards

compiler transforms int code to ASM (very static)
*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ___unused() {
	fmt.Println("")
}

func cc(v ...*IntInstr) *IntInstr {
	if len(v) > 0 {
		var begin, last *IntInstr
		for i := 0; i < len(v); i++ {
			if v[i] != nil && v[i].Code != intEmpty {
				if begin == nil {
					begin = v[i]
					last = begin
					continue
				}

				cca(last, v[i])
				last = v[i]
			}
		}

		return begin
	}

	return nil
}

func cca(a, b *IntInstr) *IntInstr {
	search := a
	for search.Next != nil {
		search = search.Next
	}

	search.Next = b

	return a
}

// Concatenate two block of code together
func concat(blk1 *IntInstr, blk2 *IntInstr) *IntInstr {
	if blk2.Code == intEmpty {
		return blk1
	}

	if blk1.Code == intEmpty {
		return blk2
	}

	search := blk1
	for search.Next != nil {
		search = search.Next
	}

	search.Next = blk2

	return blk1
}

func newJumpInstr(op Instr) (*IntInstr, *IntInstr) {
	push := newIntInstr(intPush4, "")
	cons := newIntInstr(intConst, "")
	cons.size = 4
	cons.Target = push

	jump := newIntInstr(op, "")
	jump.TargetNum = cons
	concat(push, cons)
	concat(cons, jump)

	return push, jump
}

// Recursive intermediate code generator
func (gen *IntGen) MakeIntCode(tree *SyntaxTree) *IntInstr {
	switch tree.Type {
	case StatementListTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])

		return concat(blk1, blk2)
	case AssignmentTy:
		var blk1 *IntInstr
		if len(tree.Children) == 2 {
			return gen.setVariable(tree.Children[0], tree.Children[1])
		} else {
			blk2 := gen.MakeIntCode(tree.Children[1])
			blk1 = gen.setVariable(tree.Children[0], tree.Children[2])

			concat(blk1, blk2)
		}

		return blk1
	case IfThenTy:
		cond := gen.MakeIntCode(tree.Children[0])
		not := newIntInstr(intNot, "")
		//jmpi := newIntInstr(intJumpi, "")
		p, jmpi := newJumpInstr(intJumpi)
		then := gen.MakeIntCode(tree.Children[1])
		target := newIntInstr(intTarget, "")
		jmpi.Target = target

		concat(cond, not)
		concat(not, p)
		concat(jmpi, then)
		concat(then, target)

		return cond
	case IfThenElseTy:
		cond := gen.MakeIntCode(tree.Children[0])
		// TODO reverse "else/if" in order to get rid of the "NOT"
		not := newIntInstr(intNot, "")
		//jump2else := newIntInstr(intJumpi, "")
		p, jump2else := newJumpInstr(intJumpi)
		then := gen.MakeIntCode(tree.Children[1])
		//jump2end := newIntInstr(intJump, "")
		p2, jump2end := newJumpInstr(intJump)
		elsetarget := newIntInstr(intTarget, "")
		elsethen := gen.MakeIntCode(tree.Children[2])
		end := newIntInstr(intTarget, "")

		jump2end.Target = end
		jump2else.Target = elsetarget

		concat(cond, not)
		concat(not, p)
		concat(jump2else, then)
		concat(then, p2)
		concat(jump2end, elsetarget)
		concat(elsetarget, elsethen)
		concat(elsethen, end)

		return cond
	case ForThenTy:
		var ignore, init, cond, then, aft *IntInstr
		ignore = newIntInstr(intIgnore, "")

		// Cast to not
		not := newIntInstr(intNot, "")
		// Jump to end if statement is false
		p, jmpi := newJumpInstr(intJumpi)
		// Jump back to the start of the loop (targetBack)
		p2, jmp := newJumpInstr(intJump)
		// Target for the conditional jump (jmpi)
		target := newIntInstr(intTarget, "")

		switch len(tree.Children) {
		case 2: // while
			// the condition
			cond = gen.MakeIntCode(tree.Children[0])
			// block
			then = gen.MakeIntCode(tree.Children[1])
			concat(ignore, cond)
			concat(cond, not)
			concat(not, p)
			concat(jmpi, then)
			concat(then, p2)
			concat(jmp, target)
		case 4: // for
			// Init part
			init = gen.MakeIntCode(tree.Children[0])
			// The condition for the loop
			cond = gen.MakeIntCode(tree.Children[1])
			// Iterator
			aft = gen.MakeIntCode(tree.Children[2])
			// Body of the loop
			then = gen.MakeIntCode(tree.Children[3])

			concat(ignore, init)
			concat(init, cond)
			concat(cond, not)
			concat(not, p)
			concat(jmpi, then)
			concat(then, aft)
			concat(aft, p2)
			concat(jmp, target)
		}
		// Set targets
		jmpi.Target = target
		jmp.Target = cond

		return ignore
	case EqualTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := gen.MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		return concat(blk1, newIntInstr(intEqual, ""))
	case IdentifierTy:
		c, err := gen.getMemory(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case RefTy:
		c, err := gen.getPtr(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case ConstantTy:
		blk1, blk2 := pushConstant(tree.Constant)
		concat(blk1, blk2)

		return blk1
	case BoolTy:
		var value string
		if tree.Constant == "true" {
			value = "1"
		} else {
			value = "0"
		}
		blk1, blk2 := pushConstant(value)
		concat(blk1, blk2)

		return blk1
	case SetLocalTy:
		c, err := gen.setMemory(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case DerefPtrTy:
		c, err := gen.pushDerefPtr(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case SetStoreTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := newIntInstr(intSStore, "")

		return concat(blk1, blk2)
	case StoreTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := newIntInstr(intSLoad, "")

		return concat(blk1, blk2)

	case OpTy:
		// TODO clean this up
		var blk1, blk2, blk3 *IntInstr
		blk1 = gen.MakeIntCode(tree.Children[0])

		var op Instr
		switch tree.Constant {
		case "==":
			op = intEqual
		case ">":
			op = intGt
		case "<":
			op = intLt
		case "*":
			op = intMul
		case "/":
			op = intDiv
		case "+":
			op = intAdd
		case "-":
			op = intSub
		case "%":
			op = intMod
		case "&":
			op = intAnd
		case "|":
			op = intOr
		case "^":
			op = intXor
		case "**":
			op = intExp
		case "<<", ">>":
			push, cons := pushConstant("2")
			blk2 := gen.MakeIntCode(tree.Children[1])
			exp := newIntInstr(intExp, "")
			var o *IntInstr
			if tree.Constant == "<<" {
				o = newIntInstr(intMul, "")
			} else {
				o = newIntInstr(intDiv, "")
			}
			concat(push, cons)
			concat(cons, blk2)
			concat(blk2, exp)
			concat(exp, blk1)
			concat(blk1, o)

			return push
		case "++", "--":
			one, cons := pushConstant("1")
			if tree.Constant == "++" {
				op = intAdd
			} else {
				op = intSub
			}
			opInstr := newIntInstr(op, "")
			concat(blk1, one)
			concat(one, cons)
			concat(cons, opInstr)
			// Get the child
			child := tree.Children[0]
			if child.Type == IdentifierTy {
				c, err := gen.setMemory(child)
				if err != nil {
					gen.addError(err)
					return c
				}

				concat(opInstr, c)
			} else {
				// TODO?
				c, err := tree.Errorf("++ only supported on identifiers")
				gen.addError(err)
				return c
			}

			return blk1
		case ">=":
			op = intLt
			blk3 = newIntInstr(intNot, "")
		case "<=":
			op = intGt
			blk3 = newIntInstr(intNot, "")
		case "!=":
			op = intEqual
			blk3 = newIntInstr(intNot, "")
		case "!":
			// Reconstruct this one (We ought to clean this code)
			blk1 = gen.MakeIntCode(tree.Children[1])
			opinstr := newIntInstr(intNot, "")
			return concat(blk1, opinstr)
		case "&&":

		default:
			c, err := tree.Errorf("Expected operator, got '%v'", tree.Constant)
			gen.addError(err)
			return c
		}

		blk2 = gen.MakeIntCode(tree.Children[1])
		concat(blk1, blk2)

		if blk3 != nil {
			opinstr := newIntInstr(op, "")
			concat(blk2, opinstr)
			concat(opinstr, blk3)

			return blk1
		}

		return concat(blk1, newIntInstr(op, ""))
	case StringTy:
		return gen.makeString(tree)
	case StopTy:
		return newIntInstr(intStop, "")
	case OriginTy:
		return newIntInstr(intOrigin, "")
	case AddressTy:
		return newIntInstr(intAddress, "")
	case CallerTy:
		return newIntInstr(intCaller, "")
	case CallValTy:
		return newIntInstr(intCallVal, "")
	case CallDataLoadTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		push, cons := pushConstant("32")
		mul := newIntInstr(intMul, "")
		blk2 := newIntInstr(intCallDataLoad, "")

		concat(blk1, push)
		concat(push, cons)
		concat(cons, mul)
		concat(mul, blk2)

		return blk1
	case CallDataSizeTy:
		return newIntInstr(intCallDataSize, "")
	case GasPriceTy:
		return newIntInstr(intGasPrice, "")
	case DiffTy:
		return newIntInstr(intDiff, "")
	case PrevHashTy:
		return newIntInstr(intPrevHash, "")
	case TimestampTy:
		return newIntInstr(intTimestamp, "")
	case CoinbaseTy:
		return newIntInstr(intCoinbase, "")
	case BalanceTy:
		return newIntInstr(intBalance, "")
	case GasTy:
		return newIntInstr(intGas, "")
	case BlockNumTy:
		return newIntInstr(intBlockNum, "")
	case NewArrayTy: // Create a new array
		c, err := gen.initNewArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case ArrayTy: // Get a value of an array
		c, err := gen.getArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case AssignArrayTy: // Assign a value to an element in the array
		c, err := gen.setArray(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case NewVarTy:
		c, err := gen.initNewNumber(tree)
		if err != nil {
			gen.addError(err)

			return c
		}

		return newIntInstr(intIgnore, "")
	case ArgTy:
		next := gen.MakeIntCode(tree.Children[0])
		arg := gen.MakeIntCode(tree.Children[1])

		return concat(arg, next)
	case CallTy:
		arg, err := gen.makeArg(tree.Children[3])
		if err != nil {
			gen.addError(err)
			return arg
		}
		ret, err := gen.makeArg(tree.Children[4])
		if err != nil {
			gen.addError(err)
			return ret
		}
		sender := gen.MakeIntCode(tree.Children[0])
		value := gen.MakeIntCode(tree.Children[1])
		gas := gen.MakeIntCode(tree.Children[2])
		call := newIntInstr(intCall, "")

		concat(ret, arg)
		concat(arg, gas)
		concat(gas, value)
		concat(value, sender)
		concat(sender, call)

		return ret
	case TransactTy:
		ret := gen.pushNil()
		arg, err := gen.makeArg(tree.Children[2])
		if err != nil {
			gen.addError(err)
			return arg
		}
		sender := gen.MakeIntCode(tree.Children[0])
		value := gen.MakeIntCode(tree.Children[1])
		gas := gen.makePush("0")
		call := newIntInstr(intCall, "")

		concat(ret, arg)
		concat(arg, gas)
		concat(gas, value)
		concat(value, sender)
		concat(sender, call)

		return ret
	case CreateTy:
		script, err := gen.makeArg(tree.Children[1])
		if err != nil {
			gen.addError(err)
			return script
		}
		val := gen.MakeIntCode(tree.Children[0])
		create := newIntInstr(intCreate, "")

		concat(script, val)
		concat(val, create)

		return script
	case ReturnTy:
		return gen.CurrentScope().MakeReturn(tree.Children[0], gen)
	case ExitTy:
		switch tree.Children[0].Type {
		case LambdaTy:
			retVal, num := gen.compileLambda(0, tree.Children[0])
			if num != 0 {
				size := gen.makePush("0")
				offset := gen.makePush(strconv.Itoa(num))
				concat(retVal, offset)
				concat(offset, size)

				return concat(retVal, newIntInstr(intReturn, ""))
			}
		case ConstantTy:
			cons := gen.makePush(tree.Children[0].Constant)
			store := makeStore(0)
			size := gen.makePush("32")
			offset := gen.makePush("0")

			concat(cons, store)
			concat(store, size)
			concat(size, offset)
			concat(offset, newIntInstr(intReturn, ""))

			return cons
		case StoreTy:
			blk1 := gen.MakeIntCode(tree.Children[0])
			store := makeStore(0)
			size := gen.makePush("32")
			offset := gen.makePush("0")

			concat(blk1, store)
			concat(store, size)
			concat(size, offset)
			concat(offset, newIntInstr(intReturn, ""))

			return blk1
		case IdentifierTy:
			variable := gen.GetVar(tree.Children[0].Constant)
			if variable == nil {
				i, err := tree.Errorf("Undefined variable: %v", tree.Constant)
				gen.addError(err)

				return i
			}

			offset := strconv.Itoa(variable.Offset())
			size := gen.makePush("32")
			offs := gen.makePush(offset)
			concat(size, offs)
			concat(offs, newIntInstr(intReturn, ""))

			return size
		default:
			c, err := tree.Errorf("return; '%v' not (yet) supported", tree.Children[0].Type)
			gen.addError(err)

			return c
		}

		return newIntInstr(intIgnore, "")
	case SizeofTy:
		c, err := gen.sizeof(tree)
		if err != nil {
			gen.addError(err)
		}

		return c
	case InlineAsmTy:
		// Remove tabs
		asm := strings.Replace(tree.Constant, "\t", "", -1)
		// Remove the comments
		re := regexp.MustCompile(";.*")
		asm = re.ReplaceAllString(asm, "")
		// Remove double spaces
		asm = strings.Replace(asm, "  ", " ", -1)
		// Split by \n or single space
		asmSlice := strings.FieldsFunc(asm, func(r rune) bool {
			switch r {
			case '\n', ' ':
				return true
			}

			return false
		})

		firstInstr := newIntInstr(intASM, asmSlice[0])
		if len(asmSlice) > 1 {
			for _, instr := range asmSlice[1:] {
				concat(firstInstr, newIntInstr(intASM, instr))
			}
		}

		return firstInstr
	case PushTy:
		return gen.MakeIntCode(tree.Children[0])
	case PopTy:
		//c, err := gen.
		return newIntInstr(intIgnore, "")
	case FuncCallTy:
		// Look up function
		fn := gen.functionTable[tree.Constant]
		if fn == nil {
			c, err := tree.Errorf("undefine: '%s'", tree.Constant)
			gen.addError(err)
			return c
		}

		return fn.Call(tree.ArgList, gen, gen.CurrentScope())
	case FuncDefTy:
		callTarget := newIntInstr(intTarget, "")
		fn := NewFunction(tree.Constant, callTarget, 0, tree.HasRet)
		fn.ArgCount = len(tree.ArgList)
		/****
		 * Stack frame
		 *
		 * | ret pos | frame size | return val |
		 */
		fn.NewVar("___retPtr", varNumTy)
		fn.NewVar("___frameSize", varNumTy)

		for _, i := range tree.ArgList {
			v, _ := fn.NewVar(i.Constant, varNumTy)
			fn.PushArg(v)
		}

		gen.PushScope(fn)

		p, jmp := newJumpInstr(intJump)
		target := newIntInstr(intTarget, "")
		jmp.Target = target

		body := gen.MakeIntCode(tree.Children[0])

		// Pop frame mechanism
		ptr := gen.loadStackPtr()
		dup := newIntInstr(intDup, "")
		// Load stack pointer as offset
		dup2 := newIntInstr(intDup, "")
		// Increment by 1 word
		offset := gen.makePush("32")
		add := newIntInstr(intAdd, "")
		sizeLoad := newIntInstr(intMLoad, "")
		// Now pop the frame off the stack
		stackPtrOffset := gen.makePush("0")
		stackPtrStore := newIntInstr(intMStore, "")
		retLoad := newIntInstr(intMLoad, "")
		jumpBack := newIntInstr(intJump, "")

		concat(jmp, callTarget)
		concat(callTarget, body)
		concat(body, ptr)
		concat(ptr, dup)
		concat(dup, dup2)
		concat(dup2, offset)
		concat(offset, add)
		concat(add, sizeLoad)
		concat(sizeLoad, stackPtrOffset)
		concat(stackPtrOffset, stackPtrStore)
		concat(stackPtrStore, retLoad)
		concat(retLoad, jumpBack)
		concat(jumpBack, target)

		gen.PopScope()

		// TODO do checking if function exists
		gen.functionTable[tree.Constant] = fn

		return p
	case ImportTy:
		return newIntInstr(intIgnore, "")
	case LambdaTy:
		panic("auto lambda triggered in int code gen. report this issue")
	case EmptyTy:
		return newIntInstr(intEmpty, "")
	}

	return nil
}
