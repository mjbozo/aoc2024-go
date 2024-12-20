package day06

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"time"
)

type Position utils.Pair[int, int]
type Orientation utils.Triple[int, int, byte]

func Run() {
	input, err := utils.ReadInput("day06/input.txt", 6)
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
	grid, guardPos := buildGridMap(lines)
	guardLocations := buildGuardPath(&grid, &guardPos)
	return len(guardLocations)
}

func part2(lines []string) int {
	grid, guardPos := buildGridMap(lines)
	startPos := Position{First: guardPos.First, Second: guardPos.Second}
	guardLocations := buildGuardPath(&grid, &guardPos)
	guardLocations.Remove(startPos)

	c := make(chan int, len(guardLocations))
	for location := range guardLocations {
		go func(grid utils.Grid[GridCell], loc, start Position, c chan int) {
			c <- pathLoops(grid, start, loc.First, loc.Second)
		}(grid, location, startPos, c)
	}

	numLoopsConcurrent := 0
	for range len(guardLocations) {
		select {
		case x := <-c:
			numLoopsConcurrent += x
		}
	}

	return numLoopsConcurrent
}

// Convert text input to 2D array of GridCells, and cell where guard starts
func buildGridMap(mapLines []string) (utils.Grid[GridCell], Orientation) {
	grid := make(utils.Grid[GridCell], 0)
	guardPos := Orientation{Third: '^'}

	for y, line := range mapLines {
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

	return grid, guardPos
}

// Build HashSet of all positions guard visits in default map
func buildGuardPath(grid *utils.Grid[GridCell], guardPos *Orientation) utils.HashSet[Position] {
	width := len((*grid)[0])
	height := len(*grid)
	guardLocations := make(utils.HashSet[Position])

	for guardPos.First >= 0 && guardPos.First < width && guardPos.Second >= 0 && guardPos.Second < height {
		// add location to set
		guardLocations.Insert(Position{First: guardPos.First, Second: guardPos.Second})

		// check cell in front of guard and rotate if needed
		switch guardPos.Third {
		case '^':
			if guardPos.Second > 0 && (*grid)[guardPos.Second-1][guardPos.First].isObstacle {
				guardPos.Third = '>'
			}
		case '>':
			if guardPos.First < width-1 && (*grid)[guardPos.Second][guardPos.First+1].isObstacle {
				guardPos.Third = 'v'
			}
		case 'v':
			if guardPos.Second < height-1 && (*grid)[guardPos.Second+1][guardPos.First].isObstacle {
				guardPos.Third = '<'
			}
		case '<':
			if guardPos.First > 0 && (*grid)[guardPos.Second][guardPos.First-1].isObstacle {
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

	return guardLocations
}

// Checks if grid cell directly in front of guard is an obstacle
func isFacingObstacle(grid *utils.Grid[GridCell], width, height, x, y int, pos *Orientation) bool {
	switch pos.Third {
	case '^':
		if pos.Second > 0 && ((*grid)[pos.Second-1][pos.First].isObstacle || (x == pos.First && y == pos.Second-1)) {
			return true
		}
	case '>':
		if pos.First < width-1 && ((*grid)[pos.Second][pos.First+1].isObstacle || (x == pos.First+1 && y == pos.Second)) {
			return true
		}
	case 'v':
		if pos.Second < height-1 && ((*grid)[pos.Second+1][pos.First].isObstacle || (x == pos.First && y == pos.Second+1)) {
			return true
		}
	case '<':
		if pos.First > 0 && ((*grid)[pos.Second][pos.First-1].isObstacle || (x == pos.First-1 && y == pos.Second)) {
			return true
		}
	}

	return false
}

// Rotate guard position in the map
func rotate(grid *utils.Grid[GridCell], width, height, x, y int, pos *Orientation, prevLocations *utils.HashSet[Orientation]) bool {
	for isFacingObstacle(grid, width, height, x, y, pos) {
		if !(*prevLocations).Insert(*pos) {
			return true
		}

		switch pos.Third {
		case '^':
			if pos.Second > 0 && ((*grid)[pos.Second-1][pos.First].isObstacle || (x == pos.First && y == pos.Second-1)) {
				pos.Third = '>'
			}
		case '>':
			if pos.First < width-1 && ((*grid)[pos.Second][pos.First+1].isObstacle || (x == pos.First+1 && y == pos.Second)) {
				pos.Third = 'v'
			}
		case 'v':
			if pos.Second < height-1 && ((*grid)[pos.Second+1][pos.First].isObstacle || (x == pos.First && y == pos.Second+1)) {
				pos.Third = '<'
			}
		case '<':
			if pos.First > 0 && ((*grid)[pos.Second][pos.First-1].isObstacle || (x == pos.First-1 && y == pos.Second)) {
				pos.Third = '^'
			}
		}

	}

	return false
}

// Checks if given map grid results in a looping path
func pathLoops(grid utils.Grid[GridCell], startPos Position, newObstacleX, newObstacleY int) int {
	prevLocations := make(utils.HashSet[Orientation])
	guardPos := Orientation{First: startPos.First, Second: startPos.Second, Third: '^'}

	width := len(grid[0])
	height := len(grid)

	for guardPos.First >= 0 && guardPos.First < width && guardPos.Second >= 0 && guardPos.Second < height {
		loopFound := rotate(&grid, width, height, newObstacleX, newObstacleY, &guardPos, &prevLocations)
		if loopFound {
			return 1
		}

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

	return 0
}
