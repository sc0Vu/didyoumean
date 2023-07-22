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
	Key             string
	List            []string
	Match           string
	ThreadRate      float64
	CaseInsensitive bool
}

type MatchTest struct {
	Key             string
	List            []string
	Match           []string
	ThreadRate      float64
	CaseInsensitive bool
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
		}, {
			A:        "台灣 Taiwan",
			B:        "台味",
			Distance: 7,
		},
	}
	firstMatchTests = []FirstMatchTest{
		{
			Key:        "insargrm",
			List:       []string{"facebook", "twitter", "instagram", "linkedin", "台灣 Taiwan"},
			Match:      "instagram",
			ThreadRate: 0.4,
		}, {
			Key:        "insargrm",
			List:       []string{"facebook", "twitter", "instagram", "linkedin", "台灣 Taiwan"},
			Match:      "",
			ThreadRate: 0.3,
		},
		{
			Key:             "insarGrm",
			List:            []string{"facebook", "twiTter", "InstaGram", "linkedin", "台灣 Taiwan"},
			Match:           "",
			ThreadRate:      0.4,
			CaseInsensitive: false,
		}, {
			Key:             "insarGrm",
			List:            []string{"facebook", "twitter", "InstaGram", "linkedin", "台灣 Taiwan"},
			Match:           "InstaGram",
			ThreadRate:      0.5,
			CaseInsensitive: false,
		},
	}
	matchTests = []MatchTest{
		{
			Key:        "insargrm",
			List:       []string{"facebook", "twitter", "instagram", "linkedin", "台灣 Taiwan"},
			Match:      []string{"instagram"},
			ThreadRate: 0.4,
		}, {
			Key:        "insargrm",
			List:       []string{"facebook", "twitter", "instagram", "linkedin", "台灣 Taiwan"},
			Match:      []string{},
			ThreadRate: 0.3,
		}, {
			Key:             "insarGrm",
			List:            []string{"facebook", "twiTter", "InstaGram", "linkedin", "台灣 Taiwan"},
			Match:           []string{},
			ThreadRate:      0.4,
			CaseInsensitive: false,
		}, {
			Key:             "insarGrm",
			List:            []string{"facebook", "twitter", "InstaGram", "linkedin", "台灣 Taiwan"},
			Match:           []string{"InstaGram"},
			ThreadRate:      0.5,
			CaseInsensitive: false,
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

// FuzzFindEditDistance
func FuzzFindEditDistance(f *testing.F) {
	for _, test := range editDistanceTests {
		f.Add(test.A, test.B)
	}
	f.Fuzz(func(t *testing.T, a string, b string) {
		d := findEditDistance(a, b)
		if d < 0 {
			t.Errorf("distance should not be smaller than 0: %d", d)
		}
	})
}

// TestFirstMatch
func TestFirstMatch(t *testing.T) {
	for _, test := range firstMatchTests {
		ThresholdRate = test.ThreadRate
		CaseInsensitive = test.CaseInsensitive
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
