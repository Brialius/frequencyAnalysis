package fAnalysis

import (
	"sort"
	"strings"
	"unicode"
)

type KeyValues struct {
	Key   string
	Value int
}

// Get a string with text, return slice of strings with top 10 frequent words in a text
func fAnalysis(s string) []string {
	res := make([]string, 0, 10)
	freqTable := map[string]int{}

	// Normalize text: delete all except valid unicode letters
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		} else {
			return ' '
		}
	}, s)

	for _, w := range strings.Fields(s) {
		freqTable[w] += 1
	}

	kv := make([]KeyValues, 0, len(freqTable))
	for k, v := range freqTable {
		kv = append(kv, KeyValues{k, v})
	}

	sort.Slice(kv, func(i, j int) bool {
		return kv[i].Value > kv[j].Value
	})

	for i, ws := range kv {
		if i < 10 {
			res = append(res, ws.Key)
		} else {
			break
		}
	}
	return res
}
