package learninggo

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFileCopy(t *testing.T) {
	src := "test.txt"
	dst := "test_copy.txt"

	writeResult, err := CopyFile(src, dst)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Write result", writeResult)

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		t.Error(err)
	}

	err = os.Remove(dst)
	if err != nil {
		t.Error(err)
	}
}

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
