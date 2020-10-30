package qsort

import (
	"math/rand"
	"strings"
)

type Sorter interface {
	Swap(a int, b int)
	Child(start int, end int) Sorter
	Compare(a int, b int) int
	Len() int
}

func Quicksort(s Sorter) {
	length := s.Len()
	if length < 2 {
		return
	}

	left, right := 0, length-1
	pivot := rand.Int() % length
	s.Swap(pivot, right)

	for i := 0; i < length; i++ {
		if s.Compare(i, right) < 0 {
			s.Swap(left, i)
			left++
		}
	}
	s.Swap(left, right)

	Quicksort(s.Child(0, left))
	Quicksort(s.Child(left+1, length))

}

// Convenience functions

type stringSorter struct {
	slice []string
}

func (ss *stringSorter) Swap(a int, b int) {
	ss.slice[a], ss.slice[b] = ss.slice[b], ss.slice[a]
}
func (ss *stringSorter) Compare(a int, b int) int {
	return strings.Compare(ss.slice[a], ss.slice[b])
}
func (ss *stringSorter) Child(start int, end int) Sorter {
	return &stringSorter{ss.slice[start:end]}
}
func (ss *stringSorter) Len() int {
	return len(ss.slice)
}


func QuickSortStrings(slice []string)  []string{
	 ss := stringSorter{slice}
	 Quicksort(&ss)
	 return slice
}

type intSorter struct {
	slice []int
}

func (is *intSorter) Swap(a int, b int) {
	is.slice[a], is.slice[b] = is.slice[b], is.slice[a]
}
func (is *intSorter) Compare(a int, b int) int {
	return is.slice[a] - is.slice[b]
}
func (is *intSorter) Child(start int, end int) Sorter {
	return &intSorter{is.slice[start:end]}
}
func (is *intSorter) Len() int {
	return len(is.slice)
}

func QuickSortInts(slice []int)  []int{
	is := intSorter{slice}
	Quicksort(&is)
	return slice
}
