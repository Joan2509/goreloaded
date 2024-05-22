package goreloaded

import (
	"strings"
	"unicode"
)

package main

import (
	"strings"
	"unicode"
)

func Punctuations(lines []string) []string {
	correctedLines := make([]string, len(lines))
	for i, line := range lines {
		var result strings.Builder
		runes := []rune(line)
		n := len(runes)
		for j := 0; j < n; j++ {
			if isPunctuation(runes[j]) {
				// Write the punctuation mark
				if j > 0 && !unicode.IsSpace(runes[j-1]) {
					result.WriteRune(runes[j])
				} else {
					if result.Len() > 0 && result.String()[result.Len()-1] != ' ' {
						result.WriteRune(' ')
					}
					result.WriteRune(runes[j])
				}
				// Check for grouped punctuations
				for j+1 < n && isPunctuation(runes[j+1]) {
					result.WriteRune(runes[j+1])
					j++
				}
				// Add space after punctuation if the next character is not space and not punctuation
				if j+1 < n && !unicode.IsSpace(runes[j+1]) && !isPunctuation(runes[j+1]) {
					result.WriteRune(' ')
				}
			} else {
				// Write non-punctuation character
				result.WriteRune(runes[j])
			}
		}
		correctedLines[i] = strings.TrimSpace(result.String())
	}
	return correctedLines
}

func isPunctuation(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';'
}