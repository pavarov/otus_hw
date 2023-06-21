package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCounter struct {
	Word    string
	Counter int
}

var sliceWordCounter []wordCounter

func Top10(input string) []string {
	arr := strings.Fields(input)
	var result []string

	if len(arr) == 0 {
		return result
	}
	var mapWordCounter = make(map[string]int)

	for _, word := range arr {
		if _, ok := mapWordCounter[word]; ok {
			mapWordCounter[word]++
		} else {
			mapWordCounter[word] = 1
		}
	}

	for word, counter := range mapWordCounter {
		sliceWordCounter = append(sliceWordCounter, wordCounter{word, counter})
	}

	sort.Slice(sliceWordCounter, func(idx, idxN int) bool {
		if sliceWordCounter[idx].Counter == sliceWordCounter[idxN].Counter {
			return sliceWordCounter[idx].Word < sliceWordCounter[idxN].Word
		}
		return sliceWordCounter[idx].Counter > sliceWordCounter[idxN].Counter
	})

	for _, wc := range sliceWordCounter {
		result = append(result, wc.Word)
	}

	return result[0:10]
}
