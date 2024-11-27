package utils

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadInput(filename string) ([]string, error) {
	return ReadInputByDelim(filename, "\n")
}

func ReadInputByDelim(filename, delim string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input_string := strings.TrimSpace(string(bytes))
	segments := strings.Split(input_string, delim)

	return segments, nil
}

func ReadInputRaw(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	input_string := strings.TrimSpace(string(bytes))
	return input_string, nil
}

func Manhattan(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(y2-y1)) + math.Abs(float64(x2-x1)))
}

func FindLCM(x, y int) int {
	largest := x
	if y > x {
		largest = y
	}

	upperBound := x * y
	currentLCM := largest

	for currentLCM <= upperBound {
		if currentLCM%x == 0 && currentLCM%y == 0 {
			break
		}
		currentLCM += largest
	}

	return currentLCM
}

func Overlaps(x1, y1, x2, y2 int) bool {
	return x1 <= y2 && x2 <= y1
}

func Filter[S ~[]E, E any](s S, predicate func(E) bool) S {
	filtered := make([]E, 0)
	for _, v := range s {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func Red(input string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", input)
}
