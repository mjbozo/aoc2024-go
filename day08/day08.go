package day08

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"time"
)

type Pos utils.Pair[int, int]
type Set = utils.HashSet[Pos]

func Run() {
	input, err := utils.ReadInput("day08/input.txt", 8)
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
	height := len(lines)
	width := len(lines[0])

	antinodeLocations := make(Set)
	frequencyLocations := make(map[byte][]Pos)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if lines[y][x] != '.' {
				locations, ok := frequencyLocations[lines[y][x]]
				if ok {
					// track antinode for each new pair
					for _, location := range locations {
						xDiff := location.First - x
						yDiff := location.Second - y

						antinode1 := Pos{First: location.First + xDiff, Second: location.Second + yDiff}
						antinode2 := Pos{First: x - xDiff, Second: y - yDiff}

						if isInMap(antinode1, width, height) {
							antinodeLocations.Insert(antinode1)
						}

						if isInMap(antinode2, width, height) {
							antinodeLocations.Insert(antinode2)
						}
					}
				}
				locations = append(locations, Pos{First: x, Second: y})
				frequencyLocations[lines[y][x]] = locations
			}
		}
	}

	return len(antinodeLocations)
}

func part2(lines []string) int {
	height := len(lines)
	width := len(lines[0])

	antinodeLocations := make(utils.HashSet[Pos])
	frequencyLocations := make(map[byte][]Pos)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if lines[y][x] != '.' {
				locations, ok := frequencyLocations[lines[y][x]]
				if ok {
					for _, location := range locations {
						xDiff := location.First - x
						yDiff := location.Second - y

						antinodeLocations.Insert(location)
						antinodeLocations.Insert(Pos{First: x, Second: y})

						antinode1 := Pos{First: location.First + xDiff, Second: location.Second + yDiff}
						for isInMap(antinode1, width, height) {
							antinodeLocations.Insert(antinode1)
							antinode1.First += xDiff
							antinode1.Second += yDiff
						}

						antinode2 := Pos{First: x - xDiff, Second: y - yDiff}
						for isInMap(antinode2, width, height) {
							antinodeLocations.Insert(antinode2)
							antinode2.First -= xDiff
							antinode2.Second -= yDiff
						}
					}
				}
				locations = append(locations, Pos{First: x, Second: y})
				frequencyLocations[lines[y][x]] = locations
			}
		}
	}

	return len(antinodeLocations)
}

func isInMap(antinode Pos, width, height int) bool {
	return antinode.First < width && antinode.First >= 0 && antinode.Second < height && antinode.Second >= 0
}
