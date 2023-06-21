package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regx = regexp.MustCompile(`[^\p{L}\p{N} -]+`)

func Top10(input string) []string {
	arr := strings.Fields(input)
	result := make([]string, 0)

	if len(arr) == 0 {
		return nil
	}

	type wordCounter struct {
		Word    string
		Counter int
	}
	sliceWordCounter := make([]wordCounter, 0)
	mapWordCounter := make(map[string]int)

	for _, word := range arr {
		cleanWord := regx.ReplaceAllString(word, "")
		lowWord := strings.ToLower(cleanWord)
		if lowWord == "-" {
			continue
		}
		mapWordCounter[lowWord]++
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
