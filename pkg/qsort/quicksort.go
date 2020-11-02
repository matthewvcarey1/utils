package qsort

import (
	"math/rand"
	"strings"
)

// An interface of the object passed to Quicksort
type Sorter interface {
	Swap(a int, b int)
	Child(start int, end int) Sorter
	Compare(a int, b int) int
	Len() int
}

// A type agnostic implentation of Quicksort
func Quicksort(s Sorter, ascending bool) {
	length := s.Len()
	if length < 2 {
		return
	}

	left, right := 0, length-1
	pivot := rand.Int() % length
	s.Swap(pivot, right)

	for i := 0; i < length; i++ {
		if (ascending && s.Compare(i, right) < 0) || (!ascending && s.Compare(i, right) > 0) {
			s.Swap(left, i)
			left++
		}
	}
	s.Swap(left, right)

	Quicksort(s.Child(0, left), ascending)
	Quicksort(s.Child(left+1, length), ascending)

}

// Convenience functions and types

// The structure need for string sorting
type stringSorter struct {
	slice []string
}

// Swap swaps two strings in a slice
func (ss *stringSorter) Swap(a int, b int) {
	ss.slice[a], ss.slice[b] = ss.slice[b], ss.slice[a]
}

// Compare two strings in a slice
func (ss *stringSorter) Compare(a int, b int) int {
	return strings.Compare(ss.slice[a], ss.slice[b])
}

// Child returns a stringSorter pointer with a sub slice
func (ss *stringSorter) Child(start int, end int) Sorter {
	return &stringSorter{ss.slice[start:end]}
}

// Len returns the stringSorter's slice length
func (ss *stringSorter) Len() int {
	return len(ss.slice)
}

// QuickSortStrings a convenience fuction for sorting a slice of strings
func QuickSortStrings(slice []string, ascending bool) []string {
	ss := stringSorter{slice}
	Quicksort(&ss, ascending)
	return slice
}

// A structure for int sorting
type intSorter struct {
	slice []int
}

// Swap swaps two ints in a slice
func (is *intSorter) Swap(a int, b int) {
	is.slice[a], is.slice[b] = is.slice[b], is.slice[a]
}

// Compare two ints in a slice
func (is *intSorter) Compare(a int, b int) int {
	return is.slice[a] - is.slice[b]
}

// Child returns an intSorter pointer with a sub slice
func (is *intSorter) Child(start int, end int) Sorter {
	return &intSorter{is.slice[start:end]}
}

// Len returns the intSorter's slice length
func (is *intSorter) Len() int {
	return len(is.slice)
}

// QuickSortStrings a convenience fuction for sorting a slice of ints
func QuickSortInts(slice []int, ascending bool) []int {
	is := intSorter{slice}
	Quicksort(&is, ascending)
	return slice
}
