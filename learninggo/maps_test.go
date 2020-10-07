package learninggo

import (
	"fmt"
	"testing"
)

func printMapItems(someMap map[int]string) {
	for key, val := range someMap {
		fmt.Printf("key: %d, val: %s\n", key, val)
	}
}

func TestMap(t *testing.T) {
	var mySimpleMap map[int]string
	mySimpleMap = make(map[int]string)

	mySimpleMap[8] = "go"
	mySimpleMap[42] = "go beyond"
	mySimpleMap[8] = "don't go"

	printMapItems(mySimpleMap)

	anotherMap := map[int]string{1: "test"}
	printMapItems(anotherMap)

}
