package types

//type MapFx[T any, R any] func(key string, value T) (string, []R)
//
//type ReduceFx[T any, R any] func(key string, values Iterator[T]) (string, R)
//
//type Iterator[T any] interface {
//	HasNext() bool
//	Next() (string, T)
//}

type Splitter func(data any, m int) []any

type Partitioner func(key string) int

type Intermediate map[string][]any
