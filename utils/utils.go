package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func linesGenerator(filePath string, ch chan string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	close(ch)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadFileToStringAndSplit(filePath string, sep string) []string {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(bytes), sep)
}
