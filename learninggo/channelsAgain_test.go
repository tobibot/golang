package learninggo

import (
	"fmt"
	"testing"
)

func TestChannelsAgain(t *testing.T) {
	myChannel := make(chan Person)

	go SomeFunc(myChannel)

	var myPerson Person
	myPerson = <-myChannel
	fmt.Println(myPerson)

	<-myChannel
}
