package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
  scoreMap = map[string]int{"X": 1, "Y": 2, "Z": 3, "WIN": 6, "LOSE": 0, "DRAW": 3}
  moveBasedDesiredOutcome = map[string]map[string]string{"Z": {"A": "Y", "B": "Z", "C": "X"},
                                                        "Y": {"A": "X", "B": "Y", "C": "Z"},
                                                        "X": {"A": "Z", "B": "X", "C": "Y"}}
)

func getRoundResult(oponent string, me string) (string, error) {
  if (me == "X" && oponent == "A") || (me == "Y" && oponent == "B") || (me == "Z" && oponent == "C") {
    return "DRAW", nil
  } else if (me == "X" && oponent == "C") || (me == "Y" && oponent == "A") || (me == "Z" && oponent == "B") {
            return "WIN", nil 
  }
  return "LOSE", nil
}

func calculateRoundScore(oponent string, me string) (int, error) {
  roundResult, _ := getRoundResult(oponent, me)
  return scoreMap[me] + scoreMap[roundResult], nil
}

func resolveMyMove(oponent string, desiredOutcome string) (string, error) {
  return moveBasedDesiredOutcome[desiredOutcome][oponent], nil
}

func main() {
  totalScore := 0
  file, err := os.Open("input/input.txt")
  if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        splittedLine := strings.Split(line, " ")
        oponent, desiredOutcome := splittedLine[0], splittedLine[1]
        me, _ := resolveMyMove(oponent, desiredOutcome)
        roundScore, _ := calculateRoundScore(oponent, me)
        totalScore += roundScore
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
  
  fmt.Println(totalScore)
}
