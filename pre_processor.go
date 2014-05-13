package mutan

import (
	"regexp"
	"strings"
)

func findDefines(source string) (string, [][]string) {
	const reg = "\\s?(#define\\s.*)\\n"
	re := regexp.MustCompile(reg)

	var defs [][]string
	return re.ReplaceAllStringFunc(source, func(f string) string {
		splitted := strings.SplitN(strings.TrimSpace(f)[8:], " ", 2)
		if len(splitted) == 1 {
			splitted = append(splitted, " ")
		}

		defs = append(defs, splitted)

		return "\n"
	}), defs
}

func PreProcess(s string) (source string, err error) {
	// Pre process defines
	var macros [][]string
	source, macros = findDefines(s)
	for _, macro := range macros {
		source = regexp.MustCompile("\\b"+macro[0]+"\\b").ReplaceAllString(source, macro[1])
	}

	return
}
