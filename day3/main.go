package main

import (
	"AdventOfCode22/utils"
	"fmt"
	"strings"
	"unicode"
)

func getDuplicateItemPriority(line string) (int, error) {
	lineMiddle := len(line) / 2
	firstCompartment := line[:lineMiddle]
	secondCompartment := line[lineMiddle:]

	for _, ch := range firstCompartment {
		_ch := string(ch)
		if strings.Contains(secondCompartment, _ch) {
			if unicode.IsUpper(ch) {
				return int(ch) - 38, nil
			}
			return int(ch) - 96, nil
		}
	}
	return 0, nil
}

func getGroupBadgePriority(linesArray []string) (int, error) {
	firstLine := linesArray[0]
	secondLine := linesArray[1]
	thirdLine := linesArray[2]

	for _, ch := range firstLine {
		_ch := string(ch)
		if (strings.Contains(secondLine, _ch)) && (strings.Contains(thirdLine, _ch)) {
			if unicode.IsUpper(ch) {
				return int(ch) - 38, nil
			}
			return int(ch) - 96, nil
		}
	}
	return 0, nil
}

func main() {
	lines := utils.ReadFileToStringAndSplit("input/input.txt", "\n")
	prioritiesSum := 0

	var groupLines []string

	for _, line := range lines {
		groupLines = append(groupLines, line)
		if len(groupLines) == 3 {
			priority, _ := getGroupBadgePriority(groupLines)
			prioritiesSum += priority
			groupLines = nil
		}
	}
	fmt.Println(prioritiesSum)
}
