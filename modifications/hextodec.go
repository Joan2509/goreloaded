package goreloaded

import (
	"strconv"
)

func HextoInt(s []string) []string {
	for i, word := range s {
		if word == "(hex)" {
			decimal, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				// If conversion fails, return the original string
				return s
			}
			// Replace the previous word with its decimal equivalent
			s[i-1] = strconv.FormatInt(decimal, 10)
			// Remove "(hex)" from the current word
			s = append(s[:i], s[i+1:]...)
			// Break the loop after the first replacement
			break
		}
	}
	return s
}
