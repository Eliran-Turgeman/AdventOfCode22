package main

import (
	"fmt"
	"main/utils"
)

const (
  markerLength = 14
)

func allCharsUnique(s string) (bool) {
    var a [256]bool
    for _, ascii := range s {

        if a[ascii] {
            return false
        }

        a[ascii] = true
    }
    return true
}

func findMarkerHighIndex(line string) (int) {
  low, high := 0, markerLength

  for i := 1; i < len(line) - markerLength; i++ {
    if allCharsUnique(line[low:high]){
      return high
    }
    low += 1; high += 1
  }
  return -1
}


func main() {
  lines := utils.ReadFileToStringAndSplit("input/input.txt", " ")
  markerIndex := findMarkerHighIndex(lines[0])
  fmt.Println(markerIndex)
}
