package main

import (
	"flag"
	"fmt"
	"github.com/obscuren/mutan"
	"os"
)

var Debug = flag.Bool("debug", false, "enable debug output")

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mutan [filename]")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	asm, err := mutan.Compile(file, *Debug)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(asm)
}
