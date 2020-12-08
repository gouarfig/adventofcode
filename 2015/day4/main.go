package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	getNumberFor([]byte("abcdef"), "00000")
	getNumberFor([]byte("pqrstuv"), "00000")
	getNumberFor([]byte("yzbqklnj"), "00000")
	getNumberFor([]byte("yzbqklnj"), "000000")
}

func getNumberFor(key []byte, prefix string) uint64 {
	var number uint64 = 0
	for {
		if isValidNumber(key, number, prefix) {
			fmt.Printf("md5 starting with %q for key %q: number = %d\n", prefix, string(key), number)
			break
		}
		number++
	}
	return number
}

func isValidNumber(key []byte, number uint64, prefix string) bool {
	data := strconv.AppendUint(key, number, 10)
	h := md5.Sum(data)
	temp := fmt.Sprintf("%x", h)
	// fmt.Printf("%s: %x\n", string(data), h)
	return strings.HasPrefix(temp, prefix)
}
