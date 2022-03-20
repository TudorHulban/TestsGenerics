package main

func low[T any](s []T) T {
	if len(s) == 0 {
		var t T

		return t
	}

	return s[0]
}

func high[T any](s []T) T {
	if len(s) == 0 {
		var t T

		return t
	}

	return s[len(s)-1]
}

func insertAtIndex[T any](s *[]T, ix int, el T) {
	if len(*s) == 0 {
		*s = append(*s, el)
		return
	}

	var buf T

	*s = append(*s, buf)

	copy((*s)[ix+1:], (*s)[ix:])
	(*s)[ix] = el
}

func forEach[T any](s []T, f func(int, T)) {
	for i, t := range s {
		f(i, t)
	}
}
