package main

import "fmt"

func main() {
	ints := map[string]int64{
		"one": 23,
		"two": 43,
	}

	floats := map[string]float64{
		"one": 45.44,
		"two": 23.43,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("Generic Sum: %v and %v\n", Sum(ints), Sum(floats))

	employees := Stack[string]{}

	employees.Push("John")
	employees.Push("Daniel")
	employees.Push("Deborah")

	fmt.Println("Employees stack: ", employees)
	fmt.Println("Employees stack size: ", employees.Size())

	pe, _ := employees.Pop()

	fmt.Println("Employee popped: ", pe)
	fmt.Println("Employees stack size: ", employees.Size())
}

func SumInts(m map[string]int64) int64 {
	var sum int64

	for _, v := range m {
		sum += v
	}

	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64

	for _, v := range m {
		sum += v
	}

	return sum
}

type Number interface {
	int64 | float64
}

// Generic function to handle both use cases
func Sum[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

// Stack data structure with generics

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(e T) {
	s.data = append(s.data, e)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T

		return zero, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return value, true
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}
