package mutan

import (
	_ "fmt"
	"testing"
)

const src = `#define MY_MACRO 10
	#define OTHER 20
	big test = OTHER`

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
	PreProcess(src)
}
