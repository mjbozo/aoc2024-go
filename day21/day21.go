package day21

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

var (
	GAP_NUM = Pos{x: 0, y: 3}
	A_NUM   = Pos{x: 2, y: 3}
	ZERO    = Pos{x: 1, y: 3}
	ONE     = Pos{x: 0, y: 2}
	TWO     = Pos{x: 1, y: 2}
	THREE   = Pos{x: 2, y: 2}
	FOUR    = Pos{x: 0, y: 1}
	FIVE    = Pos{x: 1, y: 1}
	SIX     = Pos{x: 2, y: 1}
	SEVEN   = Pos{x: 0, y: 0}
	EIGHT   = Pos{x: 1, y: 0}
	NINE    = Pos{x: 2, y: 0}

	GAP_DIR = Pos{x: 0, y: 0}
	A_DIR   = Pos{x: 2, y: 0}
	UP      = Pos{x: 1, y: 0}
	LEFT    = Pos{x: 0, y: 1}
	DOWN    = Pos{x: 1, y: 1}
	RIGHT   = Pos{x: 2, y: 1}

	GAP_VAL   = -1
	LEFT_VAL  = 0
	DOWN_VAL  = 1
	RIGHT_VAL = 2
	UP_VAL    = 3
)

type Pos struct {
	x int
	y int
}

type Robot struct {
	pos        Pos
	controller *Robot
}

func Run() {
	input, err := utils.ReadInput("day21/input.txt", 21)
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
	// notable property that after each iteration, robots will all be returned to starting positions
	dirRobot1 := Robot{
		pos: A_DIR,
	}

	dirRobot2 := Robot{
		pos:        A_DIR,
		controller: &dirRobot1,
	}

	doorRobot := Robot{
		pos:        A_NUM,
		controller: &dirRobot2,
	}

	total := 0
	for _, line := range lines {
		var moveset string
		numeric, _ := strconv.Atoi(line[:len(line)-1])

		// for each number, figure out moves to get there
		for _, x := range line {
			c := getCoord(x)
			xDiff := c.x - doorRobot.pos.x
			yDiff := c.y - doorRobot.pos.y

			horizontalMoves := make([]Pos, 0)
			verticalMoves := make([]Pos, 0)

			if xDiff < 0 {
				for range -xDiff {
					horizontalMoves = append(horizontalMoves, LEFT)
				}
			} else {
				for range xDiff {
					horizontalMoves = append(horizontalMoves, RIGHT)
				}
			}

			if yDiff < 0 {
				for range -yDiff {
					verticalMoves = append(verticalMoves, UP)
				}
			} else {
				for range yDiff {
					verticalMoves = append(verticalMoves, DOWN)
				}
			}

			if xDiff != 0 && yDiff != 0 {
				// need to compare vertical first and horizontal first
				// but also need to check if either will run into the gap
				horizontalFirstController := &Robot{}
				verticalFirstController := &Robot{}
				*horizontalFirstController = *doorRobot.controller
				*verticalFirstController = *doorRobot.controller

				var horizontalFirstMoveset string
				horizontalHitsGap := doorRobot.pos.y == GAP_NUM.y && doorRobot.pos.x+xDiff == GAP_NUM.x
				if !horizontalHitsGap {
					hThenV := make([]Pos, 0)
					hThenV = append(hThenV, horizontalMoves...)
					hThenV = append(hThenV, verticalMoves...)
					horizontalFirstMoveset = generateMoveset(hThenV, horizontalFirstController)
				}

				var verticalFirstMoveset string
				verticalHitsGap := doorRobot.pos.x == GAP_NUM.x && doorRobot.pos.y+yDiff == GAP_NUM.y
				if !verticalHitsGap {
					vThenH := make([]Pos, 0)
					vThenH = append(vThenH, verticalMoves...)
					vThenH = append(vThenH, horizontalMoves...)
					verticalFirstMoveset = generateMoveset(vThenH, verticalFirstController)
				}

				// set moveset accordingly
				if horizontalHitsGap {
					moveset += verticalFirstMoveset
					doorRobot.controller = verticalFirstController
				} else if verticalHitsGap {
					moveset += horizontalFirstMoveset
					doorRobot.controller = horizontalFirstController
				} else {
					if len(horizontalFirstMoveset) <= len(verticalFirstMoveset) {
						moveset += horizontalFirstMoveset
						doorRobot.controller = horizontalFirstController
					} else {
						moveset += verticalFirstMoveset
						doorRobot.controller = verticalFirstController
					}
				}
			} else if xDiff != 0 {
				// just go horizontal
				moveset += generateMoveset(horizontalMoves, doorRobot.controller)
			} else if yDiff != 0 {
				// just go vertical
				moveset += generateMoveset(verticalMoves, doorRobot.controller)
			}

			moveset += generateMoveset([]Pos{A_DIR}, doorRobot.controller)
			doorRobot.pos = c
		}

		fmt.Printf("MOVES FOR %s: %s (Len = %d, Num = %d)\n", line, moveset, len(moveset), numeric)
		total += numeric * len(moveset)
	}

	return total
}

