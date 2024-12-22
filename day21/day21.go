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

func getDirPadNeighbours(n int) []int {
	switch n {
	case 0:
		return []int{1}
	case 1:
		return []int{0, 2, 3}
	case 2:
		return []int{1, 10}
	case 3:
		return []int{2, 10}
	case 10:
		return []int{2, 3}
	}

	return []int{}
}

// type Keypad = [][]int // gap = -1, A = 10, < = 0, v = 1, > = 2, ^ = 3

type Keypad = utils.Grid[int]

type Pos struct {
	x int
	y int
}

type Robot struct {
	name       string
	keypad     Keypad
	pos        Pos
	controller *Robot
}

func (r Robot) String() string {
	return fmt.Sprintf("%s{Pos: (%d, %d), Keypad:\n%v}", r.name, r.pos.x, r.pos.y, r.keypad)
}

type State struct {
	numPadPos Pos
	keypad1   Pos
	keypad2   Pos
	dir       Pos
}

type ScoredState struct {
	state   State
	moveSet string
}

type DirPadState struct {
	keypad1 Pos
	keypad2 Pos
	dir     Pos
}

type ScoredDirPadState struct {
	dirPadState DirPadState
	moveSet     string
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
	numericKeypad := Keypad{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
		{-1, 0, 10},
	}

	directionalKeypad1 := Keypad{
		{-1, 3, 10},
		{0, 1, 2},
	}
	directionalKeypad2 := Keypad{
		{-1, 3, 10},
		{0, 1, 2},
	}

	moves := make(map[State]string)

	// solution space is small enough to write them out

	//test := dirRobot2.moveToAndPress(2, 1)
	//test := doorRobot.moveToAndPress(ZERO.x, ZERO.y)
	//fmt.Println(test)
	//test = doorRobot.moveToAndPress(TWO.x, TWO.y)
	//fmt.Println(test)

	total := 0
	lines = []string{"0"}
	for _, line := range lines {
		dirRobot1 := Robot{
			name:   "DirRobot1",
			keypad: directionalKeypad1,
			pos:    A_DIR,
		}

		dirRobot2 := Robot{
			name:       "DirRobot2",
			keypad:     directionalKeypad2,
			pos:        A_DIR,
			controller: &dirRobot1,
		}

		doorRobot := Robot{
			name:       "DoorRobot",
			keypad:     numericKeypad,
			pos:        A_NUM,
			controller: &dirRobot2,
		}

		numeric, _ := strconv.Atoi(line[:len(line)-1])

		startState := State{numPadPos: A_NUM, keypad1: Pos{x: 2, y: 0}, keypad2: Pos{x: 2, y: 0}}
		numPadQueue := make(chan State, 11*5*5*4)
		numPadQueue <- startState

		finalStates := make([]ScoredState, 0)
		var inputSequence string
		for _, c := range line {
			coord := getCoord(c)
			for len(numPadQueue) > 0 {
				current := <-numPadQueue
				fmt.Println("1")
				neighbours := getNumPadNeighbourStates(numericKeypad, current, doorRobot)

				for _, neighbour := range neighbours {
					if oldMoves, ok := moves[neighbour.state]; ok {
						if len(neighbour.moveSet) < len(oldMoves) {
							moves[neighbour.state] = neighbour.moveSet
							numPadQueue <- neighbour.state
						}
					} else {
						moves[neighbour.state] = neighbour.moveSet
						numPadQueue <- neighbour.state
					}

					if neighbour.state.numPadPos == coord {
						finalStates = append(finalStates, neighbour)
					}
				}
			}

			fmt.Println("NUMPAD Final States:")
			fmt.Println(finalStates)
			//inputSequence += moveToAndPress(doorRobot, coord.x, coord.y)
			//fmt.Println(inputSequence)
			//inputSequence += dirRobot1.sendToController()
		}
		fmt.Printf("%s: Len = %d Value = %d\n", inputSequence, len(inputSequence), numeric)
		total += numeric * len(inputSequence)
	}

	fmt.Printf("\n\nTest Area:\n")
	dirRobot1 := Robot{
		name:   "DirRobot1",
		keypad: directionalKeypad1,
		pos:    A_DIR,
	}
	dirRobot2 := Robot{
		name:       "DirRobot2",
		keypad:     directionalKeypad2,
		pos:        A_DIR,
		controller: &dirRobot1,
	}
	//currentState := DirPadState{keypad1: Pos{x: 0, y: 1}}
	ways := make(utils.HashSet[string])
	currentState := State{keypad1: A_DIR, keypad2: A_DIR}
	result := getControllerMoveset2(UP, currentState, &dirRobot2)
	for _, r := range result {
		newState := State{keypad1: r.dirPadState.keypad1, keypad2: r.dirPadState.keypad2}
		backToA := getControllerMoveset2(A_DIR, newState, &dirRobot2)
		for _, a := range backToA {
			ways.Insert(r.moveSet + a.moveSet + "A")
		}
	}

	for way := range ways {
		fmt.Println(way)
	}

	return total
	// 160800 too high
	// 125767 too low
}

