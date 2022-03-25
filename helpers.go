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

func removeDups[T comparable](s []T) []T {
	var res []T
	buf := make(map[T]bool)

	for _, value := range s {
		if _, exists := buf[value]; exists {
			continue
		}

		buf[value] = true
		res = append(res, value)
	}

	return res
}

func occurences[T comparable](s []T) map[T]int {
	res := make(map[T]int)

	for _, value := range s {
		if _, exists := res[value]; exists {
			res[value]++
			continue
		}

		res[value] = 1
	}

	return res
}
