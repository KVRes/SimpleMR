package splitter

import "strings"

func TextSplit(data any, n int) []any {
	txt := data.(string)
	lines := strings.Split(txt, "\n")
	splt := splitArr(lines, n)
	return mapAll(splt, func(item []string) any { return strings.Join(item, "\n") })
}

func ArrSplit(data any, n int) []any {
	arr := data.([]any)
	return mapAll(splitArr(arr, n), func(item []any) any { return item })
}

func ArrTSplit[T any](data any, n int) []any {
	arr := data.([]T)
	return mapAll(splitArr[T](arr, n), func(item []T) any { return item })
}
