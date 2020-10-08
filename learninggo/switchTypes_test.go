package learninggo

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitchWithTypes(t *testing.T) {
	DoSth(1)
	DoSth(1.11)
	DoSth(int16(3))
	DoSth(true)
	DoSth("true")
}

func DoSth(myType interface{}) {

	switch tpye := myType.(type) {
	case string:
		time.Sleep(time.Millisecond * 10)
		println("string type")
	case float64, float32, int:
		println("something with numbers")
	default:
		fmt.Printf("%T\n", tpye)
	}

}
