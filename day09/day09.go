package day09

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

type File struct {
	id    int
	size  int
	start int
	end   int
}

type FreeSpace struct {
	start int
	end   int
}

func Run() {
	input, err := utils.ReadInputRaw("day09/input.txt", 9)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

func part1(line string) int {
	files, freeSpace := getFragments(line)
	isFile := true
	currentFileIndex := 0
	currentSpaceIndex := 0

	checksum := 0
	for files[len(files)-1].end >= freeSpace[currentSpaceIndex].start {
		if isFile {
			for i := files[currentFileIndex].start; i <= files[currentFileIndex].end; i++ {
				checksum += files[currentFileIndex].id * i
			}

			currentFileIndex++
		} else {
			for i := freeSpace[currentSpaceIndex].start; i <= freeSpace[currentSpaceIndex].end; i++ {
				lastFile := files[len(files)-1]
				files = files[:len(files)-1]
				checksum += lastFile.id * i
				lastFile.end--
				lastFile.size--
				if lastFile.size > 0 {
					files = append(files, lastFile)
				}
			}

			currentSpaceIndex++
		}
		isFile = !isFile
	}

	for _, file := range files[currentFileIndex:] {
		for i := file.start; i <= file.end; i++ {
			checksum += files[currentFileIndex].id * i
		}
	}

	return checksum
}

func part2(line string) int {
	files, freeSpace := getFragments(line)
	checksum := 0

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		for j, space := range freeSpace {
			if space.start > file.start || file.end < space.start {
				checksum += (triangleNum(file.end) - triangleNum(file.start-1)) * file.id
				break
			}

			if file.size <= space.end-space.start+1 {
				file.start = space.start
				file.end = file.start + file.size - 1
				files[i] = file
				checksum += (triangleNum(file.end) - triangleNum(file.start-1)) * file.id

				space.start = file.end + 1
				if space.start > space.end {
					for k := j; k < len(freeSpace)-1; k++ {
						freeSpace[k] = freeSpace[k+1]
					}
				} else {
					freeSpace[j] = space
				}

				break
			}
		}
	}

	return checksum
}

func getFragments(line string) ([]File, []FreeSpace) {
	files := make([]File, 0)
	freeSpace := make([]FreeSpace, 0)
	isFile := true
	offset := 0

	for i, c := range line {
		size, _ := strconv.Atoi(string(c))
		if isFile {
			file := File{id: (i + 1) / 2, size: size, start: offset, end: offset + size - 1}
			files = append(files, file)
		} else {
			freeSpot := FreeSpace{start: offset, end: offset + size - 1}
			freeSpace = append(freeSpace, freeSpot)
		}

		offset += size
		isFile = !isFile
	}

	return files, freeSpace
}

func triangleNum(x int) int {
	return x * (x + 1) / 2
}
