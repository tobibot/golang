package learninggo

import (
	"fmt"
	"testing"
)

func TestFuncAsParam(t *testing.T) {
	printWithFunc(area, 2, 4)
	printWithFunc(sum, 2, 4)
}

func printWithFunc(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}

func area(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}
