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
	timeStr := strings.Join(timeParts, "")
	distanceParts := strings.Fields(string(lines[1]))[1:]
	recordStr := strings.Join(distanceParts, "")
	raceTime, _ := strconv.Atoi(timeStr)
	record, _ := strconv.Atoi(recordStr)
	start := 0
	stop := raceTime
	for j := 0; j <= raceTime; j++ {
		if distance(j, raceTime) > record {
			start = j
			break
		}
	}
	if start == stop {
		// went through the list either we have 1 or 0 winners
		total := 0
		if distance(start, raceTime) > record {
			total = 1
		}
		fmt.Println("Total of good times is: ", total)
		return
	}
	for j := raceTime; j >= 0; j-- {
		if distance(j, raceTime) > record {
			stop = j
			break
		}
	}
	fmt.Println("Total of good times is: ", stop-start+1)
}

func distance(hold, raceTime int) int {
	return (raceTime - hold) * hold
}
