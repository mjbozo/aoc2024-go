package day21

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"math"
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
)

type Pos struct {
	x int
	y int
}

type RobotChainState int64 // needs to store 5^25 states
type Robot struct {
	idx             int
	pos             Pos
	controller      *Robot
	prevRobotStates RobotChainState
	memo            map[utils.Pair[Pos, RobotChainState]]utils.Pair[int, RobotChainState]
}

func getState(robot Robot) RobotChainState {
	return RobotChainState(math.Pow(5.0, float64(robot.idx)) * float64(dpadIndex(robot.pos)))
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
	dpadRobots := make([]Robot, 0)
	numBots := 2

	for i := range numBots {
		newRobot := Robot{
			idx:             i,
			pos:             A_DIR,
			prevRobotStates: RobotChainState(dpadIndex(A_DIR)),
			memo:            make(map[utils.Pair[Pos, RobotChainState]]utils.Pair[int, RobotChainState]),
		}

		if i > 0 {
			newRobot.controller = &dpadRobots[i-1]
			newRobot.prevRobotStates = dpadRobots[i-1].prevRobotStates + getState(newRobot)
		}

		dpadRobots = append(dpadRobots, newRobot)
	}

	doorRobot := Robot{
		idx:        numBots,
		pos:        A_NUM,
		controller: &dpadRobots[len(dpadRobots)-1],
		memo:       make(map[utils.Pair[Pos, RobotChainState]]utils.Pair[int, RobotChainState]),
	}

	total := 0
	for _, line := range lines {
		var movesetLength int
		numeric, _ := strconv.Atoi(line[:len(line)-1])

		// for each number, figure out moves to get there
		for _, x := range line {
			movesetLength += passInstructionToDoorRobot(&doorRobot, x)
		}

		total += numeric * movesetLength
	}

	return total
}

func part2(lines []string) int {
	// notable property that after each iteration, robots will all be returned to starting positions
	dpadRobots := make([]Robot, 0)
	numBots := 25

	for i := range numBots {
		newRobot := Robot{
			idx:             i,
			pos:             A_DIR,
			prevRobotStates: RobotChainState(dpadIndex(A_DIR)),
			memo:            make(map[utils.Pair[Pos, RobotChainState]]utils.Pair[int, RobotChainState]),
		}

		if i > 0 {
			newRobot.controller = &dpadRobots[i-1]
			newRobot.prevRobotStates = dpadRobots[i-1].prevRobotStates + getState(newRobot)
		}

		dpadRobots = append(dpadRobots, newRobot)
	}

	doorRobot := Robot{
		idx:             numBots,
		pos:             A_NUM,
		controller:      &dpadRobots[len(dpadRobots)-1],
		prevRobotStates: dpadRobots[len(dpadRobots)-1].prevRobotStates,
		memo:            make(map[utils.Pair[Pos, RobotChainState]]utils.Pair[int, RobotChainState]),
	}

	total := 0
	for _, line := range lines {
		var movesetLength int
		numeric, _ := strconv.Atoi(line[:len(line)-1])

		// for each number, figure out moves to get there
		for _, x := range line {
			movesetLength += passInstructionToDoorRobot(&doorRobot, x)
		}

		total += numeric * movesetLength
	}

	return total
}

func passInstructionToDoorRobot(doorRobot *Robot, x rune) int {
	var movesetLength int
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

	// need to compare vertical first and horizontal first
	// but also need to check if either will run into the gap
	// weird pointer shit needed here to prevent overwriting data behind memory addresses
	horizontalFirstController := &Robot{}
	verticalFirstController := &Robot{}
	*horizontalFirstController = *doorRobot.controller
	*verticalFirstController = *doorRobot.controller

	var horizontalFirstMovesetLength int
	horizontalHitsGap := doorRobot.pos.y == GAP_NUM.y && doorRobot.pos.x+xDiff == GAP_NUM.x
	if !horizontalHitsGap {
		hThenV := make([]Pos, 0)
		hThenV = append(hThenV, horizontalMoves...)
		hThenV = append(hThenV, verticalMoves...)
		horizontalFirstMovesetLength, _ = generateMoveset(hThenV, horizontalFirstController)
		aMovesLength, _ := generateMoveset([]Pos{A_DIR}, horizontalFirstController)
		horizontalFirstMovesetLength += aMovesLength
	}

	var verticalFirstMovesetLength int
	verticalHitsGap := doorRobot.pos.x == GAP_NUM.x && doorRobot.pos.y+yDiff == GAP_NUM.y
	if !verticalHitsGap {
		vThenH := make([]Pos, 0)
		vThenH = append(vThenH, verticalMoves...)
		vThenH = append(vThenH, horizontalMoves...)
		verticalFirstMovesetLength, _ = generateMoveset(vThenH, verticalFirstController)
		aMovesLength, _ := generateMoveset([]Pos{A_DIR}, verticalFirstController)
		verticalFirstMovesetLength += aMovesLength
	}

	// set moveset accordingly
	if horizontalHitsGap {
		movesetLength += verticalFirstMovesetLength
		doorRobot.controller = verticalFirstController
	} else if verticalHitsGap {
		movesetLength += horizontalFirstMovesetLength
		doorRobot.controller = horizontalFirstController
	} else {
		if horizontalFirstMovesetLength <= verticalFirstMovesetLength {
			movesetLength += horizontalFirstMovesetLength
			doorRobot.controller = horizontalFirstController
		} else {
			movesetLength += verticalFirstMovesetLength
			doorRobot.controller = verticalFirstController
		}
	}

	doorRobot.pos = c
	doorRobot.prevRobotStates += getState(*doorRobot)

	return movesetLength
}

