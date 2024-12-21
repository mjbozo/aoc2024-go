package day20

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"slices"
	"time"
)

type Cell struct {
	x   int
	y   int
	val byte
	idx int
}

type Pos struct {
	x int
	y int
}

func Run() {
	input, err := utils.ReadInput("day20/input.txt", 20)
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
	grid := make(utils.Grid[Cell], 0)
	var startPos Cell

	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			row = append(row, Cell{x: x, y: y, val: byte(c), idx: -1})
			if c == 'S' {
				row[x].val = '.'
				row[x].idx = 0
				startPos = row[x]
			}
			if c == 'E' {
				row[x].val = '.'
			}
		}
		grid = append(grid, row)
	}

	w := len(grid[0])
	h := len(grid)

	seen := make(utils.HashSet[Cell])
	seen.Insert(startPos)

	queue := make(chan Cell, w*h)
	queue <- startPos
	dirs := Dirs(1)

	for len(queue) > 0 {
		current := <-queue
		neighbours := getNeighbours(&grid, &current, dirs)
		for _, neighbour := range neighbours {
			if !seen.Contains(neighbour) {
				neighbour.idx = current.idx + 1
				grid[neighbour.y][neighbour.x].idx = current.idx + 1
				queue <- neighbour
				seen.Insert(neighbour)
			}
		}
	}

	numCheats := 0
	cheatMap := make(map[int]int)
	dirs = Dirs(2)

	for current := range seen {
		cheats := findCheats(&grid, &current, dirs)
		for _, cheat := range cheats {
			cheatMap[cheat]++
			if cheat >= 100 {
				numCheats++
			}
		}
	}

	return numCheats
}

func part2(lines []string) int {
	grid := make(utils.Grid[Cell], 0)
	var startPos Cell

	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			row = append(row, Cell{x: x, y: y, val: byte(c), idx: -1})
			if c == 'S' {
				row[x].val = '.'
				row[x].idx = 0
				startPos = row[x]
			}
			if c == 'E' {
				row[x].val = '.'
			}
		}
		grid = append(grid, row)
	}

	w := len(grid[0])
	h := len(grid)

	seen := make(utils.HashSet[Cell])
	seen.Insert(startPos)

	queue := make(chan Cell, w*h)
	queue <- startPos
	dirs := Dirs(1)

	for len(queue) > 0 {
		current := <-queue
		neighbours := getNeighbours(&grid, &current, dirs)
		for _, neighbour := range neighbours {
			if !seen.Contains(neighbour) {
				neighbour.idx = current.idx + 1
				grid[neighbour.y][neighbour.x].idx = current.idx + 1
				queue <- neighbour
				seen.Insert(neighbour)
			}
		}
	}

	path := make([]Cell, 0)

	for spot := range seen {
		path = append(path, spot)
	}

	slices.SortFunc(path, func(a, b Cell) int {
		return a.idx - b.idx
	})

	numCheats := 0
	minCheatThreshold := 100

	for i, spot := range path {
		if i >= len(path)-(minCheatThreshold+2)+1 {
			break
		}

		possiblePositions := path[i+minCheatThreshold+2:]
		for _, position := range possiblePositions {
			p1 := utils.Pair[int, int]{First: spot.x, Second: spot.y}
			p2 := utils.Pair[int, int]{First: position.x, Second: position.y}
			manhattan := utils.Manhattan(p1, p2)
			if manhattan <= 20 && position.idx-spot.idx-manhattan >= minCheatThreshold {
				numCheats++
			}
		}
	}

	return numCheats
}

func getNeighbours(grid *utils.Grid[Cell], current *Cell, dirs []Pos) []Cell {
	neighbours := make([]Cell, 0)

	for _, dir := range dirs {
		if isInGrid(grid, current, dir) && (*grid)[current.y+dir.y][current.x+dir.x].val == '.' {
			neighbours = append(neighbours, (*grid)[current.y+dir.y][current.x+dir.x])
		}
	}

	return neighbours
}

func findCheats(grid *utils.Grid[Cell], current *Cell, dirs []Pos) []int {
	cheats := make([]int, 0)
	for _, dir := range dirs {
		if isInGrid(grid, current, dir) && (*grid)[current.y+dir.y][current.x+dir.x].val == '.' {
			cheatValue := (*grid)[current.y+dir.y][current.x+dir.x].idx - (*grid)[current.y][current.x].idx - 2
			cheats = append(cheats, cheatValue)
		}
	}
	return cheats
}

func isInGrid(grid *utils.Grid[Cell], current *Cell, dir Pos) bool {
	w := len((*grid)[0])
	h := len(*grid)

	newX := current.x + dir.x
	newY := current.y + dir.y
	return newX >= 0 && newX < w && newY >= 0 && newY < h
}

func Dirs(d int) []Pos {
	return []Pos{{x: 0, y: d}, {x: 0, y: -d}, {x: d, y: 0}, {x: -d, y: 0}}
}
