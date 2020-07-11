package tests

func F1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return F1(n-1) + F1(n-2)
}

func F2(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 2
	}

	return F2(n-1) + F2(n-2)
}

func S1(s string) int {
	if s == "" {
		return 0
	}

	n := 1
	for range s {
		n++
	}
	return n
}

func S2(s string) int {
	return len(s)
}
