package learninggo

import "testing"

func add(numbersToAdd ...int) int {
	sum := 0

	for _, number := range numbersToAdd {
		sum += number
	}
	return sum
}

func TestVariadicFunc(t *testing.T) {
	println(add(2, 3))
	println(add(1, 2, 3))
	mySlice := []int{2, 3, 4}
	println(add(mySlice...))
}
