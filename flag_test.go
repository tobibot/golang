package learninggo

import "testing"

type Flags int

const (
	Red   Flags = 0 // 000
	Blue        = 1 // 001
	Green       = 2 // 010
	Pink        = 3 // 011
	Gold        = 4 // 100

)

func TestFlags(t *testing.T) {
	println(Red | Blue)
	println(Red | Blue | Gold)

	println(Green & Pink)
	println(Blue & Pink)

}
