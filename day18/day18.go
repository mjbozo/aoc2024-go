package day18

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	x   int
	y   int
	val byte
}

type Pos struct {
	x int
	y int
}

func (c Cell) String() string {
	return string(c.val)
}

func Run() {
	input, err := utils.ReadInput("day18/input.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input, 71, 71, 1024)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input, 71, 71, 1024)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %s (%v)\n", part2Result, elapsed)
}

func part1(lines []string, w, h, n int) int {
	grid := make(utils.Grid[Cell], 0)
	for y := range h {
		row := make([]Cell, 0)
		for x := range w {
			row = append(row, Cell{x: x, y: y, val: '.'})
		}
		grid = append(grid, row)
	}

	for i := range n {
		b := lines[i]
		coords := strings.Split(b, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x].val = '#'
	}

	//fmt.Println(grid)

	scores := make(map[Pos]int)
	start := Pos{x: 0, y: 0}
	queue := make(chan Pos, w*h)
	queue <- start

	for len(queue) > 0 {
		current := <-queue
		currentScore := scores[current]

		neighbours := getNeighbours(&grid, &current)
		for _, neighbour := range neighbours {
			neighbourPos := Pos{x: neighbour.x, y: neighbour.y}
			if neighbourScore, ok := scores[neighbourPos]; ok {
				if neighbourScore > currentScore+1 {
					scores[neighbourPos] = currentScore + 1
					queue <- neighbourPos
				}
			} else {
				scores[neighbourPos] = currentScore + 1
				queue <- neighbourPos
			}
		}
	}

	endPos := Pos{x: w - 1, y: h - 1}
	return scores[endPos]
}

func part2(lines []string, w, h, n int) string {
	grid := make(utils.Grid[Cell], 0)
	for y := range h {
		row := make([]Cell, 0)
		for x := range w {
			row = append(row, Cell{x: x, y: y, val: '.'})
		}
		grid = append(grid, row)
	}

	for i := range n {
		b := lines[i]
		coords := strings.Split(b, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x].val = '#'
	}

	lastByte := ""
	for i := n; i < len(lines); i++ {
		b := lines[i]
		lastByte = b
		coords := strings.Split(b, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x].val = '#'

		scores := make(map[Pos]int)
		start := Pos{x: 0, y: 0}
		queue := make(chan Pos, w*h)
		queue <- start

		for len(queue) > 0 {
			current := <-queue
			currentScore := scores[current]

			neighbours := getNeighbours(&grid, &current)
			for _, neighbour := range neighbours {
				neighbourPos := Pos{x: neighbour.x, y: neighbour.y}
				if neighbourScore, ok := scores[neighbourPos]; ok {
					if neighbourScore > currentScore+1 {
						scores[neighbourPos] = currentScore + 1
						queue <- neighbourPos
					}
				} else {
					scores[neighbourPos] = currentScore + 1
					queue <- neighbourPos
				}
			}
		}

		endPos := Pos{x: w - 1, y: h - 1}
		if _, ok := scores[endPos]; !ok {
			break
		}
	}

	return lastByte
}

func getNeighbours(grid *utils.Grid[Cell], current *Pos) []Cell {
	neighbours := make([]Cell, 0)

	if current.y > 0 && (*grid)[current.y-1][current.x].val != '#' {
		neighbours = append(neighbours, (*grid)[current.y-1][current.x])
	}

	if current.y < len(*grid)-1 && (*grid)[current.y+1][current.x].val != '#' {
		neighbours = append(neighbours, (*grid)[current.y+1][current.x])
	}

	if current.x > 0 && (*grid)[current.y][current.x-1].val != '#' {
		neighbours = append(neighbours, (*grid)[current.y][current.x-1])
	}

	if current.x < len((*grid)[0])-1 && (*grid)[current.y][current.x+1].val != '#' {
		neighbours = append(neighbours, (*grid)[current.y][current.x+1])
	}

	return neighbours
}

