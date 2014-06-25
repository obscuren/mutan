package frontend

import (
	_ "fmt"
	"regexp"
	"testing"
)

const src = `#define MY_MACRO 10
	#define OTHER 20
	big test = OTHER`

func TestFindDefines(t *testing.T) {
	_, macros := findDefines(src)
	if macros[0][0] != "MY_MACRO" {
		t.Error("expected MY_MACRO, got", macros[0][0])
	}

	if len(macros) != 2 {
		t.Error("Expected 2 macros, got", len(macros))
	}
}

func TestInnerDefines(t *testing.T) {
	src, _ := PreProcess(`
		#define HELLO 10
		#define HELLOWORLD 20

		newstring = HELLOWORLD
	`)

	if matched, _ := regexp.MatchString("newstring = 20", src); !matched {
		t.Error("Expected newstring to be 20")
	}
}

func TestPreProcess(t *testing.T) {
	PreProcess(src)
}
