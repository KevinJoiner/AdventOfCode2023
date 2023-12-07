package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read input %q: %v", inputFile, err)
	}
	lines := bytes.Split(input, []byte("\n"))
	timeParts := strings.Fields(string(lines[0]))[1:]
	distanceParts := strings.Fields(string(lines[1]))[1:]
	total := 1
	for i := range timeParts {
		raceTime, _ := strconv.Atoi(timeParts[i])
		record, _ := strconv.Atoi(distanceParts[i])
		winners := 0
		for j := 0; j <= raceTime; j++ {
			if distance(j, raceTime) > record {
				winners++
			}
		}
		total *= winners
	}
	fmt.Println("Product of good times is: ", total)
}

func distance(hold, raceTime int) int {
	return (raceTime - hold) * hold
}
