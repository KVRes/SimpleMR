package splitter

func mapAll[T any, R any](a []T, f func(T) R) []R {
	b := make([]R, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}

func splitArr[T any](arr []T, n int) [][]T {
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return [][]T{arr}
	}
	if n >= len(arr) {
		return mapAll(arr, func(item T) []T { return []T{item} })
	}

	baseSize := len(arr) / n
	remainder := len(arr) % n

	result := make([][]T, 0, n)
	for i := 0; i < n; i++ {
		l := i * baseSize
		r := l + baseSize
		if i == n-1 {
			r += remainder
		}

		result = append(result, arr[l:r])
	}

	return result
}
