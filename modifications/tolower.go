package goreloaded

import (
	"strconv"
	"strings"
)

func ToLower(s []string) []string {
	for i, word := range s {
		if word == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if word == "(low," {
			number := strings.Trim(string(s[i+1]), s[i+1][1:])
			num, _ := strconv.Atoi(string(number))
			for j := 1; j <= num; j++ {
				s[i-j] = strings.ToLower((s[i-j]))
			}
			s = append(s[:i], s[i+2:]...)
		}
	}

	return s
}
