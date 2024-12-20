package day16

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"slices"
	"time"
)

type Cell struct {
	pos Pos
	val byte
}

func (c Cell) String() string {
	return string(c.val)
}

type Pos struct {
	x int
	y int
}

type State struct {
	pos     Pos
	heading byte
}

type ScoredState struct {
	state State
	score int
}

func Run() {
	input, err := utils.ReadInput("day16/input.txt", 16)
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
	startPos := State{heading: '>'}
	endPos := Pos{}

	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			row = append(row, Cell{pos: Pos{x: x, y: y}, val: byte(c)})
			if c == 'S' {
				startPos.pos.x = x
				startPos.pos.y = y
			}

			if c == 'E' {
				endPos.x = x
				endPos.y = y
			}
		}
		grid = append(grid, row)
	}

	scores := make(map[State]int)
	queue := make([]State, 0)
	queue = append(queue, startPos)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		neighbours := getNeighbours(&grid, &current, &scores)
		for _, neighbour := range neighbours {
			if prevScore, exists := scores[neighbour.state]; exists {
				// compare
				if neighbour.score < prevScore {
					scores[neighbour.state] = neighbour.score
					queue = append(queue, neighbour.state)
				}
			} else {
				scores[neighbour.state] = neighbour.score
				queue = append(queue, neighbour.state)
			}
		}
	}

	upEnd := State{pos: endPos, heading: '^'}
	downEnd := State{pos: endPos, heading: 'v'}
	leftEnd := State{pos: endPos, heading: '<'}
	rightEnd := State{pos: endPos, heading: '>'}

	endScores := make([]int, 0)
	if s, ok := scores[upEnd]; ok {
		endScores = append(endScores, s)
	}
	if s, ok := scores[downEnd]; ok {
		endScores = append(endScores, s)
	}
	if s, ok := scores[leftEnd]; ok {
		endScores = append(endScores, s)
	}
	if s, ok := scores[rightEnd]; ok {
		endScores = append(endScores, s)
	}
	slices.Sort(endScores)

	return endScores[0]
}

func part2(lines []string) int {
	grid := make(utils.Grid[Cell], 0)
	startPos := State{heading: '>'}
	endPos := Pos{}

	for y, line := range lines {
		row := make([]Cell, 0)
		for x, c := range line {
			row = append(row, Cell{pos: Pos{x: x, y: y}, val: byte(c)})
			if c == 'S' {
				startPos.pos.x = x
				startPos.pos.y = y
			}

			if c == 'E' {
				endPos.x = x
				endPos.y = y
			}
		}
		grid = append(grid, row)
	}

	scores := make(map[State]int)
	scores[startPos] = 0

	queue := make([]State, 0)
	queue = append(queue, startPos)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		neighbours := getNeighbours(&grid, &current, &scores)
		for _, neighbour := range neighbours {
			if prevScore, exists := scores[neighbour.state]; exists {
				// compare
				if neighbour.score < prevScore {
					scores[neighbour.state] = neighbour.score
					queue = append(queue, neighbour.state)
				}
			} else {
				scores[neighbour.state] = neighbour.score
				queue = append(queue, neighbour.state)
			}
		}
	}

	upEnd := State{pos: endPos, heading: '^'}
	downEnd := State{pos: endPos, heading: 'v'}
	leftEnd := State{pos: endPos, heading: '<'}
	rightEnd := State{pos: endPos, heading: '>'}

	endScores := make([]int, 0)
	if s, ok := scores[upEnd]; ok {
		endScores = append(endScores, s)
	}
	if s, ok := scores[downEnd]; ok {
		endScores = append(endScores, s)
	}
	if s, ok := scores[leftEnd]; ok {
		endScores = append(endScores, s)
	}
	if s, ok := scores[rightEnd]; ok {
		endScores = append(endScores, s)
	}
	slices.Sort(endScores)

	bestScore := endScores[0]

	bestPositions := make(utils.HashSet[Pos])
	bestPositions.Insert(startPos.pos)
	bestPositions.Insert(endPos)

	endStates := getPositionStates(scores, endPos.x, endPos.y)
	for _, end := range endStates {
		if end.score == bestScore {
			queue = append(queue, end.state)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		neighbours := getNeighboursBacktrack(&grid, &current, &scores)
		for _, neighbour := range neighbours {
			bestPositions.Insert(neighbour.state.pos)
			queue = append(queue, neighbour.state)
		}
	}

	return len(bestPositions)
}

