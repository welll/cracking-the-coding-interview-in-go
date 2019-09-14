/*

Factorial

The factorial is a quantity defined for all non-negative integers:
- for zero, the factorial is defined as 1
- for all other positive integers, factorial is defined as the product
  of the integer and all positive integers less than it.

Examples:
0 =>                       =   1 (by definition)
1 =>                    1  =   1 (there are no positive integers less than 1)
2 =>                2 * 1  =   2
3 =>            3 * 2 * 1  =   6
4 =>        4 * 3 * 2 * 1  =  24
5 =>    5 * 4 * 3 * 2 * 1  = 120
6 => 6 * 5 * 4 * 3 * 2 * 1 = 720

*/

package misc

import (
	"testing"
)

// The factorial function takes an integer and returns the product of the
// integer and all integers less than it
type FactorialFunc func(int) int

// map of inputs to expected results for factorial
var tests = [][]int{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{5, 120},
	{6, 720},
}

// factorial implementations to test
var funcs = map[string]FactorialFunc{
	"factorial":          factorial,
	"factorialRecursive": factorialRecursive,
}

func TestFactorial(t *testing.T) {
	for name, f := range funcs {
		t.Run(name, func(t *testing.T) {
			for _, test := range tests {
				if actual := f(test[0]); actual != test[1] {
					t.Fatalf("\n[%s]: for n=%d, expected: %d, actual: %d",
						name, test[0], test[1], actual)
				}
			}
		})
	}
}

func factorial(n int) int {
	x := 1
	for i := n; i > 1; i-- {
		x *= i
	}
	return x
}

func factorialRecursive(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	default:
		return n * factorialRecursive(n-1)
	}
}