func getNumPadNeighbourStates(numPad Keypad, current State, numPadRobot Robot) []ScoredState {
	neighbours := make([]ScoredState, 0)

	// UP
	if current.numPadPos.y > 0 {
		upMoves := getControllerMoveset2(UP, current, numPadRobot.controller)
		for _, move := range upMoves {
			nextPos := Pos{x: current.numPadPos.x, y: current.numPadPos.y - 1}
			nextState := State{numPadPos: nextPos, keypad2: move.dirPadState.keypad2, keypad1: move.dirPadState.keypad1}
			upScoredState := ScoredState{moveSet: move.moveSet, state: nextState}
			neighbours = append(neighbours, upScoredState) //numPad[current.numPadPos.y-1][current.numPadPos.x])
		}
	}

	// DOWN
	if current.numPadPos.y < len(numPad)-1 && numPad[current.numPadPos.y+1][current.numPadPos.x] != -1 {
		downMoves := getControllerMoveset2(UP, current, numPadRobot.controller)
		for _, move := range downMoves {
			nextPos := Pos{x: current.numPadPos.x, y: current.numPadPos.y + 1}
			nextState := State{numPadPos: nextPos, keypad2: move.dirPadState.keypad2, keypad1: move.dirPadState.keypad1}
			downScoredState := ScoredState{moveSet: move.moveSet, state: nextState}
			neighbours = append(neighbours, downScoredState) //numPad[current.numPadPos.y+1][current.numPadPos.x])

		}
	}

	// LEFT
	if current.numPadPos.x > 0 && numPad[current.numPadPos.y][current.numPadPos.x-1] != -1 {
		leftMoves := getControllerMoveset2(UP, current, numPadRobot.controller)
		for _, move := range leftMoves {
			nextPos := Pos{x: current.numPadPos.x - 1, y: current.numPadPos.y}
			nextState := State{numPadPos: nextPos, keypad2: move.dirPadState.keypad2, keypad1: move.dirPadState.keypad1}
			leftScoredState := ScoredState{moveSet: move.moveSet, state: nextState}
			neighbours = append(neighbours, leftScoredState) //numPad[current.numPadPos.y][current.numPadPos.x-1])
		}
	}

	// RIGHT
	if current.numPadPos.x < len(numPad[0])-1 {
		rightMoves := getControllerMoveset2(UP, current, numPadRobot.controller)
		for _, move := range rightMoves {
			nextPos := Pos{x: current.numPadPos.x + 1, y: current.numPadPos.y}
			nextState := State{numPadPos: nextPos, keypad2: move.dirPadState.keypad2, keypad1: move.dirPadState.keypad1}
			rightScoredState := ScoredState{moveSet: move.moveSet, state: nextState}
			neighbours = append(neighbours, rightScoredState) //numPad[current.numPadPos.y][current.numPadPos.x+1])
		}
	}

	return neighbours
}

