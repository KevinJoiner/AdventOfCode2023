package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	lines := bytes.Split(input, []byte("\n"))
	lines = lines[0 : len(lines)-1]
	total := 0
	for _, line := range lines {
		fields := bytes.Fields(line)
		nums := make([]int, len(fields))
		for i, field := range fields {
			nums[i], err = strconv.Atoi(string(field))
			if err != nil {
				panic(err)
			}
		}
		total += getNext(nums)
	}
	fmt.Println("Total of previous values: ", total)
}

func getNext(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	same := true
	first := nums[0]
	for i := 0; i < n-1; i++ {
		nums[i] = nums[i+1] - nums[i]
		if i > 0 && same && nums[i] != nums[i-1] {
			same = false
		}
	}
	next := nums[0]
	if !same {
		next = getNext(nums[:n-1])
	}
	return first - next
}
