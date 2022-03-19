package main

func low[T any](slice []T) T {
	return slice[0]
}

func high[T any](slice []T) T {
	return slice[len(slice)-1]
}

func insertAtIndex[T any](s *[]T, ix int, el T) {
	var buf T

	*s = append(*s, buf)

	copy((*s)[ix+1:], (*s)[ix:])
	(*s)[ix] = el
}
