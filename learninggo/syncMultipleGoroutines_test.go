// https://stackoverflow.com/questions/26927479/go-language-fatal-error-all-goroutines-are-asleep-deadlock
package learninggo

import (
	"fmt"
	"sync"
	"testing"
)

var wg2 sync.WaitGroup // 1

func routine2(i int) {
	defer wg2.Done() // 3
	fmt.Printf("routine %v finished\n", i)
}

func TestSyncMultipleGoroutines(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg2.Add(1)     // 2
		go routine2(i) // *
	}
	wg2.Wait() // 4
	fmt.Println("main finished")
}

// WaitGroup usage in order of execution.

//     Declaration of global variable. Making it global is the easiest way to make it visible to all functions and methods.
//     Increasing the counter. This must be done in main goroutine because there is no guarantee that newly started goroutine will execute before 4 due to memory model guarantees.
//     Decreasing the counter. This must be done at the exit of goroutine. Using deferred call, we make sure that it will be called whenever function ends no matter but no matter how it ends.
//     Waiting for the counter to reach 0. This must be done in main goroutine to prevent program exit.

// * The actual parameters are evaluated before starting new gouroutine. Thus it is needed to evaluate them explicitly before wg.Add(1) so the possibly panicking code would not leave increased counter.

// Use

// param := f(x)
// wg.Add(1)
// go g(param)

// instead of

// wg.Add(1)
// go g(f(x))