// Given a set of moves requested by next robot, figure out best moves to do those actions
func generateMoveset(desiredMoves []Pos, robot *Robot) (int, RobotChainState) {
	if robot == nil {
		return len(getHumanInput(desiredMoves)), 0
	}

	var movesetLength int

	// move is the button THIS ROBOT wants to press
	for _, move := range desiredMoves {
		robotOriginalState := robot.prevRobotStates
		if existingResult, ok := robot.memo[utils.Pair[Pos, RobotChainState]{First: move, Second: robot.prevRobotStates}]; ok {
			// seen this exact move and state before
			movesetLength += existingResult.First

			robot.pos = move

			// need to update all the robot states still
			currentRobot := robot.controller
			controllerRobots := make([]*Robot, 0)
			for currentRobot != nil {
				controllerRobots = append(controllerRobots, currentRobot)
				currentRobot = currentRobot.controller
			}

			for i := len(controllerRobots) - 1; i >= 0; i-- {
				if i == len(controllerRobots)-1 {
					controllerRobots[i].prevRobotStates = getState(*controllerRobots[i])
				} else {
					controllerRobots[i].prevRobotStates = controllerRobots[i+1].prevRobotStates + getState(*controllerRobots[i])
				}
			}

			var controllerState RobotChainState
			if robot.controller != nil {
				controllerState = robot.controller.prevRobotStates
			}
			robot.prevRobotStates = getState(*robot) + controllerState

			continue
		}

		xDiff := move.x - robot.pos.x
		yDiff := move.y - robot.pos.y
		horizontalMoves := make([]Pos, 0, utils.Abs(xDiff))
		verticalMoves := make([]Pos, 0, utils.Abs(yDiff))

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

		// need to compare vertical first and horizontal first
		// but also need to check if either will run into the gap
		// weird pointer shit needed here to prevent overwriting data behind memory addresses
		horizontalFirstController := &Robot{}
		verticalFirstController := &Robot{}
		if robot.controller != nil {
			*horizontalFirstController = *robot.controller
			*verticalFirstController = *robot.controller
		} else {
			horizontalFirstController = nil
			verticalFirstController = nil
		}

		var horizontalFirstMovesetLength int
		var horizontalFirstNewState RobotChainState
		var aMovesLength int
		horizontalHitsGap := robot.pos.y == GAP_DIR.y && robot.pos.x+xDiff == GAP_DIR.x
		if !horizontalHitsGap {
			hThenV := make([]Pos, 0, len(horizontalMoves)+len(verticalMoves))
			hThenV = append(hThenV, horizontalMoves...)
			hThenV = append(hThenV, verticalMoves...)
			horizontalFirstMovesetLength, horizontalFirstNewState = generateMoveset(hThenV, horizontalFirstController)
			aMovesLength, horizontalFirstNewState = generateMoveset([]Pos{A_DIR}, horizontalFirstController)
			horizontalFirstMovesetLength += aMovesLength
		}

		var verticalFirstMovesetLength int
		var verticalFirstNewState RobotChainState
		verticalHitsGap := robot.pos.x == GAP_DIR.x && robot.pos.y+yDiff == GAP_DIR.y
		if !verticalHitsGap {
			vThenH := make([]Pos, 0, len(verticalMoves)+len(horizontalMoves))
			vThenH = append(vThenH, verticalMoves...)
			vThenH = append(vThenH, horizontalMoves...)
			verticalFirstMovesetLength, verticalFirstNewState = generateMoveset(vThenH, verticalFirstController)
			aMovesLength, verticalFirstNewState = generateMoveset([]Pos{A_DIR}, verticalFirstController)
			verticalFirstMovesetLength += aMovesLength
		}

		// set moveset accordingly
		var newMovesLength int
		if horizontalHitsGap {
			newMovesLength = verticalFirstMovesetLength
			robot.controller = verticalFirstController
			robot.prevRobotStates = verticalFirstNewState
		} else if verticalHitsGap {
			newMovesLength = horizontalFirstMovesetLength
			robot.controller = horizontalFirstController
			robot.prevRobotStates = horizontalFirstNewState
		} else {
			if horizontalFirstMovesetLength <= verticalFirstMovesetLength {
				newMovesLength = horizontalFirstMovesetLength
				robot.controller = horizontalFirstController
				robot.prevRobotStates = horizontalFirstNewState
			} else {
				newMovesLength = verticalFirstMovesetLength
				robot.controller = verticalFirstController
				robot.prevRobotStates = verticalFirstNewState
			}
		}

		robot.pos = move
		robot.prevRobotStates += getState(*robot)

		// memoise the state for later
		memoKey := utils.Pair[Pos, RobotChainState]{First: move, Second: robotOriginalState}
		memoValue := utils.Pair[int, RobotChainState]{First: newMovesLength, Second: robot.prevRobotStates}
		robot.memo[memoKey] = memoValue

		movesetLength += newMovesLength
	}

	return movesetLength, robot.prevRobotStates
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

func dpadIndex(pos Pos) int {
	switch pos {
	case LEFT:
		return 0
	case DOWN:
		return 1
	case RIGHT:
		return 2
	case UP:
		return 3
	case A_DIR:
		return 4
	}

	return 0
}
