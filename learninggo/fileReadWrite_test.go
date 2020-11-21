package learninggo

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

const (
	fileName = "test.txt"
)

func TestReadFromFile(t *testing.T) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return
	}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		println(line)
	}

}

func TestWriteToFile(t *testing.T) {
	textToWrite := fmt.Sprintf("%v", time.Now())
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(fileName, options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, textToWrite)
	check(err)

}
