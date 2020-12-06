package helpers

import (
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
