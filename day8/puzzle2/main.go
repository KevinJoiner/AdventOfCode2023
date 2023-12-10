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

	currs := []string{}
	// create adjMap
	adjMap := map[string][2]string{}
	for _, line := range lines {
		fields := bytes.Fields(line)
		parent := strings.TrimFunc(string(fields[0]), IsNotLetter)
		left := strings.TrimFunc(string(fields[2]), IsNotLetter)
		right := strings.TrimFunc(string(fields[3]), IsNotLetter)
		adjMap[parent] = [2]string{left, right}
		if parent[2] == 'A' {
			currs = append(currs, parent)
		}
	}

	total := int64(0)
	idx := new(int)
	finished := false
	tots := make([]int64, len(currs))
	seens := make([]map[string]int64, len(currs))
	for i := range currs {
		seens[i] = map[string]int64{}
	}
	finds := 0
	fmt.Println(currs)
	for !finished {
		direction := LorR(seq, idx)
		for i, curr := range currs {
			next := adjMap[curr][direction]

			if seens[i][next] != 0 && next[2] == 'Z' && tots[i] == 0 {
				fmt.Println(tots)
				tots[i] = total - seens[i][next]
				finds++
			}
			if seens[i][next] == 0 {
				seens[i][next] = total
			}

			currs[i] = next
		}
		finished = finds == len(currs)
		total++
	}
	fmt.Println(tots, finds)
	fmt.Println("Total steps: ", LCM(tots[0], tots[1], tots[2:]...))
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

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
