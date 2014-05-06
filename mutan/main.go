package main

import (
	"flag"
	"fmt"
	"github.com/obscuren/mutan"
	"os"
	"strings"
)

var Debug = flag.Bool("d", false, "enable debug output")
var DisableAssembler = flag.Bool("asm", false, "disable assembler stage")
var StrCode = flag.String("s", "", "code string")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var asm []interface{}
	var e []error
	if len(*StrCode) > 0 {
		asm, e = mutan.CompileStage(strings.NewReader(*StrCode), *Debug)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
	} else {
		file, err := os.Open(os.Args[len(os.Args)-1])
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		asm, e = mutan.CompileStage(file, *Debug)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
	}

	if *DisableAssembler {
		s := fmt.Sprintln(asm)
		fmt.Println(s[1 : len(s)-2])
	} else {
		bytes := mutan.Assemble(asm...)

		fmt.Printf("0x%x\n", bytes)
	}
}
