package utils

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// Interface for all number types, useful for generics
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Data structure for pairing two units of data
//
// First and Second fields do not need to be of the same type
type Pair[T, U any] struct {
	First  T
	Second U
}

// Default formatting for Pair datatype
//
// Returns string as "(First, Second)"
func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// Data structure for grouping three units of data
//
// First, Second, and Third fields do not need to be of the same type
type Triple[T, U, V any] struct {
	First  T
	Second U
	Third  V
}

// Default formatting for Triple datatype
//
// Returns string as "(First, Second, Third)"
func (t Triple[T, U, V]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", t.First, t.Second, t.Third)
}

// Data structure for representing a range of integers
type intRange struct {
	First  int
	Second int
}

// IntRange constructor, ensures First is always lower bound, and Second is always upper bound
func IntRange(f, s int) intRange {
	lower := f
	upper := s
	if f > s {
		lower = s
		upper = f
	}
	return intRange{First: lower, Second: upper}
}

// Read input from specified filename, separated by newlines
func ReadInput(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input_string := strings.TrimSpace(string(bytes))
	input_string_normalised := strings.ReplaceAll(input_string, "\r\n", "\n")
	segments := strings.Split(input_string_normalised, "\n")

	return segments, nil
}

// Read input from specific files, separated by delim parameter
func ReadInputByDelim(filename, delim string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input_string := strings.TrimSpace(string(bytes))
	segments := strings.Split(input_string, delim)

	return segments, nil
}

// Read input from specified filename, unseparated
func ReadInputRaw(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	input_string := strings.TrimSpace(string(bytes))
	return input_string, nil
}

// Calculate manhattan distance between two points
func Manhattan[T, U Number](p1, p2 Pair[T, U]) int {
	return int(math.Abs(float64(p2.Second-p1.Second)) + math.Abs(float64(p2.First-p1.First)))
}

// Calculate lowest common multilpe of two integers
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

// Returns true if two integer ranges are overlapping
func Overlaps(r1, r2 intRange) bool {
	return r1.First <= r2.Second && r2.First <= r1.Second
}

// Filter a slice based on a predicate function
//
// Returned slice will contain elements from original slice where predicate returned true
func Filter[S ~[]E, E any](s S, predicate func(E) bool) S {
	filtered := make([]E, 0)
	for _, v := range s {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Maps each element in slice according to the transformation function
func Map[S ~[]E, E any, T any](s S, transformer func(E) T) []T {
	mapped := make([]T, 0)
	for _, v := range s {
		mapped = append(mapped, transformer(v))
	}
	return mapped
}

// Green text colour
func Green(input string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", input)
}

// Red text colour
func Red(input string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", input)
}
