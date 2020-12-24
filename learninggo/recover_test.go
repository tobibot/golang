package learninggo

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {

	defer recoverFromPanic()

	fmt.Println("doing something")
	panic("ooops")

	fmt.Println("Finisch")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
