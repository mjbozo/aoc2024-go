package day13

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Button struct {
	xMove int
	yMove int
}

type Pos struct {
	x int
	y int
}

func Run() {
	input, err := utils.ReadInputRaw("day13/input.txt")
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

func part1(lines string) int {
	machines := strings.Split(lines, "\n\n")

	total := 0
	for _, machine := range machines {
		machineSpecs := strings.Split(machine, "\n")
		aNums := strings.Split(machineSpecs[0], ": ")[1]
		aDeltas := strings.Split(aNums, ", ")
		aChangeX, _ := strconv.Atoi(aDeltas[0][2:])
		aChangeY, _ := strconv.Atoi(aDeltas[1][2:])

		aButton := Button{xMove: aChangeX, yMove: aChangeY}

		bNums := strings.Split(machineSpecs[1], ": ")[1]
		bDeltas := strings.Split(bNums, ", ")
		bChangeX, _ := strconv.Atoi(bDeltas[0][2:])
		bChangeY, _ := strconv.Atoi(bDeltas[1][2:])

		bButton := Button{xMove: bChangeX, yMove: bChangeY}

		prizeLocation := strings.Split(machineSpecs[2], ": ")[1]
		labels := strings.Split(prizeLocation, ", ")
		prizeLocationX, _ := strconv.Atoi(labels[0][2:])
		prizeLocationY, _ := strconv.Atoi(labels[1][2:])
		prizePos := Pos{x: prizeLocationX, y: prizeLocationY}

		currentPos := Pos{x: 0, y: 0}

		currentTokens := 0
		gotPrize := false

		for currentPos.x < prizePos.x && currentPos.y < prizePos.y {
			if isOnGradient(prizePos, currentPos, bButton) {
				currentTokens += (prizePos.x - currentPos.x) / bButton.xMove
				gotPrize = true
				break
			}

			currentPos.x += aButton.xMove
			currentPos.y += aButton.yMove
			currentTokens += 3
		}

		if gotPrize {
			total += currentTokens
		}
	}

	return total
}

func part2(lines string) int {
	machines := strings.Split(lines, "\n\n")

	total := 0
	for _, machine := range machines {
		machineSpecs := strings.Split(machine, "\n")
		aNums := strings.Split(machineSpecs[0], ": ")[1]
		aDeltas := strings.Split(aNums, ", ")
		aChangeX, _ := strconv.Atoi(aDeltas[0][2:])
		aChangeY, _ := strconv.Atoi(aDeltas[1][2:])

		aButton := Button{xMove: aChangeX, yMove: aChangeY}

		bNums := strings.Split(machineSpecs[1], ": ")[1]
		bDeltas := strings.Split(bNums, ", ")
		bChangeX, _ := strconv.Atoi(bDeltas[0][2:])
		bChangeY, _ := strconv.Atoi(bDeltas[1][2:])

		bButton := Button{xMove: bChangeX, yMove: bChangeY}

		prizeLocation := strings.Split(machineSpecs[2], ": ")[1]
		labels := strings.Split(prizeLocation, ", ")
		prizeLocationX, _ := strconv.Atoi(labels[0][2:])
		prizeLocationY, _ := strconv.Atoi(labels[1][2:])
		prizePos := Pos{x: prizeLocationX + 10000000000000, y: prizeLocationY + 10000000000000}

		xMoveA := float64(aButton.xMove)
		yMoveA := float64(aButton.yMove)
		deltaA := yMoveA / xMoveA

		xMoveB := float64(bButton.xMove)
		yMoveB := float64(bButton.yMove)
		deltaB := yMoveB / xMoveB

		prizeX := float64(prizePos.x)
		prizeY := float64(prizePos.y)

		bInterceptY := prizeY - (prizeX * deltaB)
		xCross := bInterceptY / (deltaA - deltaB)

		aButtonPresses := round(xCross / xMoveA)
		bButtonPresses := round((prizeX - xCross) / xMoveB)

		if (aButtonPresses*aButton.xMove)+(bButtonPresses*bButton.xMove) == prizePos.x &&
			(aButtonPresses*aButton.yMove)+(bButtonPresses*bButton.yMove) == prizePos.y {
			total += aButtonPresses*3 + bButtonPresses
		}
	}

	return total
}

func isOnGradient(prizePos, currentPos Pos, button Button) bool {
	xDiff := prizePos.x - currentPos.x
	yDiff := prizePos.y - currentPos.y
	return xDiff%button.xMove == 0 && yDiff%button.yMove == 0 && (xDiff/button.xMove) == (yDiff/button.yMove)
}

func round(f float64) int {
	return int(math.Round(f))
}
