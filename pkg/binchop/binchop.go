package binchop

import (
	"strings"
)

// An interface that describes the function neeed by BinChop
type SliceElementComparitor interface {
	Cmp(model interface{}, index int) int
}

// A type agnostic BinChop implementation
func BinChop(value interface{}, start int, end int, sec SliceElementComparitor) (bool, int) {
	for {
		size := end - start
		// Not found case
		if size < 1 {
			return false, -1
		}
		half := size >> 1

		which := sec.Cmp(value, start+half)
		// Found case
		if which == 0 {
			return true, start + half
		} else {
			if which > 0 {
				// Look in top half
				start += half + 1
			} else {
				// Look in bottom half
				end = start + half
			}
		}
	}
}

// struct used for searching sorted slices of strings
type stringSliceComp struct {
	slice []string
}

// Cmp compare the model string with one at an index of the slice
func (sc stringSliceComp) Cmp(model interface{}, index int) int {
	return strings.Compare(model.(string), sc.slice[index])
}

// struct used for seaching sorted slices of ints
type intSliceComp struct {
	slice []int
}

// Cmp compare the model int with one at an index of the slice
func (ic intSliceComp) Cmp(model interface{}, index int) int {
	return model.(int) - ic.slice[index]
}

// BinChopInts convenience function to search ordered slices of ints
func BinChopInts(value int, slice []int) (bool, int){
	isc := intSliceComp{slice}
	return BinChop(value, 0, len(slice), isc)
}

// BinChopInts convenience function to search ordered slices of strings
func BinChopStrings(value string, slice []string) (bool, int){
	ssc := stringSliceComp{slice}
	return BinChop(value, 0, len(slice), ssc)
}