package day04

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day04/input.txt", 4)
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

func part1(lines []string) int {
	xmasOccurrences := 0
	h := len(lines)
	w := len(lines[0])

	for y, line := range lines {
		for x := range line {
			if x <= w-4 {
				// look forward
				if line[x] == 'X' && line[x+1] == 'M' && line[x+2] == 'A' && line[x+3] == 'S' {
					xmasOccurrences++
				}

				if y <= h-4 {
					// look down-left
					if lines[y][x] == 'X' && lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
						xmasOccurrences++
					}
				}

				if y >= 3 {
					// look up-left
					if lines[y][x] == 'X' && lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
						xmasOccurrences++
					}
				}
			}

			if x >= 3 {
				// look backwards
				if line[x] == 'X' && line[x-1] == 'M' && line[x-2] == 'A' && line[x-3] == 'S' {
					xmasOccurrences++
				}

				if y <= h-4 {
					// look down-right
					if lines[y][x] == 'X' && lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S' {
						xmasOccurrences++
					}
				}

				if y >= 3 {
					// look up-right
					if lines[y][x] == 'X' && lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S' {
						xmasOccurrences++
					}
				}
			}

			// look down
			if y <= h-4 {
				if lines[y][x] == 'X' && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
					xmasOccurrences++
				}
			}

			// look up
			if y >= 3 {
				if lines[y][x] == 'X' && lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S' {
					xmasOccurrences++
				}
			}
		}
	}

	return xmasOccurrences
}

func part2(lines []string) int {
	xMASOccurrences := 0

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[0])-1; x++ {
			if lines[y][x] == 'A' {
				crosses := 0
				if (lines[y-1][x-1] == 'M' && lines[y+1][x+1] == 'S') || (lines[y-1][x-1] == 'S' && lines[y+1][x+1] == 'M') {
					crosses++
				}

				if (lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S') || (lines[y+1][x-1] == 'S' && lines[y-1][x+1] == 'M') {
					crosses++
				}

				if crosses == 2 {
					xMASOccurrences++
				}
			}
		}
	}

	return xMASOccurrences
}
