#### Compiler & Language definition for the Ethereum project

This is the Mutan 2.0 branch.

Mutan is a C-Like language for the Ethereum project. Mutan supports a
full, statically typed higher level language that compiles to native
Ethereum Assembler. The language definition and documentation
can be found on [go-ethereum wiki](https://github.com/ethereum/go-ethereum/wiki/Mutan).

A simple online editor and compiler can be found [here](http://mutan.jeffew.com)


### Syntax

```go
func fn(var a, var b) {
	var[2] c
	c[0] = a
	c[1] = b
	return c[1]
}

var a = fn(0, 1)
b := 10

if a > b {
    exit()
} else {
    // :-)
    if !a {
        if this.data[0] ^ 10 >= 10 {
            this.data[0] = 1000;
        }
    }
}

this.store[a] = 10000
this.store[b] = this.origin()

for i := 0; i < 10; i++ {
    var[10] out
    call(0xaabbccddeeff112233445566, 0, 10000, i, out)
}

// tx without input data
transact(0xa78f6abe, 10000, nil)
// no args and return values
call(0xab, 0, 10000, nil, nil)
// create contract
var ret = create(value, 0xaabbccddeeff0099887766552211)

asm (
    PUSH1 10
    PUSH1 0
    MSTORE
)
```

Mutan &copy; Jeffrey Wilcke
