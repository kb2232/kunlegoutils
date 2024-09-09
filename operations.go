package utils

import "math"

func Add[T int | float64 | string](a, b T) T {
	return a + b
}
func Sub[T int | float64](a, b T) T {
	return a - b
}
func Div[T int | float64](a, b T) T {
	return a / b
}
func Pow[T int | float64](a, b float64) T {
	return T(math.Pow(a, b))
}
