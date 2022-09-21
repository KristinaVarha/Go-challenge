package piscine

import "github.com/01-edu/z01"

func ContainsPlusMinus(s string) bool {
	for _, c := range s {
		if c == '+' || c == '-' {
			return true
		}
	}
	return false
}

func UniqueChar(s string) bool {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n-1; i++ {
		if InStr(r[i], s[i+1:n]) {
			return false
		}
	}
	return true
}

func InStr(c rune, s string) bool {
	for _, r := range s {
		if c == r {
			return true
		}
	}
	return false
}

func ValidBase(base string) bool {
	return len(base) >= 2 && !ContainsPlusMinus(base) && UniqueChar(base)
}

func PrintNbrBase(n int, base string) {
	if ValidBase(base) {
		length := len(base)
		sign := 1
		rbase := []rune(base)
		if n < 0 {
			z01.PrintRune('-')
			sign = -1
		}
		if n < length && n >= 0 {
			z01.PrintRune(rbase[n])
		} else {
			PrintNbrBase(sign*(n/length), base)
			z01.PrintRune(rbase[sign*(n%length)])
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
