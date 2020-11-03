package learninggo

import (
	"fmt"
	"testing"
	"time"
)

func Test_Select(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(sleep int) {
		time.Sleep(time.Duration(sleep) * time.Second)
		c1 <- "one"
	}(3)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
