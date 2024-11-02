package node

func mapAll[T any, R any](a []T, f func(T) R) []R {
	b := make([]R, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}

func applyAll[T any](a []T, f func(T)) {
	for _, v := range a {
		f(v)
	}
}
