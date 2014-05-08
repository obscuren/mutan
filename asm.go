package mutan

import (
	"fmt"
	"math/big"
	"os"
)

// Op codes
var OpCodes = map[string]byte{
	// 0x0 range - arithmetic ops
	"STOP": 0x00,
	"ADD":  0x01,
	"MUL":  0x02,
	"SUB":  0x03,
	"DIV":  0x04,
	"SDIV": 0x05,
	"MOD":  0x06,
	"SMOD": 0x07,
	"EXP":  0x08,
	"NEG":  0x09,
	"LT":   0x0a,
	"GT":   0x0b,
	"EQ":   0x0c,
	"NOT":  0x0d,

	// 0x10 range - bit ops
	"AND":  0x10,
	"OR":   0x11,
	"XOR":  0x12,
	"BYTE": 0x13,

	// 0x20 range - crypto
	"SHA3": 0x20,

	// 0x30 range - closure state
	"ADDRESS":      0x30,
	"BALANCE":      0x31,
	"ORIGIN":       0x32,
	"CALLER":       0x33,
	"CALLVALUE":    0x34,
	"CALLDATALOAD": 0x35,
	"CALLDATASIZE": 0x36,
	"GASPRICE":     0x38,

	// 0x40 range - block operations
	"PREVHASH":   0x40,
	"COINBASE":   0x41,
	"TIMESTAMP":  0x42,
	"NUMBER":     0x43,
	"DIFFICULTY": 0x44,
	"GASLIMIT":   0x45,

	// 0x50 range - 'storage' and execution
	"POP":     0x51,
	"DUP":     0x52,
	"SWAP":    0x53,
	"MLOAD":   0x54,
	"MSTORE":  0x55,
	"MSTORE8": 0x56,
	"SLOAD":   0x57,
	"SSTORE":  0x58,
	"JUMP":    0x59,
	"JUMPI":   0x5a,
	"PC":      0x5b,
	"MSIZE":   0x5c,

	// 0x60 range - 'push'
	"PUSH1":  0x60,
	"PUSH2":  0x61,
	"PUSH3":  0x62,
	"PUSH4":  0x63,
	"PUSH5":  0x64,
	"PUSH6":  0x65,
	"PUSH7":  0x66,
	"PUSH8":  0x67,
	"PUSH9":  0x68,
	"PUSH10": 0x69,
	"PUSH11": 0x6a,
	"PUSH12": 0x6b,
	"PUSH13": 0x6c,
	"PUSH14": 0x6d,
	"PUSH15": 0x6e,
	"PUSH16": 0x6f,
	"PUSH17": 0x70,
	"PUSH18": 0x71,
	"PUSH19": 0x72,
	"PUSH20": 0x73,
	"PUSH21": 0x74,
	"PUSH22": 0x75,
	"PUSH23": 0x76,
	"PUSH24": 0x77,
	"PUSH25": 0x78,
	"PUSH26": 0x79,
	"PUSH27": 0x7a,
	"PUSH28": 0x7b,
	"PUSH29": 0x7c,
	"PUSH30": 0x7d,
	"PUSH31": 0x7e,
	"PUSH32": 0x7f,

	// 0xf0 range - closures
	"CREATE": 0xf0,
	"CALL":   0xf1,
	"RETURN": 0xf2,

	// 0x70 range - other
	"LOG":     0x70,
	"SUICIDE": 0x7f,
}

// Big to bytes
//
// Returns the bytes of a big integer with the size specified by
// **base**
// Attempts to pad the byte array with zeros.
func bigToBytes(num *big.Int, base int) []byte {
	ret := make([]byte, base/8)

	return append(ret[:len(ret)-len(num.Bytes())], num.Bytes()...)
}

// Is op code
//
// Check whether the given string matches anything in
// the OpCode list
func IsOpCode(s string) bool {
	for key, _ := range OpCodes {
		if key == s {
			return true
		}
	}
	return false
}

// Compile instruction
//
// Attempts to compile and parse the given instruction in "s"
// and returns the byte sequence
func CompileInstr(s interface{}) ([]byte, error) {
	switch s.(type) {
	case string:
		str := s.(string)
		isOp := IsOpCode(str)
		if isOp {
			return []byte{OpCodes[str]}, nil
		}

		// Check for pre formatted byte array
		// Jumps are preformatted
		if []byte(str)[0] == 0 {
			return []byte(str), nil
		}

		num := new(big.Int)
		_, success := num.SetString(str, 0)
		// Assume regular bytes during compilation
		if !success {
			num.SetBytes([]byte(str))
		}

		return num.Bytes(), nil
	case int:
		//num := bigToBytes(big.NewInt(int64(s.(int))), 256)
		return big.NewInt(int64(s.(int))).Bytes(), nil
	case []byte:
		return new(big.Int).SetBytes(s.([]byte)).Bytes(), nil
	}

	return nil, nil
}

// Assemble
//
// Assembles the given instructions and returns EVM byte code
func Assemble(instructions ...interface{}) (script []byte) {
	//script = make([]string, len(instructions))

	for _, val := range instructions {
		instr, err := CompileInstr(val)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(instr) == 0 {
			instr = []byte{0}
		}

		//script[i] = string(instr)
		script = append(script, instr...)
	}

	return
}

// Pre process script
//
// Take data apart and attempt to find the "init" section and
// "main" section. `main { } init { }`
func PreProcess(data string) (mainInput, initInput string) {
	mainInput = getCodeSectionFor("main", data)
	if mainInput == "" {
		mainInput = data
	}
	initInput = getCodeSectionFor("init", data)

	return
}

// Very, very dumb parser. Heed no attention :-)
func getCodeSectionFor(blockMatcher, input string) string {
	curCount := -1
	length := len(blockMatcher)
	matchfst := rune(blockMatcher[0])
	var currStr string

	for i, run := range input {
		// Find init
		if curCount == -1 && run == matchfst && input[i:i+length] == blockMatcher {
			curCount = 0
		} else if curCount > -1 {
			if run == '{' {
				curCount++
				if curCount == 1 {
					continue
				}
			} else if run == '}' {
				curCount--
				if curCount == 0 {
					// we are done
					curCount = -1
					break
				}
			}

			if curCount > 0 {
				currStr += string(run)
			}
		}
	}

	return currStr
}
