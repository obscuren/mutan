package frontend

import (
	"regexp"
	"strings"
)

func findConsts(source string) (string, [][]string) {
	const reg = "\\s?(const [a-zA-Z0-9]+\\s?=\\s?.*)\\n"
	re := regexp.MustCompile(reg)

	var defs [][]string
	return re.ReplaceAllStringFunc(source, func(f string) string {
		fields := strings.Fields(f)

		var key, value string
		if len(fields) > 1 {
			if len(fields) == 2 {
				pair := strings.Split(fields[1], "=")
				if len(pair) != 2 {
					return "\n"
				}
				key = pair[0]
				value = pair[1]
			} else {
				key = strings.TrimRight(fields[1], "=")
				value = fields[len(fields)-1]
			}

			defs = append(defs, []string{key, value})
		}

		return "\n"
	}), defs
}

func PreProcess(s string) (source string, err error) {
	// Pre process constants
	var macros [][]string
	source, macros = findConsts(s)
	for _, macro := range macros {
		source = regexp.MustCompile("\\b"+macro[0]+"\\b").ReplaceAllString(source, macro[1])
	}

	return
}
