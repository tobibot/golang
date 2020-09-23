package learninggo

import "testing"

func TestRemoveFromSlice(t *testing.T) {
	var mySlice []int = []int{1, 2, 3, 4, 5}

	sliceWithoutMiddle := append(mySlice[:2], mySlice[3:]...)

	for _, item := range sliceWithoutMiddle {
		println(item)
	}
	// println(sliceWithoutMiddle)
}
