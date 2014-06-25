package main

import (
	"flag"
	"fmt"
	"github.com/obscuren/mutan"
	"github.com/obscuren/mutan/backends"
	"io"
	"os"
	"strings"
)

var (
	StrCode   = flag.String("s", "", "code string")
	ByteArray = flag.Bool("b", false, "output byte array instead of hex")

	ShowAssembler = flag.Bool("asm", false, "output assembler")
)

func Panic(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func Compile(compiler *mutan.Compiler, reader io.Reader) (asm []interface{}, errors []error) {
	code, err := compiler.ReadAll(reader)
	if err != nil {
		errors = append(errors, err)
		return
	}

	code, err = compiler.PreProcessorStage(code)
	if err != nil {
		errors = append(errors, err)
		return
	}

	asm, errors = compiler.CompileStage(code)
	if errors != nil {
		return
	}

	return
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var (
		reader   io.Reader
		compiler = mutan.NewCompiler(backend.NewEthereumBackend())
	)
	compiler.Debug = *Debug

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

	asm, err := Compile(compiler, reader)
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
