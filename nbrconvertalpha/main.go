package main

import (
	"os"

	"github.com/01-edu/z01"
)

func tonumber(str string) int {
	check := []rune(str)
	ans := 0
	i := 0
	for _, s := range check {
		if s >= 48 && s <= 57 {
			ans = ans*10 + (int(s) - 48)
		} else {
			return 32
		}
		i++
		if i == 3 {
			return 32
		}
	}

	if ans >= 1 && ans <= 26 {
		return ans + 96
	} else {
		return 32
	}
}

func main() {
	len := 0
	for range os.Args {
		len++
	}
	if len > 0 {
		a := 0
		l := 1
		if len > 1 && os.Args[1] == "--upper" {
			z01.PrintRune(32)
			a = -32
			l = 2
		}
		for i := l; i < len; i++ {
			num := tonumber(os.Args[i])
			z01.PrintRune(rune(num + a))
		}
	}
	z01.PrintRune(10)
}
