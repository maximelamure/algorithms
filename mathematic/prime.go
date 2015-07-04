package mathematic

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	for x := 2; x*x <= n; x = x + 1 {
		if n%x == 0 {
			return false
		}
	}
	return true
}
