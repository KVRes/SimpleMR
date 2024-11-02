package splitter

import "strings"

func TextSplit(data any, n int) []any {
	txt := data.(string)
	lines := strings.Split(txt, "\n")
	splt := splitArr(lines, n)
	return mapAll(splt, func(item []string) any { return strings.Join(item, "\n") })
}
