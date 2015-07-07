package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/obscuren/mutan"
	"github.com/obscuren/mutan/backends"
)

var (
	StrCode       = flag.String("s", "", "code string")
	ByteArray     = flag.Bool("b", false, "output byte array instead of hex")
	ShowAssembler = flag.Bool("asm", false, "output assembler")
	ShowVersion   = flag.Bool("v", false, "show version")

	Version = "0.2"
)

func Panic(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *ShowVersion {
		fmt.Fprintf(os.Stdout, "Mutan %s (EVM 0.6, Frontend %s) (%s, %s)\n", mutan.Version, Version, runtime.GOOS, runtime.Version())
		os.Exit(0)
	}

	var (
		reader   io.Reader
		compiler = mutan.NewCompiler(backend.NewEthereumBackend())
	)

	if len(*StrCode) > 0 {
		reader = strings.NewReader(*StrCode)
	} else {
		if len(flag.Arg(0)) > 0 {
			file, err := os.Open(os.Args[len(os.Args)-1])
			if err != nil {
				Panic("%v\n", err)
			}

			reader = file
		} else {
			Panic("no file specified\n")
		}
	}

	asm, err := compiler.Assemble(reader)
	if err != nil {
		Panic("%v\n", err)
	}

	if *ShowAssembler {
		for _, c := range asm {
			switch c := c.(type) {
			case vm.OpCode:
				fmt.Printf("%v ", c)
			case string:
				fmt.Printf("%s ", c)
			default:
				fmt.Printf("%x ", c)
			}
		}
		fmt.Println()
	} else {
		bytes, err := compiler.AssemblerStage(asm...)
		if err != nil {
			Panic("%v\n", err)
		}

		if *ByteArray {
			fmt.Println(bytes)
		} else {
			fmt.Printf("%x\n", bytes)
		}
	}
}
