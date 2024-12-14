package day14

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	px int
	py int
	vx int
	vy int
}

type Pos struct {
	x int
	y int
}

func (r *Robot) move(w, h int) {
	nextX := r.px + r.vx
	if nextX < 0 {
		nextX += w
	} else if nextX >= w {
		nextX -= w
	}

	nextY := r.py + r.vy
	if nextY < 0 {
		nextY += h
	} else if nextY >= h {
		nextY -= h
	}

	r.px = nextX
	r.py = nextY
}

func Run() {
	input, err := utils.ReadInput("day14/input.txt")
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input, 101, 103)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input, 101, 103)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

func part1(lines []string, w, h int) int {
	robots := make([]Robot, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		position := parts[0]
		velocity := parts[1]

		posElements := strings.Split(position, ",")
		px, _ := strconv.Atoi(posElements[0][2:])
		py, _ := strconv.Atoi(posElements[1])

		velElements := strings.Split(velocity, ",")
		vx, _ := strconv.Atoi(velElements[0][2:])
		vy, _ := strconv.Atoi(velElements[1])

		robot := Robot{px: px, py: py, vx: vx, vy: vy}
		robots = append(robots, robot)
	}

	for range 100 {
		for i, robot := range robots {
			robot.move(w, h)
			robots[i] = robot
		}
	}

	quadrant1 := 0
	quadrant2 := 0
	quadrant3 := 0
	quadrant4 := 0

	for _, robot := range robots {
		if robot.px < w/2 && robot.py < h/2 {
			quadrant1++
		}
		if robot.px > w/2 && robot.py < h/2 {
			quadrant2++
		}
		if robot.px < w/2 && robot.py > h/2 {
			quadrant3++
		}
		if robot.px > w/2 && robot.py > h/2 {
			quadrant4++
		}
	}

	return quadrant1 * quadrant2 * quadrant3 * quadrant4
}

func part2(lines []string, w, h int) int {
	robots := make([]Robot, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		position := parts[0]
		velocity := parts[1]

		posElements := strings.Split(position, ",")
		px, _ := strconv.Atoi(posElements[0][2:])
		py, _ := strconv.Atoi(posElements[1])

		velElements := strings.Split(velocity, ",")
		vx, _ := strconv.Atoi(velElements[0][2:])
		vy, _ := strconv.Atoi(velElements[1])

		robot := Robot{px: px, py: py, vx: vx, vy: vy}
		robots = append(robots, robot)
	}

	n := 1
	for {
		positions := make(map[Pos]int)
		for i := range robots {
			robots[i].move(w, h)
			robotPos := Pos{x: robots[i].px, y: robots[i].py}
			positions[robotPos]++
		}

		seenPositions := make(utils.HashSet[Pos])
		for _, robot := range robots {
			p := Pos{x: robot.px, y: robot.py}
			if seenPositions.Contains(p) {
				continue
			}

			numberInBlob := 0
			queue := make([]Pos, 0)
			queue = append(queue, p)

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				seenPositions.Insert(current)
				numberInBlob += positions[current]

				up := Pos{x: current.x, y: current.y - 1}
				down := Pos{x: current.x, y: current.y + 1}
				left := Pos{x: current.x - 1, y: current.y}
				right := Pos{x: current.x + 1, y: current.y}

				if _, exists := positions[up]; !seenPositions.Contains(up) && exists {
					queue = append(queue, up)
				}
				if _, exists := positions[down]; !seenPositions.Contains(down) && exists {
					queue = append(queue, down)
				}
				if _, exists := positions[left]; !seenPositions.Contains(left) && exists {
					queue = append(queue, left)
				}
				if _, exists := positions[right]; !seenPositions.Contains(right) && exists {
					queue = append(queue, right)
				}
			}

			if numberInBlob > len(robots)/2 {
				for y := range h {
					row := ""
					for x := range w {
						if _, ok := positions[Pos{x: x, y: y}]; ok {
							row += "█"
						} else {
							row += " "
						}
					}
					fmt.Println(row)
				}
				return n
			}
		}

		n++
	}
}

