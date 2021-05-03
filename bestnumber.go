package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func checkNumberRaw(s string) bool {
	checker := true
	for _, r := range s {
		if !unicode.IsDigit(r) {
			fmt.Fprintln(os.Stderr, string(r), " non è un numero ")
			checker = false
		}
	}
	return checker
}
func checkLenNumberRaw(s string) bool {
	checker := true

	if len(s) < 3 {
		fmt.Fprintln(os.Stderr, "NaN")

		checker = false
	}
	return checker
}

// case 1:= 45678
func main() {
	numeroRaw := os.Args[1]
	bestnumber := 0
	if !checkNumberRaw(numeroRaw) {
		return
	} else if !checkLenNumberRaw(numeroRaw) {
		return
	}

	window := 3
	for i := 0; i < len(numeroRaw); i++ {
		//sinistra := string(numeroRaw[:i])
		//destra := string(numeroRaw[window+i:])
		temp, _ := strconv.Atoi(string(numeroRaw[:i]) + string(numeroRaw[window+i:]))
		//fmt.Println(temp)
		if temp > bestnumber {
			bestnumber = temp
		}
		if window+i == len(numeroRaw) {
			break
		}
	}

	fmt.Println(bestnumber)
}
