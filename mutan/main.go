package main

import (
	"flag"
	"fmt"
	"github.com/obscuren/mutan"
	"os"
)

var Debug = flag.Bool("d", false, "enable debug output")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	asm, e := mutan.Compile(file, *Debug)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	fmt.Println(asm)
}
