package didyoumean

import (
	"strings"
)

var (
	// ThresholdRate is the rate that allows the edit distanse less than, eg 0.4
	//  means the edit distance less than 40%
	ThresholdRate float64
	// CaseInsensitive compare the edit distance in case insensitive mode
	CaseInsensitive bool
)

// minimum returns minimum value
func minimum(values ...int) (min int) {
	min = values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return
}

// findEditDistance returns edit distance between a and b
// it's use Levenshtein distance, see: https://en.wikipedia.org/wiki/Levenshtein_distance for
// more information
func findEditDistance(a, b string) (distance int) {
	lenA := len(a)
	lenB := len(b)
	totalLen := lenB + 1
	// only use two rows
	v0 := make([]int, totalLen)
	v1 := make([]int, totalLen)
	// initialize first row
	for i := 0; i < totalLen; i++ {
		v0[i] = i
	}
	// loop through the text
	for i := 0; i < lenA; i++ {
		v1[0] = i + 1
		for j := 0; j < lenB; j++ {
			dc := v0[j+1] + 1
			ic := v1[j] + 1
			sc := v0[j]
			if a[i] != b[j] {
				sc++
			}
			min := minimum(dc, ic, sc)
			v1[j+1] = min
		}
		// copy v1 to v0
		for i, v := range v1 {
			v0[i] = v
		}
	}
	return v0[lenB]
}

// FirstMatch returns first match of didyoumean
func FirstMatch(key string, list []string) (result string) {
	if len(key) == 0 {
		return
	}
	if CaseInsensitive {
		key = strings.ToLower(key)
	}
	var winner int
	if ThresholdRate > 0 {
		winner = int(ThresholdRate * float64(len(key)))
	} else {
		winner = -1
	}
	for _, str := range list {
		distance := findEditDistance(key, str)
		if winner < 0 || distance <= winner {
			// winner = distance
			result = str
			return
		}
	}
	return
}

// Match returns all match of didyoumean
func Match(key string, list []string) (results []string) {
	if len(key) == 0 {
		return
	}
	if CaseInsensitive {
		key = strings.ToLower(key)
	}
	var winner int
	if ThresholdRate > 0 {
		winner = int(ThresholdRate * float64(len(key)))
	} else {
		winner = -1
	}
	for _, result := range list {
		distance := findEditDistance(key, result)
		if winner < 0 || distance <= winner {
			winner = distance
			results = append(results, result)
		}
	}
	return
}
