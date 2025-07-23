// Package adder provides a simple example of an addition function.
//
// For a basic introduction to addition, see:
// https://www.mathsisfun.com/numbers/addition.html
package adder

// Add returns the sum of any number of integers.
func Add(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
