package main

import (
	"flag"
	"fmt"
	"github.com/obscuren/mutan"
	"github.com/obscuren/mutan/backends"
	"io"
	"os"
	"runtime"
	"strings"
)

var (
	StrCode       = flag.String("s", "", "code string")
	ByteArray     = flag.Bool("b", false, "output byte array instead of hex")
	ShowAssembler = flag.Bool("asm", false, "output assembler")
	Version       = flag.Bool("v", false, "show version")
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

	if *Version {
		fmt.Fprintf(os.Stdout, "Mutan version 0.4 (EVM 0.5) (%s)\n", runtime.GOOS)
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
		s := fmt.Sprintln(asm)

		fmt.Println(s[1 : len(s)-2])
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
