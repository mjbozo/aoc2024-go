package day10

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Cell struct {
	x      int
	y      int
	height int
}

func Run() {
	input, err := utils.ReadInput("day10/input.txt", 10)
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
	grid, trailheads := buildGrid(lines)
	total := 0
	for _, trailhead := range trailheads {
		peaks := make(utils.HashSet[Cell])
		queue := make([]*Cell, 0)
		queue = append(queue, trailhead)

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			neighbours := getNeighbours(grid, current)
			for _, neighbour := range neighbours {
				if neighbour.height == 9 {
					peaks.Insert(*neighbour)
				} else {
					queue = append(queue, neighbour)
				}
			}
		}

		total += len(peaks)
	}

	return total
}

func part2(lines []string) int {
	grid, trailheads := buildGrid(lines)
	peaks := 0
	for _, trailhead := range trailheads {
		queue := make([]*Cell, 0)
		queue = append(queue, trailhead)

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			neighbours := getNeighbours(grid, current)
			for _, neighbour := range neighbours {
				if neighbour.height == 9 {
					peaks++
				} else {
					queue = append(queue, neighbour)
				}
			}
		}
	}

	return peaks
}

func buildGrid(lines []string) (utils.Grid[Cell], []*Cell) {
	grid := make(utils.Grid[Cell], 0)
	trailheads := make([]*Cell, 0)

	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			h, _ := strconv.Atoi(string(c))
			cell := Cell{x: x, y: y, height: h}
			row = append(row, cell)

			if h == 0 {
				trailheads = append(trailheads, &cell)
			}
		}
		grid = append(grid, row)
	}

	return grid, trailheads
}

func getNeighbours(grid utils.Grid[Cell], trailhead *Cell) []*Cell {
	neighbours := make([]*Cell, 0)
	width := len(grid[0])
	height := len(grid)

	if trailhead.x > 0 && grid[trailhead.y][trailhead.x-1].height == trailhead.height+1 {
		neighbours = append(neighbours, &grid[trailhead.y][trailhead.x-1])
	}
	if trailhead.y > 0 && grid[trailhead.y-1][trailhead.x].height == trailhead.height+1 {
		neighbours = append(neighbours, &grid[trailhead.y-1][trailhead.x])
	}
	if trailhead.x < width-1 && grid[trailhead.y][trailhead.x+1].height == trailhead.height+1 {
		neighbours = append(neighbours, &grid[trailhead.y][trailhead.x+1])
	}
	if trailhead.y < height-1 && grid[trailhead.y+1][trailhead.x].height == trailhead.height+1 {
		neighbours = append(neighbours, &grid[trailhead.y+1][trailhead.x])
	}

	return neighbours
}
