package helpers

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func GetInput(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func GetLines(filename string) [][]byte {
	content := GetInput(filename)
	lines := bytes.Split(content, []byte("\n"))
	return lines
}

func ReadLines(filename string, lineChan chan string) {
	defer close(lineChan)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("reading standard input:", err)
	}
}
