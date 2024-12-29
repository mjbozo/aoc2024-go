package day25

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

func Run() {
	input, err := utils.ReadInputRaw("day25/input.txt", 25)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2()
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %s (%v)\n", part2Result, elapsed)
}

type Key = []int
type Lock = []int

func part1(input string) int {
	keys := make([]Key, 0)
	locks := make([]Lock, 0)
	keysAndLocks := strings.Split(input, "\n\n")

	for _, keyOrLock := range keysAndLocks {
		input := strings.Split(keyOrLock, "\n")
		if input[0] == "....." {
			// key
			key := make(Key, 0)
			for i := range 5 {
				filled := 0
				for _, col := range input {
					if col[i] == '#' {
						filled++
					}
				}
				key = append(key, filled-1)
			}

			keys = append(keys, key)
		} else {
			// lock
			lock := make(Lock, 0)
			for i := range 5 {
				filled := 0
				for _, col := range input {
					if col[i] == '#' {
						filled++
					}
				}
				lock = append(lock, filled-1)
			}

			locks = append(locks, lock)
		}
	}

	possiblePairs := 0
	for _, key := range keys {
		for _, lock := range locks {
			isValid := true
			for i := range 5 {
				if key[i]+lock[i] > 5 {
					isValid = false
					break
				}
			}

			if isValid {
				possiblePairs++
			}
		}
	}

	return possiblePairs
}

func part2() string {
	return "Delivered the chronicle"
}

