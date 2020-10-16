package didyoumean

import (
	"testing"
)

type EditDistanceTest struct {
	A        string
	B        string
	Distance int
}

type FirstMatchTest struct {
	Key           string
	List          []string
	Match         string
	ThreadRate    float64
	CaseSensitive bool
}

type MatchTest struct {
	Key           string
	List          []string
	Match         []string
	ThreadRate    float64
	CaseSensitive bool
}

var (
	editDistanceTests = []EditDistanceTest{
		{
			A:        "kitten",
			B:        "sitting",
			Distance: 3,
		}, {
			A:        "Saturday",
			B:        "Sunday",
			Distance: 3,
		},
	}
	firstMatchTests = []FirstMatchTest{
		{
			Key:           "insargrm",
			List:          []string{"facebook", "twitter", "instagram", "linkedin"},
			Match:         "instagram",
			ThreadRate:    0.4,
			CaseSensitive: false,
		}, {
			Key:           "insargrm",
			List:          []string{"facebook", "twitter", "instagram", "linkedin"},
			Match:         "",
			ThreadRate:    0.3,
			CaseSensitive: false,
		},
		{
			Key:           "insarGrm",
			List:          []string{"facebook", "twiTter", "InstaGram", "linkedin"},
			Match:         "",
			ThreadRate:    0.4,
			CaseSensitive: true,
		}, {
			Key:           "insarGrm",
			List:          []string{"facebook", "twitter", "InstaGram", "linkedin"},
			Match:         "InstaGram",
			ThreadRate:    0.5,
			CaseSensitive: true,
		},
	}
	matchTests = []MatchTest{
		{
			Key:           "insargrm",
			List:          []string{"facebook", "twitter", "instagram", "linkedin"},
			Match:         []string{"instagram"},
			ThreadRate:    0.4,
			CaseSensitive: false,
		}, {
			Key:           "insargrm",
			List:          []string{"facebook", "twitter", "instagram", "linkedin"},
			Match:         []string{},
			ThreadRate:    0.3,
			CaseSensitive: false,
		}, {
			Key:           "insarGrm",
			List:          []string{"facebook", "twiTter", "InstaGram", "linkedin"},
			Match:         []string{},
			ThreadRate:    0.4,
			CaseSensitive: true,
		}, {
			Key:           "insarGrm",
			List:          []string{"facebook", "twitter", "InstaGram", "linkedin"},
			Match:         []string{"InstaGram"},
			ThreadRate:    0.5,
			CaseSensitive: true,
		},
	}
)

// TestFindEditDistance
func TestFindEditDistance(t *testing.T) {
	for _, test := range editDistanceTests {
		d := findEditDistance(test.A, test.B)
		if d != test.Distance {
			t.Fatalf("The distance should be %d but got %d", test.Distance, d)
		}
	}
}

// TestFirstMatch
func TestFirstMatch(t *testing.T) {
	for _, test := range firstMatchTests {
		ThresholdRate = test.ThreadRate
		CaseSensitive = test.CaseSensitive
		m := FirstMatch(test.Key, test.List)
		if m != test.Match {
			t.Fatalf("The match should be %s but got %s", test.Match, m)
		}
	}
}

// TestMatch
func TestMatch(t *testing.T) {
	for _, test := range matchTests {
		ThresholdRate = test.ThreadRate
		m := Match(test.Key, test.List)
		if len(m) != len(test.Match) {
			t.Fatalf("The match should be equal")
		}
		for i, v := range m {
			if m[i] != v {
				t.Fatalf("The match should be %s but got %s", m[i], v)
			}
		}
	}
}
