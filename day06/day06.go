package day06

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"time"
)

func Run() {
	input, err := utils.ReadInput("day06/input.txt")
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

type GridCell struct {
	x          int
	y          int
	isObstacle bool
}

func part1(lines []string) int {
	grid := make(utils.Grid[GridCell], 0)
	guardPos := utils.Triple[int, int, byte]{First: 0, Second: 0, Third: '^'}

	for y, line := range lines {
		row := make([]GridCell, 0)
		for x, element := range line {
			cell := GridCell{x: x, y: y, isObstacle: element == '#'}
			row = append(row, cell)
			if element == '^' {
				guardPos.First = x
				guardPos.Second = y
			}
		}

		grid = append(grid, row)
	}

	width := len(grid[0])
	height := len(grid)
	guardLocations := make(utils.HashSet[utils.Pair[int, int]])

	for guardPos.First >= 0 && guardPos.First < width && guardPos.Second >= 0 && guardPos.Second < height {
		// add location to set
		guardLocations.Insert(utils.Pair[int, int]{First: guardPos.First, Second: guardPos.Second})

		// check cell in front of guard and rotate if needed
		switch guardPos.Third {
		case '^':
			if guardPos.Second > 0 && grid[guardPos.Second-1][guardPos.First].isObstacle {
				guardPos.Third = '>'
			}
		case '>':
			if guardPos.First < width-1 && grid[guardPos.Second][guardPos.First+1].isObstacle {
				guardPos.Third = 'v'
			}
		case 'v':
			if guardPos.Second < height-1 && grid[guardPos.Second+1][guardPos.First].isObstacle {
				guardPos.Third = '<'
			}
		case '<':
			if guardPos.First > 0 && grid[guardPos.Second][guardPos.First-1].isObstacle {
				guardPos.Third = '^'
			}
		}

		// move the guard depending on heading
		switch guardPos.Third {
		case '^':
			guardPos.Second--
		case '>':
			guardPos.First++
		case 'v':
			guardPos.Second++
		case '<':
			guardPos.First--
		}
	}

	return len(guardLocations)
}

func part2(lines []string) int {
	grid := make(utils.Grid[GridCell], 0)
	guardPos := utils.Triple[int, int, byte]{First: 0, Second: 0, Third: '^'}

	for y, line := range lines {
		row := make([]GridCell, 0)
		for x, element := range line {
			cell := GridCell{x: x, y: y, isObstacle: element == '#'}
			row = append(row, cell)
			if element == '^' {
				guardPos.First = x
				guardPos.Second = y
			}
		}

		grid = append(grid, row)
	}

	startPos := utils.Pair[int, int]{First: guardPos.First, Second: guardPos.Second}

	width := len(grid[0])
	height := len(grid)
	guardLocations := make(utils.HashSet[utils.Pair[int, int]])

	for guardPos.First >= 0 && guardPos.First < width && guardPos.Second >= 0 && guardPos.Second < height {
		// add location to set
		guardLocations.Insert(utils.Pair[int, int]{First: guardPos.First, Second: guardPos.Second})

		// check cell in front of guard and rotate if needed
		switch guardPos.Third {
		case '^':
			if guardPos.Second > 0 && grid[guardPos.Second-1][guardPos.First].isObstacle {
				guardPos.Third = '>'
			}
		case '>':
			if guardPos.First < width-1 && grid[guardPos.Second][guardPos.First+1].isObstacle {
				guardPos.Third = 'v'
			}
		case 'v':
			if guardPos.Second < height-1 && grid[guardPos.Second+1][guardPos.First].isObstacle {
				guardPos.Third = '<'
			}
		case '<':
			if guardPos.First > 0 && grid[guardPos.Second][guardPos.First-1].isObstacle {
				guardPos.Third = '^'
			}
		}

		// move the guard depending on heading
		switch guardPos.Third {
		case '^':
			guardPos.Second--
		case '>':
			guardPos.First++
		case 'v':
			guardPos.Second++
		case '<':
			guardPos.First--
		}
	}

	guardLocations.Remove(startPos)

	numLoops := 0

	for location := range guardLocations {
		grid[location.Second][location.First].isObstacle = true
		if loops(grid, startPos) {
			numLoops++
		}
		grid[location.Second][location.First].isObstacle = false
	}

	return numLoops
}

func loops(grid [][]GridCell, startPos utils.Pair[int, int]) bool {
	prevLocations := make(utils.HashSet[utils.Triple[int, int, byte]])
	guardPos := utils.Triple[int, int, byte]{First: startPos.First, Second: startPos.Second, Third: '^'}

	width := len(grid[0])
	height := len(grid)

	for guardPos.First >= 0 && guardPos.First < width && guardPos.Second >= 0 && guardPos.Second < height {
		// check cell in front of guard and rotate if needed
		frontClear := false
		for !frontClear {
			// add location to set
			if !prevLocations.Insert(guardPos) {
				return true
			}

			switch guardPos.Third {
			case '^':
				if guardPos.Second > 0 && grid[guardPos.Second-1][guardPos.First].isObstacle {
					guardPos.Third = '>'
				} else {
					frontClear = true
				}
			case '>':
				if guardPos.First < width-1 && grid[guardPos.Second][guardPos.First+1].isObstacle {
					guardPos.Third = 'v'
				} else {
					frontClear = true
				}
			case 'v':
				if guardPos.Second < height-1 && grid[guardPos.Second+1][guardPos.First].isObstacle {
					guardPos.Third = '<'
				} else {
					frontClear = true
				}
			case '<':
				if guardPos.First > 0 && grid[guardPos.Second][guardPos.First-1].isObstacle {
					guardPos.Third = '^'
				} else {
					frontClear = true
				}
			}
		}

		// move the guard depending on heading
		switch guardPos.Third {
		case '^':
			guardPos.Second--
		case '>':
			guardPos.First++
		case 'v':
			guardPos.Second++
		case '<':
			guardPos.First--
		}
	}

	return false
}

