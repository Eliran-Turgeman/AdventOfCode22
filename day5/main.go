package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type moveInstruction struct {
	from int
	to   int
	nb   int
}

const (
	crateStart       string = "["
	instructionStart string = "m"
)

func parseCratesAndMoves(fileName string) ([][]string, []moveInstruction) {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var crates [][]string
	moves := make([]moveInstruction, 0)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		st := strings.TrimSpace(currentLine)

		if len(st) == 0 {
			continue
		}

		switch st[0] {
		case crateStart[0]:
			nbStacks := (len(currentLine) + 1) / 4
			for i := 0; i < nbStacks; i++ {
				if len(crates) == 0 {
					crates = make([][]string, nbStacks)
				}
				cratePosition := (i * 4) + 1
				if crate := string(currentLine[cratePosition]); crate != " " {
					crates[i] = append(crates[i], crate)
				}
			}
		case instructionStart[0]:
			var (
				from int
				to   int
				nb   int
			)

			fmt.Sscanf(currentLine, "move %d from %d to %d", &nb, &from, &to)
			moves = append(moves, moveInstruction{from: from, to: to, nb: nb})
		default:
		}
	}
	readFile.Close()

	return crates, moves
}

func applyMoves(crates [][]string, moves []moveInstruction) (string) {
  for _, instr := range moves {
		for i := 0; i < instr.nb; i++ {
			crateFrom := crates[instr.from-1][1:]
			crateTo := append(make([]string, 0), crates[instr.from-1][0])
			crateTo = append(crateTo, crates[instr.to-1]...)

			crates[instr.to-1] = crateTo
			crates[instr.from-1] = crateFrom
		}
	}

	tops := make([]string, 0)
	for _, crate := range crates {
		tops = append(tops, crate[0])
	}

	return strings.Join(tops, "")
}

func applyMovesP2(crates [][]string, moves []moveInstruction) (string) {
  for _, instr := range moves {
		crateFrom := crates[instr.from-1][instr.nb:]
		crateTo := append(make([]string, 0), crates[instr.from-1][0:instr.nb]...)
		crateTo = append(crateTo, crates[instr.to-1]...)

		crates[instr.to-1] = crateTo
		crates[instr.from-1] = crateFrom
	}

	tops := make([]string, 0)
	for _, crate := range crates {
		tops = append(tops, crate[0])
	}

	return strings.Join(tops, "")
}

func main() {
  crates, moves := parseCratesAndMoves("input/input.txt")
  topCrates := applyMovesP2(crates, moves)
  fmt.Println(topCrates)
}
