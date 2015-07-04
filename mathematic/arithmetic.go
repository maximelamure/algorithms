package mathematic

import "errors"

// SubstractAddOnly implements a substaction with only the Add operator
func SubstractAddOnly(a, b int) int {
	return a + Min(b)
}

func Min(b int) int {
	result := 0
	d := -1
	if b < 0 {
		d = 1
	}
	for ; b != 0; b += d {
		result += d
	}
	return result
}

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// MultiplyAddOnly implements a multiplication with only the Add operator
func MultiplyAddOnly(a, b int) int {
	if a < b {
		MultiplyAddOnly(b, a)
	}

	result := 0
	for x := 0; x < Abs(b); x++ {
		result += a
	}

	if b < 0 {
		result = Min(result)
	}
	return result
}

// DivideAddOnly implements a division with only the Add operator
func DivideAddOnly(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by 0")
	}

	result := 0
	for x := Abs(a); x >= Abs(b); x = SubstractAddOnly(x, Abs(b)) {
		result++
	}
	if (b < 0 && a > 0) || (b > 0 && a < 0) {
		result = Min(result)
	}

	return result, nil
}

func Pow(a, b int) int {
	if a == 0 && b != 0 {
		return 0
	}
	if b == 0 {
		return 1
	}

	mid := Pow(a, (b-1)/2)
	mid *= mid
	if b%2 != 0 {
		mid *= a
	}

	return mid
}

func Multiply(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < b {
		return Multiply(b, a)
	}

	if b == 1 {
		return a
	}
	if b == 2 {
		return a + a
	}
	temp := Multiply(a, (b-1)/2)
	temp += temp + a
	return temp

}
