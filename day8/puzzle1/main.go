package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[0 : len(lines)-1]
	seq := lines[0]
	lines = lines[2:]

	// create adjMap
	adjMap := map[string][2]string{}
	for _, line := range lines {
		fields := bytes.Fields(line)
		parent := strings.TrimFunc(string(fields[0]), IsNotLetter)
		left := strings.TrimFunc(string(fields[2]), IsNotLetter)
		right := strings.TrimFunc(string(fields[3]), IsNotLetter)
		adjMap[parent] = [2]string{left, right}
	}

	total := 0
	idx := new(int)
	curr := "AAA"
	for curr != "ZZZ" {
		curr = adjMap[curr][LorR(seq, idx)]
		total++
	}

	fmt.Println("Total steps: ", total)
}

func LorR(seq []byte, idx *int) uint8 {
	if *idx == len(seq) {
		*idx = 0
	}
	defer func() { *idx++ }()
	if seq[*idx] == 'L' {
		return 0
	}
	return 1
}

func IsNotLetter(r rune) bool {
	return !unicode.IsLetter(r)
}
