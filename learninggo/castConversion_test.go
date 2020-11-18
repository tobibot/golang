package learninggo

import (
	"strconv"
	"testing"
)

func TestConversion(t *testing.T) {
	number := 1
	floatingPtNumber := 3.14

	resultAsFloat := float64(number) + floatingPtNumber
	println(resultAsFloat)

	resultAsInt := number + int(floatingPtNumber)
	if resultAsInt != 4 {
		t.Errorf("you need to invest more time learning type conversions")
	}

	text := "NumberAsString"
	numberFromTextWillBe0, _ := strconv.Atoi(text)
	println(numberFromTextWillBe0)

	numberFromText1909, _ := strconv.Atoi("1909")

	if numberFromText1909 != 1909 {
		t.Errorf("1909 didn't work out")
	}

}

func TestCasting(t *testing.T) {

}
