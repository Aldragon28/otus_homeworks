package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	punctuations  = "!\"#$%&'()*+,-./:;?@[\\]^_`{|}~"
	regexpCompile = regexp.MustCompile(fmt.Sprintf(`^[%[1]s]+|[%[1]s]+$`, punctuations))
	topSize       = 10
)

func prepareWord(word string) string {
	lowerCase := strings.ToLower(word)
	return string(regexpCompile.ReplaceAll([]byte(lowerCase), []byte("")))
}

func Top10(text string) []string {
	words := strings.Fields(text)

	var uniqueWords []string
	wordFrequency := make(map[string]int)
	for _, word := range words {
		newWord := prepareWord(word)
		if newWord != "" {
			if wordFrequency[newWord] == 0 {
				uniqueWords = append(uniqueWords, newWord)
			}

			wordFrequency[newWord]++
		}
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		iFrequency := wordFrequency[uniqueWords[i]]
		jFrequency := wordFrequency[uniqueWords[j]]

		if iFrequency > jFrequency {
			return true
		}

		if iFrequency == jFrequency {
			return strings.Compare(uniqueWords[i], uniqueWords[j]) < 0
		}

		return false
	})

	if len(uniqueWords) > topSize {
		return uniqueWords[:topSize]
	}

	return uniqueWords
}
