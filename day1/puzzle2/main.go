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
	root := &Trie{}
	Add(root, "one", 1)
	Add(root, "two", 2)
	Add(root, "three", 3)
	Add(root, "four", 4)
	Add(root, "five", 5)
	Add(root, "six", 6)
	Add(root, "seven", 7)
	Add(root, "eight", 8)
	Add(root, "nine", 9)
	total := 0
	for _, line := range bytes.Split(input, []byte("\n")) {
		first := -1
		last := -1
		ptr := 0
		for i, b := range line {
			if '0' <= b && b <= '9' {
				if first == -1 {
					first = int(b - '0')
				}
				last = int(b - '0')
				ptr = i + 1
			} else {
				digit, ok := Find(root, line[ptr:i+1])
				for !ok {
					if ptr > i {
						ptr = i + 1
						break
					}
					ptr++
					digit, ok = Find(root, line[ptr:i+1])
				}
				if digit != 0 {
					if first == -1 {
						first = digit
					}
					last = digit
				}
			}
		}
		if first == -1 {
			first = 0
			last = 0
		}
		total += first * 10
		total += last
	}
	fmt.Printf("Total value is : %d\n", total)
}

type Trie struct {
	Children [26]*Trie
	Num      int
}

func Find(root *Trie, word []byte) (int, bool) {
	if root == nil {
		return 0, false
	}
	curr := root
	for _, b := range word {
		curr = curr.Children[b-'a']
		if curr == nil {
			return 0, false
		}
	}
	if curr.Num != 0 {
		return curr.Num, true
	}
	// not a digit but could be
	return 0, true
}

func Add(root *Trie, word string, val int) {
	curr := root
	for i := range word {
		b := word[i]
		next := curr.Children[b-'a']
		if next == nil {
			next = &Trie{}
			curr.Children[b-'a'] = next
		}
		curr = next
	}
	curr.Num = val
}
