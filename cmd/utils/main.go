package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/matthewvcarey1/utils/pkg/binchop"
	"github.com/matthewvcarey1/utils/pkg/qsort"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

type stringSorter struct {
	slice []string
}

func (ss *stringSorter) Swap(a int, b int) {
	ss.slice[a], ss.slice[b] = ss.slice[b], ss.slice[a]
}
func (ss *stringSorter) Compare(a int, b int) int {
	return strings.Compare(ss.slice[a], ss.slice[b])
}
func (ss *stringSorter) Child(start int, end int) qsort.Sorter {
	return &stringSorter{ss.slice[start:end]}
}
func (ss *stringSorter) Len() int {
	return len(ss.slice)
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
func (is *intSorter) Child(start int, end int) qsort.Sorter {
	return &intSorter{is.slice[start:end]}
}
func (is *intSorter) Len() int {
	return len(is.slice)
}


func pickStringValueInSlice(s []string) string {
	index := rand.Intn(len(s) - 1)
	return s[index]
}

func pickValueInSlice(s []int) int {
	index := rand.Intn(len(s) - 1)
	return s[index]
}

func pickValueNotInSlice(s []int) int {
	for {
		v := rand.Intn(999) - rand.Intn(999)
		for _, c := range s {
			if c == v {
				break
			}
			if c > v {
				fmt.Println(c, ">", v)
				return v
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

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
func generateStringsSlice(size int) []string {
	slice := make([]string, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = RandomString(10)
	}
	return slice
}

func main() {
	found := int64(0)
	notfound := int64(0)
	for x := 0; x < 1; x++ {
		slice := generateSlice(500)
		fmt.Println("\n--- Unsorted --- \n\n", slice)
		qsort.QuickSortInts(slice)
		fmt.Println("\n--- Sorted ---\n\n", slice)
		inslice := pickValueInSlice(slice)
		outslice := pickValueNotInSlice(slice)
		fmt.Println("Looking for (in slice)", inslice)
		t1 := time.Now()
		success, index := binchop.BinChopInts(inslice, slice)
		taken := time.Since(t1)
		found += taken.Nanoseconds()
		fmt.Printf("Time taken: %v\n", taken)
		if !success {
			fmt.Println(inslice, "not found in slice \n", slice)
			return
		}
		fmt.Println(success, slice[index], "found")

		fmt.Println("Looking for (not in slice)", outslice)
		t2 := time.Now()
		success, _ = binchop.BinChopInts(outslice, slice)
		taken2 := time.Since(t2)
		notfound += taken2.Nanoseconds()
		fmt.Printf("Time taken: %8v\n", taken)
		if success {
			fmt.Println("Error", outslice, "found in slice \n", slice)
			return
		}
		fmt.Println(success, outslice, "not found")
		sslice := generateStringsSlice(500)
		//ss := stringSorter{sslice}
		fmt.Println("\n--- Unsorted --- \n\n", sslice)
		//qsort.Quicksort(&ss)
		qsort.QuickSortStrings(sslice)
		fmt.Println("\n--- Sorted ---\n\n", sslice)
		lookfor := pickStringValueInSlice(sslice)
		sc := stringSliceComp{sslice}
		success, index = binchop.BinChop(lookfor, 0, len(sslice), sc)
		if success {
			fmt.Println(success, sslice[index], "found")
		}

	}
	fmt.Println("Found", found)
	fmt.Println("Not Found", notfound)

}
