package learninggo_test

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestBase64Decode(t *testing.T) {

	str := "SGVsbG8gV29ybGQ="
	want := "Hello World"
	got, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}

	if string(got) != want {
		t.Errorf("wanted %s\n, but got %s ", want, string(got))		
	}
	
	fmt.Printf("%q\n", got)
	
}

func TestBase64Encode(t *testing.T) {

	want := "SGVsbG8gV29ybGQ="
	str := "Hello World"

	got := base64.StdEncoding.EncodeToString([]byte(str))

	if string(got) != want {
		t.Errorf("wanted %s\n, but got %s ", want, string(got))		
	}
	fmt.Printf("%q\n", got)
}
