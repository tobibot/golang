package learninggo

import "testing"

func TestDefer(t *testing.T) {
	defer DeferFunc(t)
	panic("get me out of here!")

}

func DeferFunc(t *testing.T) {
	println("i was deferred")
	t.Skip()
}
