package goreloaded

import (
	"strconv"
	"strings"
)

func ToUpper(s []string) []string {
	for i, word := range s {
		if word == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if word == "(up," {
			number := strings.Trim(string(s[i+1]), s[i+1][1:])
			num, _ := strconv.Atoi(string(number))
			for j := 1; j <= num; j++ {
				s[i-j] = strings.ToUpper((s[i-j]))
			}
			s = append(s[:i], s[i+2:]...)
		}
	}

	return s
}