func getControllerMoveset2(desiredPosition Pos, current State, robot *Robot) []ScoredDirPadState {
	if robot == nil {
		fmt.Println("calculate raw input here")
		return []ScoredDirPadState{}
		// return topLevelMove(current.keypad1, desiredPosition)
	}

	if current.keypad2 == desiredPosition {
		s := DirPadState{keypad1: current.keypad1, keypad2: current.keypad2, dir: current.dir}
		return []ScoredDirPadState{{s, ""}}
	}

	moves := make(map[DirPadState]string)
	//moveSet := ""
	// now pathfind on the dirPad

	queue := make(chan DirPadState, 5*5*4)
	startState := DirPadState{keypad2: current.keypad2, keypad1: current.keypad1, dir: UP}
	queue <- startState

	finalStates := make([]ScoredDirPadState, 0)

	for len(queue) > 0 {
		currentDpad := <-queue

		neighbours := getDpadNeighbours(robot, currentDpad)
		for _, neighbour := range neighbours {
			if oldMoves, ok := moves[neighbour.dirPadState]; ok {
				if len(neighbour.moveSet) < len(oldMoves) {
					moves[neighbour.dirPadState] = neighbour.moveSet
					queue <- neighbour.dirPadState
				}
			} else {
				moves[neighbour.dirPadState] = neighbour.moveSet
				queue <- neighbour.dirPadState
			}

			if neighbour.dirPadState.keypad2 == desiredPosition {
				finalStates = append(finalStates, neighbour)
			}
		}
	}

	// fmt.Println("DPAD2 Final States:")
	// fmt.Println(finalStates)

	return finalStates
}

func getDpadNeighbours(robot *Robot, dpadState DirPadState) []ScoredDirPadState {
	neighbours := make([]ScoredDirPadState, 0)

	// UP
	if dpadState.keypad2.y > 0 {
		upMoves := getControllerMoveset1(UP, dpadState, robot.controller)
		nextState := DirPadState{keypad1: UP, keypad2: Pos{x: dpadState.keypad2.x, y: dpadState.keypad2.y - 1}}
		upScoredState := ScoredDirPadState{moveSet: upMoves, dirPadState: nextState}
		neighbours = append(neighbours, upScoredState) //numPad[current.numPadPos.y-1][current.numPadPos.x])
	}

	// DOWN
	if dpadState.keypad2.y < len(robot.keypad)-1 && robot.keypad[dpadState.keypad2.y+1][dpadState.keypad2.x] != -1 {
		downMoves := getControllerMoveset1(DOWN, dpadState, robot.controller)
		nextState := DirPadState{keypad1: DOWN, keypad2: Pos{x: dpadState.keypad2.x, y: dpadState.keypad2.y + 1}}
		downScoredState := ScoredDirPadState{moveSet: downMoves, dirPadState: nextState}
		neighbours = append(neighbours, downScoredState) //numPad[current.numPadPos.y+1][current.numPadPos.x])
	}

	// LEFT
	if dpadState.keypad2.x > 0 && robot.keypad[dpadState.keypad2.y][dpadState.keypad2.x-1] != -1 {
		leftMoves := getControllerMoveset1(LEFT, dpadState, robot.controller)
		nextState := DirPadState{keypad1: LEFT, keypad2: Pos{x: dpadState.keypad2.x - 1, y: dpadState.keypad2.y}}
		leftScoredState := ScoredDirPadState{moveSet: leftMoves, dirPadState: nextState}
		neighbours = append(neighbours, leftScoredState) //numPad[current.numPadPos.y][current.numPadPos.x-1])
	}

	// RIGHT
	if dpadState.keypad2.x < len(robot.keypad[0])-1 {
		rightMoves := getControllerMoveset1(RIGHT, dpadState, robot.controller)
		nextState := DirPadState{keypad1: RIGHT, keypad2: Pos{x: dpadState.keypad2.x + 1, y: dpadState.keypad2.y}}
		rightScoredState := ScoredDirPadState{moveSet: rightMoves, dirPadState: nextState}
		neighbours = append(neighbours, rightScoredState) //numPad[current.numPadPos.y][current.numPadPos.x+1])
	}

	return neighbours
}

func getControllerMoveset1(desiredPosition Pos, current DirPadState, robot *Robot) string {
	// rawdog this one??
	if robot.controller != nil {
		panic("somtin gon areye ere")
	}

	var inputSequence string

	currentPos := current.keypad1
	xDiff := desiredPosition.x - currentPos.x
	yDiff := desiredPosition.y - currentPos.y

	for xDiff != 0 || yDiff != 0 {
		if xDiff > 0 {
			for xDiff != 0 {
				inputSequence += ">"
				xDiff--
			}
		}

		if yDiff > 0 {
			for yDiff != 0 {
				inputSequence += "v"
				yDiff--
			}
		}

		if xDiff < 0 {
			for xDiff != 0 {
				inputSequence += "<"
				xDiff++
			}
		}

		if yDiff < 0 {
			for yDiff != 0 {
				inputSequence += "^"
				yDiff++
			}
		}
	}

	return inputSequence + "A"
}

func part2(lines []string) int {
	return 0
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
