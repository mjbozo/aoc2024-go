package day15

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

type Cell struct {
	x   int
	y   int
	val byte
}

func (c Cell) String() string {
	return string(c.val)
}

type Pos struct {
	x int
	y int
}

func Run() {
	input, err := utils.ReadInputRaw("day15/input.txt")
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

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	movements := parts[1]
	robot := Pos{x: 0, y: 0}

	grid := make(utils.Grid[Cell], 0)
	for y, line := range gridLines {
		row := make([]Cell, 0)
		for x, c := range line {
			cell := Cell{x: x, y: y, val: byte(c)}
			row = append(row, cell)

			if byte(c) == '@' {
				robot.x = x
				robot.y = y
			}
		}
		grid = append(grid, row)
	}

	for _, movement := range movements {
		switch movement {
		case '^':
			y := robot.y - 1
			for grid[y][robot.x].val == byte('O') {
				y--
			}
			if grid[y][robot.x].val == byte('.') {
				for y < robot.y {
					grid[y][robot.x].val = grid[y+1][robot.x].val
					y++
				}
				grid[robot.y][robot.x].val = '.'
				robot.y--
			}

		case 'v':
			y := robot.y + 1
			for grid[y][robot.x].val == byte('O') {
				y++
			}
			if grid[y][robot.x].val == byte('.') {
				for y > robot.y {
					grid[y][robot.x].val = grid[y-1][robot.x].val
					y--
				}
				grid[robot.y][robot.x].val = '.'
				robot.y++
			}

		case '<':
			x := robot.x - 1
			for grid[robot.y][x].val == byte('O') {
				x--
			}
			if grid[robot.y][x].val == byte('.') {
				for x < robot.x {
					grid[robot.y][x].val = grid[robot.y][x+1].val
					x++
				}
				grid[robot.y][robot.x].val = '.'
				robot.x--
			}

		case '>':
			x := robot.x + 1
			for grid[robot.y][x].val == byte('O') {
				x++
			}
			if grid[robot.y][x].val == byte('.') {
				for x > robot.x {
					grid[robot.y][x].val = grid[robot.y][x-1].val
					x--
				}
				grid[robot.y][robot.x].val = '.'
				robot.x++
			}
		}
	}

	total := 0
	for y, line := range grid {
		for x, cell := range line {
			if cell.val == 'O' {
				total += 100*y + x
			}
		}
	}

	return total
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	movements := parts[1]
	robot := Pos{x: 0, y: 0}

	for i, line := range gridLines {
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")
		gridLines[i] = line
	}

	for _, line := range gridLines {
		fmt.Println(line)
	}

	grid := make(utils.Grid[Cell], 0)
	for y, line := range gridLines {
		row := make([]Cell, 0)
		for x, c := range line {
			cell := Cell{x: x, y: y, val: byte(c)}
			row = append(row, cell)

			if byte(c) == '@' {
				robot.x = x
				robot.y = y
			}
		}
		grid = append(grid, row)
	}

	//fmt.Println(grid)

	for _, movement := range movements {
		switch movement {
		case '^':
			positionsToMove := make([]Cell, 0)
			hasWall := false
			queue := make([]Cell, 0)
			queue = append(queue, grid[robot.y][robot.x])

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				alreadyMoving := false
				for _, pos := range positionsToMove {
					if pos.x == current.x && pos.y == current.y-1 {
						alreadyMoving = true
					}
				}

				if !alreadyMoving {
					positionsToMove = append(positionsToMove, Cell{x: current.x, y: current.y - 1})
				}

				if grid[current.y-1][current.x].val == '#' {
					// need to stop immediately
					hasWall = true
					break
				}

				if grid[current.y-1][current.x].val == '[' {
					queue = append(queue, grid[current.y-1][current.x])
					queue = append(queue, grid[current.y-1][current.x+1])
				} else if grid[current.y-1][current.x].val == ']' {
					queue = append(queue, grid[current.y-1][current.x])
					queue = append(queue, grid[current.y-1][current.x-1])
				}
			}

			if hasWall {
				break
			}

			// should all be empty spots
			for i := len(positionsToMove) - 1; i >= 0; i-- {
				pos := positionsToMove[i]
				grid[pos.y][pos.x].val = grid[pos.y+1][pos.x].val
				grid[pos.y+1][pos.x].val = '.'
			}

			grid[robot.y][robot.x].val = '.'
			robot.y--

		case 'v':
			positionsToMove := make([]Cell, 0)
			hasWall := false
			queue := make([]Cell, 0)
			queue = append(queue, grid[robot.y][robot.x])

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				alreadyMoving := false
				for _, pos := range positionsToMove {
					if pos.x == current.x && pos.y == current.y+1 {
						alreadyMoving = true
					}
				}

				if !alreadyMoving {
					positionsToMove = append(positionsToMove, Cell{x: current.x, y: current.y + 1})
				}

				if grid[current.y+1][current.x].val == '#' {
					// need to stop immediately
					hasWall = true
					break
				}

				if grid[current.y+1][current.x].val == '[' {
					queue = append(queue, grid[current.y+1][current.x])
					queue = append(queue, grid[current.y+1][current.x+1])
				} else if grid[current.y+1][current.x].val == ']' {
					queue = append(queue, grid[current.y+1][current.x])
					queue = append(queue, grid[current.y+1][current.x-1])
				}
			}

			if hasWall {
				break
			}

			// should all be empty spots
			for i := len(positionsToMove) - 1; i >= 0; i-- {
				pos := positionsToMove[i]
				grid[pos.y][pos.x].val = grid[pos.y-1][pos.x].val
				grid[pos.y-1][pos.x].val = '.'
			}

			grid[robot.y][robot.x].val = '.'
			robot.y++

		case '<':
			x := robot.x - 1
			for grid[robot.y][x].val == ']' || grid[robot.y][x].val == '[' {
				x--
			}
			if grid[robot.y][x].val == byte('.') {
				for x < robot.x {
					grid[robot.y][x].val = grid[robot.y][x+1].val
					x++
				}
				grid[robot.y][robot.x].val = '.'
				robot.x--
			}

		case '>':
			x := robot.x + 1
			for grid[robot.y][x].val == '[' || grid[robot.y][x].val == ']' {
				x++
			}
			if grid[robot.y][x].val == byte('.') {
				for x > robot.x {
					grid[robot.y][x].val = grid[robot.y][x-1].val
					x--
				}
				grid[robot.y][robot.x].val = '.'
				robot.x++
			}
		}

		//fmt.Println(string(movement))
		//fmt.Println(grid)
	}

	total := 0
	for y, line := range grid {
		for x, cell := range line {
			if cell.val == '[' {
				total += 100*y + x
			}
		}
	}

	return total
}

