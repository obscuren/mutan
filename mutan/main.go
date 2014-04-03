package main

import (
	"flag"
	"fmt"
	"github.com/obscuren/mutan"
	"os"
)

var Debug = flag.Bool("d", false, "enable debug output")
var Fn = flag.String("f", "", "file name")

func main() {
	flag.Parse()

	file, err := os.Open(*Fn)
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
