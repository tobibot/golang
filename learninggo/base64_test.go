package learninggo_test

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestBase64(t *testing.T) {
	str := "SGVsbG8gV29ybGQ="
	output, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", output)
}
