package day24

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Gate struct {
	left  string
	right string
	op    string
}

type GateCount struct {
	id    string
	count int
}

func Run() {
	input, err := utils.ReadInputRaw("day24/input.txt", 24)
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
	fmt.Printf("Part 2: %s (%v)\n", part2Result, elapsed)
}

func part1(lines string) int {
	sections := strings.Split(lines, "\n\n")
	initialValues := sections[0]
	connections := sections[1]

	gates := make(map[string]Gate)
	wireValues := make(map[string]int)

	zGates := make([]string, 0)

	for _, initial := range strings.Split(initialValues, "\n") {
		parts := strings.Split(initial, ": ")
		wire := parts[0]
		value, _ := strconv.Atoi(parts[1])
		wireValues[wire] = value
	}

	for _, connections := range strings.Split(connections, "\n") {
		components := strings.Split(connections, " ")
		left, op, right, id := components[0], components[1], components[2], components[4]

		gate := Gate{
			left:  left,
			right: right,
			op:    op,
		}

		gates[id] = gate

		if strings.HasPrefix(id, "z") {
			zGates = append(zGates, id)
		}
	}

	zValues := make([]int, len(zGates))
	for _, zGate := range zGates {
		bitIndex, _ := strconv.Atoi(zGate[1:])
		bitValue := calculateGateOutput(zGate, gates, wireValues)
		zValues[bitIndex] = bitValue
	}

	total := 0
	for i, val := range zValues {
		total += val * int(math.Pow(2.0, float64(i)))
	}

	return total
}

// im too smooth brain to figure it out myself, so approach stolen from here:
// https://www.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/
func part2(lines string) string {
	sections := strings.Split(lines, "\n\n")
	initialValues := sections[0]
	connections := sections[1]

	gates := make(map[string]Gate)
	wireValues := make(map[string]int)

	for _, initial := range strings.Split(initialValues, "\n") {
		parts := strings.Split(initial, ": ")
		wire := parts[0]
		value, _ := strconv.Atoi(parts[1])
		wireValues[wire] = value
	}

	inputMap := make(map[string][]Gate)

	for _, connections := range strings.Split(connections, "\n") {
		components := strings.Split(connections, " ")
		left, op, right, id := components[0], components[1], components[2], components[4]

		gate := Gate{
			left:  left,
			right: right,
			op:    op,
		}

		gates[id] = gate
		inputMap[left] = append(inputMap[left], gate)
		inputMap[right] = append(inputMap[right], gate)
	}

	faultyGates := getFaultyGates(gates, inputMap)
	faulty := make([]string, 0)
	for gate := range faultyGates {
		faulty = append(faulty, gate)
	}

	slices.Sort(faulty)

	return strings.Join(faulty, ",")
}

func calculateGateOutput(id string, gates map[string]Gate, wireValues map[string]int) int {
	if val, exists := wireValues[id]; exists {
		return val
	}

	gate := gates[id]
	leftValue := calculateGateOutput(gate.left, gates, wireValues)
	rightValue := calculateGateOutput(gate.right, gates, wireValues)

	var outputValue int
	switch gate.op {
	case "AND":
		outputValue = leftValue & rightValue
	case "OR":
		outputValue = leftValue | rightValue
	case "XOR":
		outputValue = leftValue ^ rightValue
	default:
		panic(fmt.Sprintf("Invalid operation: %s\n", gate.op))
	}

	wireValues[id] = outputValue
	return outputValue
}

func getFaultyGates(gates map[string]Gate, inputMap map[string][]Gate) utils.HashSet[string] {
	faulty := make(utils.HashSet[string])

	for k, v := range gates {
		// z output gates must be XOR, unless its the last bit z45
		if k[0] == 'z' && v.op != "XOR" && k != "z45" {
			faulty.Insert(k)
			continue
		}

		// if non-z gate is XOR, inputs must not be X and Y
		if k[0] != 'z' && v.left[0] != 'x' && v.left[0] != 'y' && v.right[0] != 'x' && v.right[0] != 'y' && v.op == "XOR" {
			faulty.Insert(k)
			continue
		}

		// if XOR gate has x and y inputs, there must be another XOR gate with this gate as input
		if v.op == "XOR" && ((v.left[0] == 'x' && v.right[0] == 'y') || (v.left[0] == 'y' && v.right[0] == 'x')) &&
			v.left != "x00" && v.right != "x00" && v.left != "y00" && v.right != "y00" {
			if _, ok := inputMap[k]; !ok {
				faulty.Insert(k)
				continue
			}

			isValid := false
			for _, possible := range inputMap[k] {
				if possible.op == "XOR" {
					isValid = true
					break
				}
			}

			if !isValid {
				faulty.Insert(k)
				continue
			}
		}

		// each AND gate must be used as input to an OR gate, otherwise AND gate is faulty
		if v.op == "AND" && ((v.left[0] == 'x' && v.right[0] == 'y') || (v.left[0] == 'y' && v.right[0] == 'x')) &&
			v.left != "x00" && v.right != "x00" && v.left != "y00" && v.right != "y00" {
			if _, ok := inputMap[k]; !ok {
				faulty.Insert(k)
				continue
			}

			isValid := false
			for _, possible := range inputMap[k] {
				if possible.op == "OR" {
					isValid = true
					break
				}

				if !isValid {
					faulty.Insert(k)
					continue
				}
			}
		}
	}

	return faulty
}
