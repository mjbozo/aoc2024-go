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

// A Set, implemented as a wrapper around Go's map type
type HashSet[T comparable] map[T]bool

func (s HashSet[T]) String() string {
	var output string
	for k := range s {
		if len(output) > 0 {
			output += ", "
		}
		output += fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("[%s]", output)
}

// Checks if key exists in the set
func (s *HashSet[T]) Contains(key T) bool {
	_, exists := (*s)[key]
	return exists
}

// Inserts a key in the set
//
// Returns true if the key is newly inserted into the set
func (s *HashSet[T]) Insert(key T) bool {
	exists := s.Contains(key)
	(*s)[key] = true
	return !exists
}

// Removes a key from the set
//
// Returns true if the key was successfully removed
func (s *HashSet[T]) Remove(key T) bool {
	if s.Contains(key) {
		delete(*s, key)
		return true
	}

	return false
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

// Node data structure for binary heap
type heapNode[T comparable] struct {
	value  T
	left   *heapNode[T]
	right  *heapNode[T]
	parent *heapNode[T]
}

// Default formatting for binary heap node
func (n heapNode[T]) String() string {
	return fmt.Sprintf("HeapNode[Val=%v, Left=%v, Right=%v]", n.value, n.left, n.right)
}

// Helper method for swapping two node values, and returning the new swapped node
func (n *heapNode[T]) swap(other *heapNode[T]) *heapNode[T] {
	temp := n.value
	n.value = other.value
	other.value = temp
	n = other
	return n
}

// Error for Binary Heap operations
type HeapError struct {
	msg string
}

// Error implementation for HeapError
func (e *HeapError) Error() string {
	return e.msg
}

// Binary heap data structure
//
// Can be used as either min-heap or max-heap depending in the user-defined comparator function
type BinaryHeap[T comparable] struct {
	Comparator func(a, b T) int
	root       *heapNode[T]
	size       int
}

// Default format for binary heap
func (h BinaryHeap[T]) String() string {
	return fmt.Sprintf("BinaryHeap[Size=%d, Root=%v]", h.size, h.root)
}

// Array representation of binary heap
func (h BinaryHeap[T]) Array() []T {
	arr := make([]T, 0)
	queue := make([]*heapNode[T], 0)
	queue = append(queue, h.root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			arr = append(arr, node.value)
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}
	}

	return arr
}

// Helper method for maintaining valid heap structure
//
// Swaps node with parent as long as node value is greater than parent value, as determined by heap comparator
func (h *BinaryHeap[T]) shiftUp(node *heapNode[T]) {
	for node.parent != nil && h.Comparator(node.value, node.parent.value) > 0 {
		node = node.swap(node.parent)
	}
}

// Helper method for maintaining valid heap structure
//
// Swaps node with largest child as long as node value is less than child value, as determined by heap comparator
func (h *BinaryHeap[T]) shiftDown(node *heapNode[T]) {
	for (node.left != nil && h.Comparator(node.value, node.left.value) < 0) || (node.right != nil && h.Comparator(node.value, node.right.value) < 0) {
		if node.right != nil {
			// then i should have left node too
			if h.Comparator(node.left.value, node.right.value) >= 0 {
				node = node.swap(node.left)
			} else {
				node = node.swap(node.right)
			}
		} else {
			node = node.swap(node.left)
		}
	}
}

// Insert a value into binary heap
func (h *BinaryHeap[T]) Insert(val T) {
	defer func() { h.size += 1 }()

	// insert new node into left-most vacant spot in last layer. find spot by BFS I guess?
	if h.root == nil {
		h.root = &heapNode[T]{value: val}
		return
	}

	var insertedNode *heapNode[T]
	queue := make([]*heapNode[T], 0)
	queue = append(queue, h.root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.left == nil {
			node.left = &heapNode[T]{value: val, parent: node}
			insertedNode = node.left
			break
		}

		if node.right == nil {
			node.right = &heapNode[T]{value: val, parent: node}
			insertedNode = node.right
			break
		}

		queue = append(queue, node.left)
		queue = append(queue, node.right)
	}

	// then shift up until correctly placed
	h.shiftUp(insertedNode)
}

// Pop the element with highest priority from the heap and return pointer to the value
func (h *BinaryHeap[T]) Pop() (*T, error) {
	if h == nil || h.root == nil {
		return nil, &HeapError{msg: "Cannot pop empty or nil heap"}
	}

	defer func() {
		if h != nil {
			h.size -= 1
		}
	}()

	maxValue := h.root.value

	// find right-most element in last layer
	last := h.root
	queue := make([]*heapNode[T], 0)
	queue = append(queue, h.root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.left == nil {
			break
		}
		last = node.left

		if node.right == nil {
			break
		}
		last = node.right

		queue = append(queue, node.left)
		queue = append(queue, node.right)
	}

	// root is the only value in the heap
	if last == h.root {
		fmt.Println(h, maxValue)
		h.root = nil
		return &maxValue, nil
	}

	h.root.value = last.value
	if last.parent.right != nil {
		last.parent.right = nil
	} else {
		last.parent.left = nil
	}

	h.shiftDown(h.root)

	return &maxValue, nil
}

// Return the highest priority value in the heap, without removing it
func (h *BinaryHeap[T]) Peek() (*T, error) {
	if h == nil || h.root == nil {
		return nil, &HeapError{msg: "Cannot peek empty or nil heap"}
	}

	return &h.root.value, nil
}

// Read input from specified filename, separated by newlines
func ReadInput(filename string) ([]string, error) {
	return ReadInputByDelim(filename, "\r\n")
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

// Green text colour
func Green(input string) string {
	return fmt.Sprintf("\033[32m%s\033[39m", input)
}

// Red text colour
func Red(input string) string {
	return fmt.Sprintf("\033[31m%s\033[39m", input)
}
