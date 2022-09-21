package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	isUpper := false
	for i, a := range arguments {
		if a == "--upper" {
			isUpper = true
		}
		if i != 0 && !IsNumeric(a) {
			z01.PrintRune(' ')
			continue
		}
		if i != 0 && (Atoi(a) > 26 || Atoi(a) < 1) {
			z01.PrintRune(' ')
			continue
		}
		if i != 0 && isUpper {
			tmp0 := Atoi(a) + 64
			z01.PrintRune(rune(tmp0))
			continue
		}
		if i != 0 && !isUpper {
			tmp1 := Atoi(a) + 96
			z01.PrintRune(rune(tmp1))
			continue
		}
	}
	z01.PrintRune('\n')
}

func IsNumeric(str string) bool {
	for _, a := range str {
		if !(a >= '0' && a <= '9') {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	newStr := []rune(s)
	count := StrLen(s)
	sstr := ""
	if s == "+" || s == "-" || s == "" {
		return 0
	}
	for i := 1; i < count; i++ {
		sstr = sstr + string(newStr[i])
	}
	if newStr[0] == '-' {
		return -1 * BasicAtoi2(sstr)
	}
	if newStr[0] == '+' {
		return BasicAtoi2(sstr)
	}
	return BasicAtoi2(s)
}

func BasicAtoi2(s string) int {
	count := StrLen(s)
	newStr := []rune(s)
	val := 0
	for i := 0; i < count; i++ {
		t := 1
		for j := 0; j < count-i-1; j++ {
			t = t * 10
		}
		sym := 0
		for k := '0'; k < newStr[i]; k++ {
			sym++
		}
		if sym == 0 && newStr[i] != '0' {
			return 0
		}
		if sym > 9 {
			return 0
		}
		val = val + sym*t
	}
	return val
}

func StrLen(str string) int {
	newStr := []rune(str)
	count := 0
	for range newStr {
		count++
	}
	return count
}
