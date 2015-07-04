package mathematic_test

import (
	"log"
	"testing"
	"time"

	"github.com/maximelamure/algorithms/mathematic"
	"github.com/maximelamure/api/common"
)

var helper common.Test = common.Test{}

func TestSub(t *testing.T) {
	helper.Assert(t, mathematic.SubstractAddOnly(5, 3) == 2, "Substract 5 3 should be 2")
	helper.Assert(t, mathematic.SubstractAddOnly(5, -3) == 8, "Substract 5 -3 should be 8")
	helper.Assert(t, mathematic.SubstractAddOnly(-5, -3) == -2, "Substract -5 -3 should be -2")
}

func TestMultiplication(t *testing.T) {
	helper.Assert(t, mathematic.MultiplyAddOnly(5, 3) == 15, "Multiply 5 3 should be 15")
	helper.Assert(t, mathematic.MultiplyAddOnly(5, -3) == -15, "Multiply 5 -3 should be -15")
	helper.Assert(t, mathematic.MultiplyAddOnly(-5, -3) == 15, "Multiply -5 -3 should be 15")
	log.Print(mathematic.Multiply(5, 1))
	helper.Assert(t, mathematic.Multiply(5, 1) == 5, "Multiply 5 1 should be 5")
	log.Print(mathematic.Multiply(5, 2))
	helper.Assert(t, mathematic.Multiply(5, 2) == 10, "Multiply 5 2 should be 10")
	log.Print(mathematic.Multiply(5, 3))
	helper.Assert(t, mathematic.Multiply(5, 3) == 15, "Multiply 5 3 should be 15")
}

func TestDivision(t *testing.T) {
	val, _ := mathematic.DivideAddOnly(5, 3)
	helper.Assert(t, val == 1, "Division 5 3 should be 1")
	val, _ = mathematic.DivideAddOnly(-5, -3)
	helper.Assert(t, val == 1, "Division -5 -3 should be 1")
	val, _ = mathematic.DivideAddOnly(5, -3)
	helper.Assert(t, val == -1, "Division 5 -3 should be -1")
	val, _ = mathematic.DivideAddOnly(180, 5)
	helper.Assert(t, val == 36, "Division 180 5 should be 36")
}

func TestFibonacci(t *testing.T) {
	start := time.Now()
	helper.Assert(t, mathematic.Fibonacci(10) == 55, "Fibonacci 10 should be 55")
	log.Printf("Fibonacci tooks %v to run.\n", time.Now().Sub(start))
	start = time.Now()
	helper.Assert(t, mathematic.FibonacciDyamic(10) == 55, "Fibonacci 10 should be 55")
	log.Printf("FibonacciDyamic tooks %v to run.\n", time.Now().Sub(start))
}

func TestIsPrime(t *testing.T) {
	helper.Assert(t, !mathematic.IsPrime(358), "358 is not prime")
	helper.Assert(t, mathematic.IsPrime(997), "997 is Prime")
}

func TestPow(t *testing.T) {
	helper.Assert(t, mathematic.Pow(2, 3) == 8, "Pow(2,3) should be 8")
	helper.Assert(t, mathematic.Pow(0, 0) == 1, "Pow(0,0) should be 1")
	helper.Assert(t, mathematic.Pow(5, 0) == 1, "Pow(5,0) should be 1")
	helper.Assert(t, mathematic.Pow(0, 5) == 0, "Pow(0,5) should be 0")
}
