package cheatgoes

import (
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {

	myNumber := int(time.Now().Second() % 10)
	println(myNumber)
	switch myNumber {
	case 0, 1, 2, 3:
		println("smaller than 4")
	case 5:
		println("five!")
	case 4:
		println("Square number!")
		fallthrough
	case 7:
		println("4 or 7")
	case 6, 9:
		println("6 or 9")
	default:
		println("must be eight: ", myNumber)
	}

}