func part2(lines []string) int {
	return 0
}

// Given a set of moves requested by next robot, figure out best moves to do those actions
func generateMoveset(desiredMoves []Pos, robot *Robot) string {
	if robot == nil {
		return getHumanInput(desiredMoves)
	}

	var moveset string

	for _, move := range desiredMoves {
		xDiff := move.x - robot.pos.x
		yDiff := move.y - robot.pos.y
		horizontalMoves := make([]Pos, 0)
		verticalMoves := make([]Pos, 0)

		if xDiff < 0 {
			for range -xDiff {
				horizontalMoves = append(horizontalMoves, LEFT)
			}
		} else {
			for range xDiff {
				horizontalMoves = append(horizontalMoves, RIGHT)
			}
		}

		if yDiff < 0 {
			for range -yDiff {
				verticalMoves = append(verticalMoves, UP)
			}
		} else {
			for range yDiff {
				verticalMoves = append(verticalMoves, DOWN)
			}
		}

		if xDiff != 0 && yDiff != 0 {
			// need to compare vertical first and horizontal first
			// but also need to check if either will run into the gap
			horizontalFirstController := &Robot{}
			verticalFirstController := &Robot{}
			if robot.controller != nil {
				*horizontalFirstController = *robot.controller
				*verticalFirstController = *robot.controller
			} else {
				horizontalFirstController = nil
				verticalFirstController = nil
			}

			var horizontalFirstMoveset string
			horizontalHitsGap := robot.pos.y == GAP_DIR.y && robot.pos.x+xDiff == GAP_DIR.x
			if !horizontalHitsGap {
				hThenV := make([]Pos, 0)
				hThenV = append(hThenV, horizontalMoves...)
				hThenV = append(hThenV, verticalMoves...)
				horizontalFirstMoveset = generateMoveset(hThenV, horizontalFirstController)
			}

			var verticalFirstMoveset string
			verticalHitsGap := robot.pos.x == GAP_DIR.x && robot.pos.y+yDiff == GAP_DIR.y
			if !verticalHitsGap {
				vThenH := make([]Pos, 0)
				vThenH = append(vThenH, verticalMoves...)
				vThenH = append(vThenH, horizontalMoves...)
				verticalFirstMoveset = generateMoveset(vThenH, verticalFirstController)
			}

			// set moveset accordingly
			if horizontalHitsGap {
				moveset += verticalFirstMoveset
				robot.controller = verticalFirstController
			} else if verticalHitsGap {
				moveset += horizontalFirstMoveset
				robot.controller = horizontalFirstController
			} else {
				if len(horizontalFirstMoveset) <= len(verticalFirstMoveset) {
					moveset += horizontalFirstMoveset
					robot.controller = horizontalFirstController
				} else {
					moveset += verticalFirstMoveset
					robot.controller = verticalFirstController
				}
			}

		} else if xDiff != 0 {
			// just go horizontal
			moveset += generateMoveset(horizontalMoves, robot.controller)
		} else if yDiff != 0 {
			// just go vertical
			moveset += generateMoveset(verticalMoves, robot.controller)
		}

		moveset += generateMoveset([]Pos{A_DIR}, robot.controller)
		robot.pos = move
	}

	return moveset
}

func getHumanInput(moves []Pos) string {
	var inputSequence string
	for _, move := range moves {
		switch move {
		case UP:
			inputSequence += "^"
		case DOWN:
			inputSequence += "v"
		case LEFT:
			inputSequence += "<"
		case RIGHT:
			inputSequence += ">"
		case A_DIR:
			inputSequence += "A"
		}
	}

	return inputSequence
}

func getCoord(val rune) Pos {
	switch val {
	case '0':
		return ZERO
	case '1':
		return ONE
	case '2':
		return TWO
	case '3':
		return THREE
	case '4':
		return FOUR
	case '5':
		return FIVE
	case '6':
		return SIX
	case '7':
		return SEVEN
	case '8':
		return EIGHT
	case '9':
		return NINE
	default:
		return A_NUM
	}
}
