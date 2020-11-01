package levenshtein

// First parameter is so that it is a compile error to call with no parameters
// minimum returns the minimum of (1..n) of int parameters
func minimum(a int, v ...int) int {
	min := a
	var n int
	for _, n = range v {
		if n < min {
			min = n
		}
	}
	return min
}

// LevenshteinDistance returns the edit difference between to strings
func LevenshteinDistance(s string, t string) int {
	slen := len(s)
	tlen := len(t)

	/*  degenerate cases */
	if s == t {
		return 0
	}
	if slen == 0 {
		return tlen
	}
	if tlen == 0 {
		return slen
	}

	sr := []rune(s)
	tr := []rune(t)
	/*  create two work vectors of integer distances */

	v0 := make([]int, tlen+1)
	v1 := make([]int, tlen+1)

	/*  initialize v0 (the previous row of distances) */
	/*  this row is A[0][i]: edit distance for an empty s */
	/*  the distance is just the number of characters to delete from t */
	for n := range v0 {
		v0[n] = n
	}

	for i, y := range sr {
		/*  calculate v1 (current row distances) from the previous row v0 */

		/*  first element of v1 is A[i+1][0] */
		/*    edit distance is delete (i+1) chars from s to match empty t */
		v1[0] = i + 1

		/*  use formula to fill in the rest of the row */

		for j, x := range tr {
			// calculating costs for A[i+1][j+1]
			substitutionCost := 0
			deletionCost := v0[j+1] + 1
			insertionCost := v1[j] + 1
			if x == y {
				substitutionCost = v0[j]
			} else {
				substitutionCost = v0[j] + 1
			}
			v1[j+1] = minimum(deletionCost, insertionCost, substitutionCost)

		}
		// copy v1 (current row) to v0 (previous row) for next iteration
		// since data in v1 is always invalidated, a swap without copy could be more efficient
		v0, v1 = v1, v0

	}
	return v0[tlen]
}