func getPositionStates(scores map[State]int, x, y int) []ScoredState {
	states := make([]ScoredState, 0)

	upState := State{pos: Pos{x: x, y: y}, heading: '^'}
	downState := State{pos: Pos{x: x, y: y}, heading: 'v'}
	leftState := State{pos: Pos{x: x, y: y}, heading: '<'}
	rightState := State{pos: Pos{x: x, y: y}, heading: '>'}

	if v, ok := scores[upState]; ok {
		states = append(states, ScoredState{state: upState, score: v})
	}
	if v, ok := scores[downState]; ok {
		states = append(states, ScoredState{state: downState, score: v})
	}
	if v, ok := scores[leftState]; ok {
		states = append(states, ScoredState{state: leftState, score: v})
	}
	if v, ok := scores[rightState]; ok {
		states = append(states, ScoredState{state: rightState, score: v})
	}
	return states

}

func getNeighbours(grid *utils.Grid[Cell], current *State, scores *map[State]int) []ScoredState {
	neighbours := make([]ScoredState, 0)
	currentPos := current.pos

	if (*grid)[currentPos.y-1][currentPos.x].val != '#' {
		score := (*scores)[*current] + 1
		if current.heading != '^' {
			score += 1000
		}
		newState := State{pos: Pos{x: currentPos.x, y: currentPos.y - 1}, heading: '^'}
		upNeighbour := ScoredState{state: newState, score: score}
		neighbours = append(neighbours, upNeighbour)
	}

	if (*grid)[currentPos.y+1][currentPos.x].val != '#' {
		score := (*scores)[*current] + 1
		if current.heading != 'v' {
			score += 1000
		}
		newState := State{pos: Pos{x: currentPos.x, y: currentPos.y + 1}, heading: 'v'}
		downNeighbour := ScoredState{state: newState, score: score}
		neighbours = append(neighbours, downNeighbour)
	}

	if (*grid)[currentPos.y][currentPos.x-1].val != '#' {
		score := (*scores)[*current] + 1
		if current.heading != '<' {
			score += 1000
		}
		newState := State{pos: Pos{x: currentPos.x - 1, y: currentPos.y}, heading: '<'}
		leftNeighbour := ScoredState{state: newState, score: score}
		neighbours = append(neighbours, leftNeighbour)
	}

	if (*grid)[currentPos.y][currentPos.x+1].val != '#' {
		score := (*scores)[*current] + 1
		if current.heading != '>' {
			score += 1000
		}
		newState := State{pos: Pos{x: currentPos.x + 1, y: currentPos.y}, heading: '>'}
		rightNeighbour := ScoredState{state: newState, score: score}
		neighbours = append(neighbours, rightNeighbour)
	}

	return neighbours
}

func getNeighboursBacktrack(grid *utils.Grid[Cell], current *State, scores *map[State]int) []ScoredState {
	neighbours := make([]ScoredState, 0)
	currentPos := current.pos
	currentScore := (*scores)[*current]

	if (*grid)[currentPos.y-1][currentPos.x].val != '#' {
		upStates := getPositionStates(*scores, currentPos.x, currentPos.y-1)
		for _, s := range upStates {
			if s.state.heading == 'v' && s.score == currentScore-1 {
				neighbours = append(neighbours, s)
			} else if s.state.heading != 'v' && s.score == currentScore-1001 {
				neighbours = append(neighbours, s)
			}
		}
	}

	if (*grid)[currentPos.y+1][currentPos.x].val != '#' {
		downStates := getPositionStates(*scores, currentPos.x, currentPos.y+1)
		for _, s := range downStates {
			if s.state.heading == '^' && s.score == currentScore-1 {
				neighbours = append(neighbours, s)
			} else if s.state.heading != '^' && s.score == currentScore-1001 {
				neighbours = append(neighbours, s)
			}
		}
	}

	if (*grid)[currentPos.y][currentPos.x-1].val != '#' {
		leftStates := getPositionStates(*scores, currentPos.x-1, currentPos.y)
		for _, s := range leftStates {
			if s.state.heading == '>' && s.score == currentScore-1 {
				neighbours = append(neighbours, s)
			} else if s.state.heading != '>' && s.score == currentScore-1001 {
				neighbours = append(neighbours, s)
			}
		}
	}

	if (*grid)[currentPos.y][currentPos.x+1].val != '#' {
		rightStates := getPositionStates(*scores, currentPos.x+1, currentPos.y)
		for _, s := range rightStates {
			if s.state.heading == '<' && s.score == currentScore-1 {
				neighbours = append(neighbours, s)
			} else if s.state.heading != '<' && s.score == currentScore-1001 {
				neighbours = append(neighbours, s)
			}
		}
	}

	return neighbours
}
