package mutan

import (
	"fmt"
	"testing"
)

const src = `#define MY_MACRO 10
	#define OTHER 20
	#define ADD_IT 2 + 2
	big test = ADD_IT`

func TestFindDefiness(t *testing.T) {
	_, macros := findDefines(src)
	if macros[0][0] != "MY_MACRO" {
		t.Error("expected MY_MACRO, got", macros[0][0])
	}

	if len(macros) != 2 {
		t.Error("Expected 2 macros, got", len(macros))
	}
}

func TestPreProcess(t *testing.T) {
	source, _ := PreProcess(src)
	fmt.Println("new source", source)
}
