package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	total := 0
	for _, line := range bytes.Split(input, []byte("\n")) {
		first := byte('0')
		last := byte('0')
		for _, b := range line {
			if '0' <= b && b <= '9' {
				first = b
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			b := line[i]
			if '0' <= b && b <= '9' {
				last = b
				break
			}
		}
		total += int(first-'0') * 10
		total += int(last - '0')
	}
	fmt.Printf("Total value is : %d\n", total)
}
