// Package adder provides a simple example of a generic addition function.
//
// For a basic introduction to addition, see:
// https://www.mathsisfun.com/numbers/addition.html
package adder

import "golang.org/x/exp/constraints"

// Number is any integer or float type.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of any number of values of type Number.
func Add[T Number](vals ...T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}
