package day11

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputRaw("day11/input.txt")
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

func part1(line string) int {
	stonesStr := strings.Split(line, " ")
	stones := make([]int, 0)
	for _, s := range stonesStr {
		num, _ := strconv.Atoi(s)
		stones = append(stones, num)
	}

	for range 25 {
		stones = blink(stones)
	}

	return len(stones)
}

func part2(line string) uint {
	stonesStr := strings.Split(line, " ")
	stones := make(map[uint]uint, 0)
	for _, s := range stonesStr {
		num, _ := strconv.Atoi(s)
		stones[uint(num)]++
	}

	memo := make(map[uint][]uint)
	var numStones uint = 0

	for range 75 {
		stones = blink2(&stones, &memo, &numStones)
	}

	var sum uint = 0
	for _, val := range stones {
		sum += val
	}

	return sum
}

func blink(stones []int) []int {
	nextStones := make([]int, 0)

	for _, stone := range stones {
		stoneStr := fmt.Sprintf("%d", stone)
		if stone == 0 {
			nextStones = append(nextStones, 1)
		} else if len(stoneStr)%2 == 0 {
			n := len(stoneStr) / 2
			first, _ := strconv.Atoi(stoneStr[:n])
			second, _ := strconv.Atoi(stoneStr[n:])
			nextStones = append(nextStones, first)
			nextStones = append(nextStones, second)
		} else {
			nextStones = append(nextStones, stone*2024)
		}
	}

	return nextStones
}

func blink2(stones *map[uint]uint, memo *map[uint][]uint, numStones *uint) map[uint]uint {
	// loop through stones map
	nextStones := make(map[uint]uint)
	for key, val := range *stones {
		nextStones[key] = val
	}

	for key, val := range *stones {
		if val == 0 {
			continue
		}

		// if num in memo, dw abt compute, otherwise compute
		if turnsInto, ok := (*memo)[key]; ok {
			// figure out how many stones it turns into
			// increase numStones
			*numStones += uint(len(turnsInto)) * val
			nextStones[key] -= val
			for _, nextStone := range turnsInto {
				nextStones[nextStone] += val
			}
		} else {
			// put in memo if not already
			stoneStr := fmt.Sprintf("%d", key)
			if key == 0 {
				nextStones[0] -= val
				nextStones[1] += val
				(*memo)[0] = []uint{1}
			} else if len(stoneStr)%2 == 0 {
				n := len(stoneStr) / 2
				first, _ := strconv.Atoi(stoneStr[:n])
				second, _ := strconv.Atoi(stoneStr[n:])
				nextStones[key] -= val
				nextStones[uint(first)] += val
				nextStones[uint(second)] += val
				(*memo)[key] = []uint{uint(first), uint(second)}
			} else {
				v := uint(key * 2024)
				nextStones[key] -= val
				nextStones[v] += val
				(*memo)[key] = []uint{v}
			}
		}
	}

	return nextStones
}
