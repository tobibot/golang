// https://stackoverflow.com/questions/26927479/go-language-fatal-error-all-goroutines-are-asleep-deadlock
package learninggo

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup // 1

func routine() {
	defer wg.Done() // 3
	fmt.Println("routine finished")
}

func TestSyncGoroutine(t *testing.T) {
	wg.Add(1)    // 2
	go routine() // *
	wg.Wait()    // 4
	fmt.Println("main finished")
}
