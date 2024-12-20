package day12

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"slices"
	"time"
)

type Cell struct {
	x      int
	y      int
	region byte
}

type FenceKey struct {
	horizontal bool
	constVal   int
}

type Fence struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
}

func Run() {
	input, err := utils.ReadInput("day12/input.txt", 12)
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
	garden := make(utils.Grid[Cell], 0)
	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			row = append(row, Cell{x: x, y: y, region: byte(c)})
		}
		garden = append(garden, row)
	}

	width := len(garden[0])
	height := len(garden)
	visited := make(utils.HashSet[Cell])

	total := 0
	for y, line := range lines {
		for x := range line {
			plot := garden[y][x]
			if !visited.Contains(plot) {
				perimeter := 0
				area := 0
				queue := make([]Cell, 0)
				queue = append(queue, plot)

				for len(queue) != 0 {
					current := queue[0]
					queue = queue[1:]
					if visited.Contains(current) {
						continue
					}

					area++
					neighbours := getNeighbours(&garden, &current, width, height)
					perimeter += 4 - len(neighbours)
					for _, neighbour := range neighbours {
						if !visited.Contains(*neighbour) {
							queue = append(queue, *neighbour)
						}
					}

					visited.Insert(current)
				}

				//fmt.Printf("Adding %d (area = %d, perimeter = %d) for region %q\n", area*perimeter, area, perimeter, plot.region)
				total += area * perimeter
			}
		}
	}

	return total
}

func part2(lines []string) int {
	garden := make(utils.Grid[Cell], 0)
	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			row = append(row, Cell{x: x, y: y, region: byte(c)})
		}
		garden = append(garden, row)
	}

	width := len(garden[0])
	height := len(garden)
	visited := make(utils.HashSet[Cell])

	total := 0
	for y, line := range lines {
		for x := range line {
			plot := garden[y][x]
			if !visited.Contains(plot) {
				fences := make(map[FenceKey][]Fence)
				area := 0
				sides := 0
				queue := make([]Cell, 0)
				queue = append(queue, plot)

				for len(queue) != 0 {
					current := queue[0]
					queue = queue[1:]
					if visited.Contains(current) {
						continue
					}

					area++
					neighbours := getNeighboursAndFences(&garden, &current, width, height, &fences)

					for _, neighbour := range neighbours {
						if !visited.Contains(*neighbour) {
							queue = append(queue, *neighbour)
						}
					}

					visited.Insert(current)
				}

				//fmt.Println(fences)

				for key, val := range fences {
					sides++
					if key.horizontal {
						slices.SortFunc(val, func(a, b Fence) int { return a.xStart - b.xStart })
						for i := 1; i < len(val); i++ {
							if val[i].yStart == 0 || val[i].yStart == height {
								if val[i].xStart != val[i-1].xEnd {
									sides++
								}
							} else {
								above1 := garden[val[i].yStart-1][val[i].xStart-1]
								above2 := garden[val[i].yStart-1][val[i].xStart]
								below1 := garden[val[i].yStart][val[i].xStart]
								below2 := garden[val[i].yStart][val[i].xStart-1]
								if val[i].xStart != val[i-1].xEnd || (above1.region != above2.region && below1.region != below2.region) {
									sides++
								}
							}
						}
					} else {
						slices.SortFunc(val, func(a, b Fence) int { return a.yStart - b.yStart })
						for i := 1; i < len(val); i++ {
							if val[i].xStart == 0 || val[i].xStart == width {
								if val[i].yStart != val[i-1].yEnd {
									sides++
								}
							} else {
								left1 := garden[val[i].yStart-1][val[i].xStart-1]
								left2 := garden[val[i].yStart][val[i].xStart-1]
								right1 := garden[val[i].yStart-1][val[i].xStart]
								right2 := garden[val[i].yStart][val[i].xStart]
								if val[i].yStart != val[i-1].yEnd || (left1.region != left2.region && right1.region != right2.region) {
									sides++
								}
							}
						}
					}
				}

				total += area * sides
			}
		}
	}

	return total
}

func getNeighbours(garden *utils.Grid[Cell], current *Cell, width, height int) []*Cell {
	neighbours := make([]*Cell, 0)

	if current.x > 0 && (*garden)[current.y][current.x-1].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y][current.x-1])
	}
	if current.y > 0 && (*garden)[current.y-1][current.x].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y-1][current.x])
	}
	if current.x < width-1 && (*garden)[current.y][current.x+1].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y][current.x+1])
	}
	if current.y < height-1 && (*garden)[current.y+1][current.x].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y+1][current.x])
	}

	return neighbours
}

func getNeighboursAndFences(garden *utils.Grid[Cell], current *Cell, width, height int, fences *map[FenceKey][]Fence) []*Cell {
	neighbours := make([]*Cell, 0)

	if current.x > 0 && (*garden)[current.y][current.x-1].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y][current.x-1])
	} else {
		fenceKey := FenceKey{horizontal: false, constVal: current.x}
		fence := Fence{xStart: current.x, xEnd: current.x, yStart: current.y, yEnd: current.y + 1}
		if currentFences, ok := (*fences)[fenceKey]; ok {
			currentFences = append(currentFences, fence)
			(*fences)[fenceKey] = currentFences
		} else {
			(*fences)[fenceKey] = []Fence{fence}
		}
	}

	if current.y > 0 && (*garden)[current.y-1][current.x].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y-1][current.x])
	} else {
		fenceKey := FenceKey{horizontal: true, constVal: current.y}
		fence := Fence{xStart: current.x, xEnd: current.x + 1, yStart: current.y, yEnd: current.y}
		if currentFences, ok := (*fences)[fenceKey]; ok {
			currentFences = append(currentFences, fence)
			(*fences)[fenceKey] = currentFences
		} else {
			(*fences)[fenceKey] = []Fence{fence}
		}
	}

	if current.x < width-1 && (*garden)[current.y][current.x+1].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y][current.x+1])
	} else {
		fenceKey := FenceKey{horizontal: false, constVal: current.x + 1}
		fence := Fence{xStart: current.x + 1, xEnd: current.x + 1, yStart: current.y, yEnd: current.y + 1}
		if currentFences, ok := (*fences)[fenceKey]; ok {
			currentFences = append(currentFences, fence)
			(*fences)[fenceKey] = currentFences
		} else {
			(*fences)[fenceKey] = []Fence{fence}
		}
	}

	if current.y < height-1 && (*garden)[current.y+1][current.x].region == current.region {
		neighbours = append(neighbours, &(*garden)[current.y+1][current.x])
	} else {
		fenceKey := FenceKey{horizontal: true, constVal: current.y + 1}
		fence := Fence{xStart: current.x, xEnd: current.x + 1, yStart: current.y + 1, yEnd: current.y + 1}
		if currentFences, ok := (*fences)[fenceKey]; ok {
			currentFences = append(currentFences, fence)
			(*fences)[fenceKey] = currentFences
		} else {
			(*fences)[fenceKey] = []Fence{fence}
		}
	}

	return neighbours
}
