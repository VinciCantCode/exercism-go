package parallelletterfrequency

import "unicode"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	m := FreqMap{}
	for _, t := range text {
		if unicode.IsLetter(t) {
			t = unicode.ToLower(t)
			m[t]++
		}
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	result := FreqMap{}
	ch := make(chan FreqMap)
	for _, text := range texts {
		go func(s string) {
			ch <- Frequency(s)
		}(text)
	}

	for range texts {
		res := <-ch
		for char, count := range res {
			result[char] += count
		}
	}
	return result
}
