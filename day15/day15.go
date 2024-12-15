package day15

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

type Pos struct {
	x int
	y int
}

type Cell struct {
	pos Pos
	val byte
}

func (c Cell) String() string {
	return string(c.val)
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

	grid, robot := buildGrid(gridLines)

	for _, movement := range movements {
		switch movement {
		case '^':
			moveVertically(&grid, &robot, -1, 1)
		case 'v':
			moveVertically(&grid, &robot, 1, -1)
		case '<':
			moveHorizontally(&grid, &robot, -1, 1)
		case '>':
			moveHorizontally(&grid, &robot, 1, -1)
		}
	}

	return gpsSum(grid, 'O')
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	movements := parts[1]

	for i, line := range gridLines {
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")
		gridLines[i] = line
	}

	grid, robot := buildGrid(gridLines)

	for _, movement := range movements {
		switch movement {
		case '^':
			moveVerticallyThicc(&grid, &robot, -1, 1)
		case 'v':
			moveVerticallyThicc(&grid, &robot, 1, -1)
		case '<':
			moveHorizontally(&grid, &robot, -1, 1)
		case '>':
			moveHorizontally(&grid, &robot, 1, -1)
		}
	}

	//fmt.Println(grid)

	return gpsSum(grid, '[')
}

func buildGrid(gridLines []string) (utils.Grid[Cell], Pos) {
	robot := Pos{x: 0, y: 0}
	grid := make(utils.Grid[Cell], 0)
	for y, line := range gridLines {
		row := make([]Cell, 0)
		for x, c := range line {
			cell := Cell{pos: Pos{x: x, y: y}, val: byte(c)}
			row = append(row, cell)

			if byte(c) == '@' {
				robot.x = x
				robot.y = y
			}
		}
		grid = append(grid, row)
	}

	return grid, robot
}

func moveHorizontally(grid *utils.Grid[Cell], robot *Pos, forward, backward int) {
	x := robot.x + forward
	for (*grid)[robot.y][x].val != '.' && (*grid)[robot.y][x].val != '#' {
		x += forward
	}
	if (*grid)[robot.y][x].val == '.' {
		for x != robot.x {
			(*grid)[robot.y][x].val = (*grid)[robot.y][x+backward].val
			x += backward
		}
		(*grid)[robot.y][robot.x].val = '.'
		robot.x += forward
	}
}

func moveVertically(grid *utils.Grid[Cell], robot *Pos, forward, backward int) {
	y := robot.y + forward
	for (*grid)[y][robot.x].val != '.' && (*grid)[y][robot.x].val != '#' {
		y += forward
	}
	if (*grid)[y][robot.x].val == '.' {
		for y != robot.y {
			(*grid)[y][robot.x].val = (*grid)[y+backward][robot.x].val
			y += backward
		}
		(*grid)[robot.y][robot.x].val = '.'
		robot.y += forward
	}
}

func moveVerticallyThicc(grid *utils.Grid[Cell], robot *Pos, forward, backward int) {
	positionsToMove := make([]Cell, 0)
	queue := make([]Cell, 0)
	queue = append(queue, (*grid)[robot.y][robot.x])

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		alreadyMoving := false
		for _, cell := range positionsToMove {
			if cell.pos.x == current.pos.x && cell.pos.y == current.pos.y+forward {
				alreadyMoving = true
			}
		}

		if !alreadyMoving {
			positionsToMove = append(positionsToMove, Cell{pos: Pos{x: current.pos.x, y: current.pos.y + forward}})
		}

		if (*grid)[current.pos.y+forward][current.pos.x].val == '#' {
			// need to stop immeiately
			return
		}

		if (*grid)[current.pos.y+forward][current.pos.x].val == '[' {
			queue = append(queue, (*grid)[current.pos.y+forward][current.pos.x])
			queue = append(queue, (*grid)[current.pos.y+forward][current.pos.x+1])
		} else if (*grid)[current.pos.y+forward][current.pos.x].val == ']' {
			queue = append(queue, (*grid)[current.pos.y+forward][current.pos.x])
			queue = append(queue, (*grid)[current.pos.y+forward][current.pos.x-1])
		}
	}

	// should all be empty spots
	for i := len(positionsToMove) - 1; i >= 0; i-- {
		cell := positionsToMove[i]
		(*grid)[cell.pos.y][cell.pos.x].val = (*grid)[cell.pos.y+backward][cell.pos.x].val
		(*grid)[cell.pos.y+backward][cell.pos.x].val = '.'
	}

	(*grid)[robot.y][robot.x].val = '.'
	robot.y += forward
}

func gpsSum(grid utils.Grid[Cell], box byte) int {
	total := 0
	for y, line := range grid {
		for x, cell := range line {
			if cell.val == box {
				total += 100*y + x
			}
		}
	}
	return total
}
