package frontend

/*
int code generator takes care of memory allocation
generate all opcode and instructions
sets pc for each instruction afterwards

compiler transforms int code to ASM (very static)
*/

import (
	"fmt"
	"io/ioutil"
	"os"
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
			if v[i] != nil && v[i].Code != IntEmpty {
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
	if blk2.Code == IntEmpty {
		return blk1
	}

	if blk1.Code == IntEmpty {
		return blk2
	}

	search := blk1
	for search.Next != nil {
		search = search.Next
	}

	search.Next = blk2

	return blk1
}
func Concat(a, b *IntInstr) *IntInstr {
	return concat(a, b)
}

func newJumpInstr(op Instr) (*IntInstr, *IntInstr) {
	push := newIntInstr(IntPush4, "")
	cons := newIntInstr(IntConst, "")
	cons.size = 4
	cons.Target = push

	jump := newIntInstr(op, "")
	jump.TargetNum = cons
	concat(push, cons)
	concat(cons, jump)

	return push, jump
}

// Recursive Intermediate code generator
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
		not := newIntInstr(IntNot, "")
		p, jmpi := newJumpInstr(IntJumpi)
		then := gen.MakeIntCode(tree.Children[1])
		target := newIntInstr(IntTarget, "")
		jmpi.Target = target

		concat(cond, not)
		concat(not, p)
		concat(jmpi, then)
		concat(then, target)

		return cond
	case IfThenElseTy:
		cond := gen.MakeIntCode(tree.Children[0])
		// TODO reverse "else/if" in order to get rid of the "NOT"
		not := newIntInstr(IntNot, "")
		//jump2else := newIntInstr(IntJumpi, "")
		p, jump2else := newJumpInstr(IntJumpi)
		then := gen.MakeIntCode(tree.Children[1])
		//jump2end := newIntInstr(IntJump, "")
		p2, jump2end := newJumpInstr(IntJump)
		elsetarget := newIntInstr(IntTarget, "")
		elsethen := gen.MakeIntCode(tree.Children[2])
		end := newIntInstr(IntTarget, "")

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
		ignore = newIntInstr(IntIgnore, "")

		// Cast to not
		not := newIntInstr(IntNot, "")
		// Jump to end if statement is false
		p, jmpi := newJumpInstr(IntJumpi)
		// Jump back to the start of the loop (targetBack)
		p2, jmp := newJumpInstr(IntJump)
		// Target for the conditional jump (jmpi)
		target := newIntInstr(IntTarget, "")

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

		return concat(blk1, newIntInstr(IntEqual, ""))
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
		blk2 := newIntInstr(IntSStore, "")

		return concat(blk1, blk2)
	case StoreTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		blk2 := newIntInstr(IntSLoad, "")

		return concat(blk1, blk2)

	case OpTy:
		// TODO clean this up
		var blk1, blk2, blk3 *IntInstr
		blk1 = gen.MakeIntCode(tree.Children[0])

		var op Instr
		switch tree.Constant {
		case "==":
			op = IntEqual
		case ">":
			op = IntGt
		case "<":
			op = IntLt
		case "*":
			op = IntMul
		case "/":
			op = IntDiv
		case "+":
			op = IntAdd
		case "-":
			op = IntSub
		case "%":
			op = IntMod
		case "&":
			op = IntAnd
		case "|":
			op = IntOr
		case "^":
			op = IntXor
		case "**":
			op = IntExp
		case "<<", ">>":
			push, cons := pushConstant("2")
			blk2 := gen.MakeIntCode(tree.Children[1])
			exp := newIntInstr(IntExp, "")
			var o *IntInstr
			if tree.Constant == "<<" {
				o = newIntInstr(IntMul, "")
			} else {
				o = newIntInstr(IntDiv, "")
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
				op = IntAdd
			} else {
				op = IntSub
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
				c, err := gen.Errorf(tree, "++ only supported on identifiers")
				gen.addError(err)
				return c
			}

			return blk1
		case ">=":
			op = IntLt
			blk3 = newIntInstr(IntNot, "")
		case "<=":
			op = IntGt
			blk3 = newIntInstr(IntNot, "")
		case "!=":
			op = IntEqual
			blk3 = newIntInstr(IntNot, "")
		case "!":
			// Reconstruct this one (We ought to clean this code)
			opinstr := newIntInstr(IntNot, "")
			return concat(blk1, opinstr)
		case "&&":
			var createdTarget bool
			if gen.sharedJumpTarget == nil {
				createdTarget = true
				gen.sharedJumpTarget = newIntInstr(IntTarget, "")
				defer func() { gen.sharedJumpTarget = nil }()
			}

			push, cons := pushConstant("1")
			lh := gen.MakeIntCode(tree.Children[0])
			not := newIntInstr(IntNot, "")
			p, jmpi := newJumpInstr(IntJumpi)
			pop := newIntInstr(IntPop, "")
			rh := gen.MakeIntCode(tree.Children[1])
			jmpi.Target = gen.sharedJumpTarget

			concat(push, cons)
			concat(cons, lh)
			concat(lh, not)
			concat(not, p)
			concat(p, pop)
			concat(pop, rh)

			if createdTarget {
				concat(rh, gen.sharedJumpTarget)
			}

			return push
		case "||":
			push, cons := pushConstant("1")
			lh := gen.MakeIntCode(tree.Children[0])
			p, jmpi := newJumpInstr(IntJumpi)
			pop := newIntInstr(IntPop, "")
			rh := gen.MakeIntCode(tree.Children[1])
			target := newIntInstr(IntTarget, "")
			jmpi.Target = target

			concat(push, cons)
			concat(cons, lh)
			concat(lh, p)
			concat(p, pop)
			concat(pop, rh)
			concat(rh, target)

			return push

		default:
			c, err := gen.Errorf(tree, "Expected operator, got '%v'", tree.Constant)
			gen.addError(err)
			return c
		}

		blk2 = gen.MakeIntCode(tree.Children[1])
		concat(blk2, blk1)

		if blk3 != nil {
			opinstr := newIntInstr(op, "")
			cc(blk2, opinstr, blk3)

			return blk2
		}

		return concat(blk2, newIntInstr(op, ""))
	case StringTy:
		return gen.makeString(tree)
	case StopTy:
		return newIntInstr(IntStop, "")
	case OriginTy:
		return newIntInstr(IntOrigin, "")
	case AddressTy:
		return newIntInstr(IntAddress, "")
	case CallerTy:
		return newIntInstr(IntCaller, "")
	case CallValTy:
		return newIntInstr(IntCallVal, "")
	case CallDataLoadTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		push, cons := pushConstant("32")
		mul := newIntInstr(IntMul, "")
		blk2 := newIntInstr(IntCallDataLoad, "")

		concat(blk1, push)
		concat(push, cons)
		concat(cons, mul)
		concat(mul, blk2)

		return blk1
	case CallDataSizeTy:
		return newIntInstr(IntCallDataSize, "")
	case GasPriceTy:
		return newIntInstr(IntGasPrice, "")
	case DiffTy:
		return newIntInstr(IntDiff, "")
	case PrevHashTy:
		return newIntInstr(IntPrevHash, "")
	case TimestampTy:
		return newIntInstr(IntTimestamp, "")
	case CoinbaseTy:
		return newIntInstr(IntCoinbase, "")
	case BalanceTy:
		blk1 := gen.MakeIntCode(tree.Children[0])
		return concat(blk1, newIntInstr(IntBalance, ""))
	case GasTy:
		return newIntInstr(IntGas, "")
	case BlockNumTy:
		return newIntInstr(IntBlockNum, "")
	case ByteTy:
		value := gen.MakeIntCode(tree.Children[0])
		th := gen.MakeIntCode(tree.Children[1])
		b := newIntInstr(IntByte, "")

		cc(value, th, b)

		return value
	case Sha3Ty:
		l := gen.MakeIntCode(tree.Children[1])
		addr, err := gen.getMemoryAddress(tree.Children[0])
		if err != nil {
			gen.addError(err)
			return addr
		}
		sha3 := newIntInstr(IntSha3, "")

		cc(l, addr, sha3)

		return l
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

		return newIntInstr(IntIgnore, "")
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
		call := newIntInstr(IntCall, "")

		cc(ret, arg, value, sender, gas, call)

		return ret
	case TransactTy:
		ret := gen.pushNil()
		arg, err := gen.makeArg(tree.Children[3])
		if err != nil {
			gen.addError(err)
			return arg
		}

		value := gen.MakeIntCode(tree.Children[2])
		sender := gen.MakeIntCode(tree.Children[0])
		gas := gen.MakeIntCode(tree.Children[1])
		call := newIntInstr(IntCall, "")

		cc(ret, arg, value, sender, gas, call)

		return ret
	case CreateTy:
		script := gen.MakeIntCode(tree.Children[1])
		val := gen.MakeIntCode(tree.Children[0])
		create := newIntInstr(IntCreate, "")

		concat(script, val)
		concat(val, create)

		return script
	case ScopeTy:
		// TODO
		return gen.MakeIntCode(tree.Children[0])
	case ReturnTy:
		if gen.CurrentScope() != gen {
			return gen.CurrentScope().MakeReturn(tree.Children[0], gen)
		}

		// Fall through to exit if this is the global scope
		fallthrough
	case ExitTy:
		switch tree.Children[0].Type {
		case LambdaTy:
			ret := gen.MakeIntCode(tree.Children[0])
			cc(ret, newIntInstr(IntReturn, ""))

			return ret
		case ConstantTy:
			cons := gen.makePush(tree.Children[0].Constant)
			store := makeStore(0)
			size := gen.makePush("32")
			offset := gen.makePush("0")

			concat(cons, store)
			concat(store, size)
			concat(size, offset)
			concat(offset, newIntInstr(IntReturn, ""))

			return cons
		case IdentifierTy:
			variable := gen.GetVar(tree.Children[0].Constant)
			if variable == nil {
				i, err := gen.Errorf(tree, "Undefined variable: %v", tree.Constant)
				gen.addError(err)

				return i
			}

			offset := strconv.Itoa(variable.Offset())
			size := gen.makePush("32")
			offs := gen.makePush(offset)
			concat(size, offs)
			concat(offs, newIntInstr(IntReturn, ""))

			return size

		default: //case StoreTy:
			blk1 := gen.MakeIntCode(tree.Children[0])
			store := makeStore(0)
			size := gen.makePush("32")
			offset := gen.makePush("0")

			concat(blk1, store)
			concat(store, size)
			concat(size, offset)
			concat(offset, newIntInstr(IntReturn, ""))

			return blk1
			/*
				default:
					c, err := gen.Errorf(tree, "return; '%v' not (yet) supported", tree.Children[0].Type)
					gen.addError(err)

					return c
			*/
		}

		return newIntInstr(IntIgnore, "")
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

		firstInstr := newIntInstr(IntASM, asmSlice[0])
		if len(asmSlice) > 1 {
			for _, instr := range asmSlice[1:] {
				concat(firstInstr, newIntInstr(IntASM, instr))
			}
		}

		return firstInstr
	case PushTy:
		return gen.MakeIntCode(tree.Children[0])
	case PopTy:
		//c, err := gen.
		return newIntInstr(IntIgnore, "")
	case FuncCallTy:
		// Look up function
		fn := gen.functionTable[tree.Constant]
		if fn == nil {
			c, err := gen.Errorf(tree, "undefine: '%s'", tree.Constant)
			gen.addError(err)
			return c
		}

		return fn.Call(tree.ArgList, gen, gen.CurrentScope())
	case FuncDefTy:
		callTarget := newIntInstr(IntTarget, "")
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
			var t varType
			if i.Ptr {
				t = varPtrTy
			}
			v, _ := fn.NewVar(i.Constant, t)
			fn.PushArg(v)
		}

		gen.PushScope(fn)

		p, jmp := newJumpInstr(IntJump)
		target := newIntInstr(IntTarget, "")
		jmp.Target = target

		body := gen.MakeIntCode(tree.Children[0])

		// Pop frame mechanism
		ptr := gen.loadStackPtr()
		dup := newIntInstr(IntDup1, "")
		// Load stack pointer as offset
		dup2 := newIntInstr(IntDup1, "")
		// Increment by 1 word
		offset := gen.makePush("32")
		add := newIntInstr(IntAdd, "")
		sizeLoad := newIntInstr(IntMLoad, "")
		// Now pop the frame off the stack
		stackPtrOffset := gen.makePush("0")
		stackPtrStore := newIntInstr(IntMStore, "")
		retLoad := newIntInstr(IntMLoad, "")
		jumpBack := newIntInstr(IntJump, "")

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
		file, err := os.Open(tree.Constant)
		if err != nil {
			instr, e := gen.Errorf(tree, "import: %v", err)
			gen.addError(e)
			return instr
		}

		buff, err := ioutil.ReadAll(file)
		if err != nil {
			instr, e := gen.Errorf(tree, "import buff: %v", err)
			gen.addError(e)
			return instr
		}

		ast, err := MakeAst(string(buff))
		if err != nil {
			instr, e := gen.Errorf(tree, "import: %v", err)
			gen.addError(e)
			return instr
		}

		ofn := gen.filename
		gen.filename = tree.Constant
		instr := gen.MakeIntCode(ast)
		gen.filename = ofn

		return instr
	case LambdaTy:
		instr, l := gen.compileLambda(0, tree)

		size := gen.makePush(strconv.Itoa(l))
		/*
			s := ((l+31)/32*32 - l)

			dup := gen.makePush(strconv.Itoa(s))
			offset := newIntInstr(IntMSize, "")
			sub := newIntInstr(IntSub, "")

			cc(instr, size, dup, offset, sub)
		*/

		cc(instr, size, gen.makePush("0"))

		return instr

	case SuicideTy:
		receiverAddr := gen.MakeIntCode(tree.Children[0])
		suicide := newIntInstr(IntSuicide, "")

		return cc(receiverAddr, suicide)
	case EmptyTy:
		return newIntInstr(IntEmpty, "")
	}

	return nil
}
