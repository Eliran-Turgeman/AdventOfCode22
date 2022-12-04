package main

import (
	"AdventOfCode22/utils"
	"fmt"
	"strconv"
	"strings"
)

func isFullyOverlappingPair(line string) bool {
	firstElfLow, firstElfHigh, secondElfLow, secondElfHigh := parsePairIntoElfRanges(line)

	if ((firstElfLow <= secondElfLow) && (secondElfLow <= secondElfHigh) && (secondElfHigh <= firstElfHigh)) ||
		((secondElfLow <= firstElfLow) && (firstElfLow <= firstElfHigh) && (firstElfHigh <= secondElfHigh)) {
		return true
	}
	return false
}

func isOverlappingPair(line string) bool {
	firstElfLow, firstElfHigh, secondElfLow, secondElfHigh := parsePairIntoElfRanges(line)

	if (firstElfLow <= secondElfHigh) && (firstElfHigh >= secondElfLow) {
		return true
	}
	return false
}

func parsePairIntoElfRanges(line string) (int, int, int, int) {
	pair := strings.Split(line, ",")
	firstElf := pair[0]
	secondElf := pair[1]

	firstElfRanges := strings.Split(firstElf, "-")
	secondElfRanges := strings.Split(secondElf, "-")

	firstElfLow, _ := strconv.Atoi(firstElfRanges[0])
	firstElfHigh, _ := strconv.Atoi(firstElfRanges[1])
	secondElfLow, _ := strconv.Atoi(secondElfRanges[0])
	secondElfHigh, _ := strconv.Atoi(secondElfRanges[1])

	return firstElfLow, firstElfHigh, secondElfLow, secondElfHigh
}

func main() {
	lines := utils.ReadFileToStringAndSplit("input/input.txt", "\n")
	fullyOverlappingPairs := 0

	for _, line := range lines {
		if isOverlappingPair(line) {
			fullyOverlappingPairs += 1
		}
	}
	fmt.Println(fullyOverlappingPairs)
}
