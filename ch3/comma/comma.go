package main

import (
	"bytes"
)

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func comma(s string) string {
	ret := bytes.Buffer{}
	dot := len(s)
	for i, c := range s {
		if c == '.' {
			dot = i
		}
	}
	sign := false
	if s != "" && (s[0] == '-' || s[0] == '+') {
		ret.WriteString(string(s[0]))
		sign = true
	}
	mod := (len(s[:dot]) - btoi(dot != 0) - btoi(sign)) % 3
	if mod == 0 {
		mod = 3
	}
	for i, c := range s {
		if mod == 0 && i < dot {
			ret.WriteString(",")
			mod = 3
		}
		mod--
		ret.WriteString(string(c))

	}
	return ret.String()
}

func isAnagram(s1, s2 string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, c := range s1 {
		m1[c]++
	}
	for _, c := range s2 {
		m2[c]++
	}
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}

func main() {
}
