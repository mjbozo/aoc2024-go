package day17

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Bit int

const (
	X Bit = iota
	ZERO
	ONE
)

func (b Bit) String() string {
	switch b {
	case ZERO:
		return "0"
	case ONE:
		return "1"
	default:
		return "X"
	}
}

type BitSet [64]Bit

func (bs BitSet) value() int {
	val := 0
	for i := range 64 {
		val += bitsetValue(bs[i]) * (1 << i)
	}
	return val
}

func (bs BitSet) String() string {
	output := "[ "
	val := 0
	for i := range 64 {
		output += bs[i].String() + " "
		val += bitsetValue(bs[i]) * (1 << i)
	}
	output += "] " + fmt.Sprintf("%d", val)
	return output
}

func Run() {
	input, err := utils.ReadInputRaw("day17/input.txt", 17)
	if err != nil {
		log.Fatalln(utils.Red(err.Error()))
	}

	start := time.Now()
	part1Result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %s (%v)\n", part1Result, elapsed)

	start = time.Now()
	part2Result := part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d (%v)\n", part2Result, elapsed)
}

func part1(input string) string {
	parts := strings.Split(input, "\n\n")
	registers := strings.Split(parts[0], "\n")
	program := strings.Split(parts[1], ": ")[1]

	A, _ := strconv.Atoi(strings.Split(registers[0], ": ")[1])
	B, _ := strconv.Atoi(strings.Split(registers[1], ": ")[1])
	C, _ := strconv.Atoi(strings.Split(registers[2], ": ")[1])
	instructions := strings.Split(program, ",")

	return evaluate(instructions, A, B, C)
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	program := strings.Split(parts[1], ": ")[1]
	instructions := strings.Split(program, ",")
	optionsA := make([]BitSet, 0)

	for s, instruction := range instructions {
		n, _ := strconv.Atoi(instruction)
		B := n ^ 4

		if len(optionsA) == 0 {
			for C := range 8 {
				bitset, validToAdd := createBitset(B, C)
				if validToAdd {
					optionsA = append(optionsA, bitset)
				}
			}
		} else {
			nextOptionsA := make([]BitSet, 0)

			for _, existingBitSet := range optionsA {
				for C := range 8 {
					bitsetToMerge, validToAdd := createBitset(B, C)
					if !validToAdd {
						continue
					}

					bitsetToMerge = leftShift(bitsetToMerge, s)
					for i := range 64 {
						if existingBitSet[i] != X && bitsetToMerge[i] != X && existingBitSet[i] != bitsetToMerge[i] {
							validToAdd = false
							break
						}

						if bitsetToMerge[i] == X {
							bitsetToMerge[i] = existingBitSet[i]
						}
					}

					if validToAdd {
						nextOptionsA = append(nextOptionsA, bitsetToMerge)
					}
				}
			}

			optionsA = nextOptionsA
		}
	}

	minimum := optionsA[0].value()
	for _, op := range optionsA {
		if op.value() < minimum {
			minimum = op.value()
		}
	}
	return minimum
}

func evaluate(instructions []string, A, B, C int) string {
	pc := 0
	var output string
	for pc < len(instructions) {
		out := process(instructions[pc], instructions[pc+1], &A, &B, &C, &pc)
		if len(out) > 0 {
			output += out + ","
		}
		pc += 2
	}

	return output[:len(output)-1]
}

func createBitset(B, C int) (BitSet, bool) {
	prevB := B ^ C
	lowerA := prevB ^ 7
	var bitset BitSet
	for i := range 3 {
		bitset[i] = intToBit(lowerA & 1)
		lowerA = lowerA >> 1
	}

	validToAdd := true
	for i := prevB; i < prevB+3; i++ {
		if bitset[i] == X {
			bitset[i] = intToBit(C & 1)
		} else if bitsetValue(bitset[i]) != C&1 {
			validToAdd = false
			break
		}
		C = C >> 1
	}

	return bitset, validToAdd
}

func leftShift(bs BitSet, offset int) BitSet {
	for i := len(bs) - 1; i >= 3*offset; i-- {
		bs[i] = bs[i-3*offset]
	}

	for i := 0; i < 3*offset; i++ {
		bs[i] = X
	}

	return bs
}

func bitsetValue(b Bit) int {
	switch b {
	case ONE:
		return 1
	case ZERO:
		return 0
	case X:
		return 0
	}

	return -1
}

func intToBit(x int) Bit {
	switch x {
	case 0:
		return ZERO
	case 1:
		return ONE
	default:
		return X
	}
}

func process(opcode, operand string, A, B, C, pc *int) string {
	switch opcode {
	case "0":
		num := *A
		x := combo(operand, *A, *B, *C)
		den := math.Pow(2.0, float64(x))
		val := int(math.Floor(float64(num) / den))
		*A = val

	case "1":
		x, _ := strconv.Atoi(operand)
		val := *B ^ x
		*B = val

	case "2":
		x := combo(operand, *A, *B, *C)
		val := x % 8
		*B = val

	case "3":
		if *A != 0 {
			x, _ := strconv.Atoi(operand)
			*pc = x - 2
		}

	case "4":
		val := *B ^ *C
		*B = val

	case "5":
		x := combo(operand, *A, *B, *C)
		val := x % 8
		return fmt.Sprintf("%d", val)

	case "6":
		num := *A
		x := combo(operand, *A, *B, *C)
		den := math.Pow(2.0, float64(x))
		val := int(math.Floor(float64(num) / den))
		*B = val

	case "7":
		num := *A
		x := combo(operand, *A, *B, *C)
		den := math.Pow(2.0, float64(x))
		val := int(math.Floor(float64(num) / den))
		*C = val
	}

	return ""
}

func combo(operand string, A, B, C int) int {
	switch operand {
	case "0", "1", "2", "3":
		x, _ := strconv.Atoi(operand)
		return x
	case "4":
		return A
	case "5":
		return B
	case "6":
		return C
	case "7":
		panic("Got combo value 7")
	}

	return 0
}
