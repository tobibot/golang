package learninggo

import (
	"fmt"
	"runtime"
	"testing"
)

func TestOsCheck(t *testing.T) {
	fmt.Print("Go runs on ")
	switch {
	case runtime.GOOS == "darwin":
		fmt.Println("OS X.")
	case runtime.GOOS == "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", runtime.GOOS)
	}
}
