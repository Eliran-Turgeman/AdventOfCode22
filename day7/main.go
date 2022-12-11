package main

import (
	"fmt"
	"AdventOfCode22/utils"
  "strings"
  "strconv"
)

const (
  minDirSize = 100000
  fileSystemSize = 70000000
  requiredSizeForUpdate = 30000000
)


type file struct {
  name string
  size int
}

type directory struct {
  name string
  files []file
  subdirs []directory
  parent *directory
}

var dirsCandidatesForDeletion []directory = []directory{}

func (d directory) filteredSum() int {
	acc := 0
	size := d.getDirSize()
	if size <= minDirSize {
		acc += size
	}

	for _, dir := range d.subdirs {
		acc += dir.filteredSum()
	}

	return acc
}

func (d directory) getDirSize() int {
	acc := 0
	for _, dir := range d.subdirs {
		acc += dir.getDirSize()
	}
	for _, file := range d.files {
		acc += file.size
	}
	return acc
}


func getFS(lines []string) (directory) {  
  root := directory{name: "/"}
	pointer := &root
  
  for _, line := range(lines){
    if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd ") {
				moveToDir := strings.TrimPrefix(line, "$ cd ")
				if moveToDir == ".." {
					pointer = pointer.parent
				} else {
					for i, dir := range pointer.subdirs {
						if dir.name == moveToDir {
							pointer = &pointer.subdirs[i]
							break
						}
					}
				}
			} else {
				continue
			}
		} else {
			line := strings.Split(line, " ")
			size, filename := line[0], line[1]
			if size == "dir" {
				pointer.subdirs = append(pointer.subdirs, directory{
					name:   filename,
					parent: pointer,
				})
			} else {
        size, _ := strconv.Atoi(size)
				pointer.files = append(pointer.files, file{
					name: filename,
					size: size,
				})
			}
		}
  }
  return root
}

func fillDirsToRemove(dir directory, needed int) {
  if dir.getDirSize() >= needed {
		dirsCandidatesForDeletion = append(dirsCandidatesForDeletion, dir)
	}
	for _, subdir := range dir.subdirs {
		fillDirsToRemove(subdir, needed)
	}
}


func main() {
  lines := utils.ReadFileToStringAndSplit("input/input.txt", "\n")
  root := getFS(lines)
  sum := root.filteredSum()

  unused := (fileSystemSize - root.getDirSize())
	needed := requiredSizeForUpdate - unused

  fillDirsToRemove(root, needed)
  smallestDir := dirsCandidatesForDeletion[0].getDirSize()

	for _, dir := range dirsCandidatesForDeletion {
		if dir.getDirSize() < smallestDir {
			smallestDir = dir.getDirSize()
		}
	}
  
  fmt.Println("Part 1: ", sum, "Part 2: ", smallestDir)
}
