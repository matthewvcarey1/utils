package binchop

import (
	"strings"
)

type SliceElementComparitor interface {
	Cmp(model interface{}, index int) int
}

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


type stringSliceComp struct {
	slice []string
}

func (sc stringSliceComp) Cmp(model interface{}, index int) int {
	return strings.Compare(model.(string), sc.slice[index])
}

type intSliceComp struct {
	slice []int
}

func (ic intSliceComp) Cmp(model interface{}, index int) int {
	return model.(int) - ic.slice[index]
}


func BinChopInts(value int, slice []int) (bool, int){
	isc := intSliceComp{slice}
	return BinChop(value, 0, len(slice), isc)
}

func BinChopStrings(value string, slice []string) (bool, int){
	ssc := stringSliceComp{slice}
	return BinChop(value, 0, len(slice), ssc)
}