package learninggo

import (
	"fmt"
	"testing"
	"time"
)

type Person struct {
	ID   int
	Name string
}

func SomeFunc(channel chan Person) {

	channel <- Person{ID: 1, Name: "Adam"}
	time.Sleep(time.Second * 5)
	channel <- Person{ID: 2, Name: "Eva"}

}

func TestChannels(t *testing.T) {
	myChannel := make(chan Person)

	go SomeFunc(myChannel)

	var myPerson Person
	myPerson = <-myChannel
	fmt.Println(myPerson)

	myPerson = <-myChannel
	fmt.Println(myPerson)
}
