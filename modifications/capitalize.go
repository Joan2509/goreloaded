package goreloaded

import (
	"strconv"
	"strings"
)

func Capitalize(s []string) []string {
	for i, word := range s {
		if word == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if word == "(cap," {
			number := strings.Trim(string(s[i+1]), s[i+1][1:])
			num, _ := strconv.Atoi(string(number))
			for j := 1; j <= num; j++ {
				s[i-j] = strings.Title((s[i-j]))
			}
			s = append(s[:i], s[i+2:]...)
		}
	}

	return s
}
