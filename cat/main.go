package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lens := 0
	argum := os.Args
	for range argum {
		lens++
	}
	for i := 1; i < lens; i++ {
		file, err := ioutil.ReadFile(os.Args[i])
		strFile := string(file)
		if err != nil {
			errmessage := "open " + os.Args[i] + ": no such file or directory"
			for _, v := range errmessage {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')

		} else {
			for _, v := range strFile {
				z01.PrintRune(v)
			}
		}
	}
	if lens == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			str := err.Error()
			for _, v := range str {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
			return
		}
	}
}
