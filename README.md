# didyoumean
![Go](https://github.com/sc0Vu/didyoumean/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sc0Vu/didyoumean)](https://goreportcard.com/report/github.com/sc0Vu/didyoumean)

Didyoumean written in golang, find the similar string from the given string list. Currently it use Levenshtein distance to calculate edit distannce between two strings. See: https://en.wikipedia.org/wiki/Levenshtein_distance

# Installation

```
$ go get https://github.com/sc0vu/didyoumean
```

or you can

```GO
import (
    "https://github.com/sc0vu/didyoumean"
)
```

# Usage

## Parameters

### ThresholdRate float64
ThresholdRate is the rate that allows the edit distanse less than, eg 0.4 means the edit distance less than 40%

### CaseInsensitive bool
CaseInsensitive compare the edit distance in case insensitive mode

## FirstMatch
Find first match of the given string list
```GO
didyoumean.FirstMatch("key", []string{"kkk", "apple", "kea"})
```

## Match
Find all match of the given string list
```GO
didyoumean.Match("key", []string{"kkk", "apple", "kea"})
```

# License
MIT
