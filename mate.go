package main

import (
	"fmt"
	"strconv"
)

func fatt(n int) int {
	if n == 1 {
		return 1
	}
	return n * fatt(n-1)
}
func confronto(s string, n int) bool {
	m := make(map[byte]int)
	//numero massimo

	for i := 0; i < len(s); i++ {
		//trasformo in numero per face check
		j, _ := strconv.Atoi(string(s[i]))
		if j > n || j == 0 {
			return false
		}
		if m[s[i]] == 0 {
			m[s[i]]++
		} else {
			return false
		}
	}
	return true
}
func main() {
	//dimmi quanti elementi
	n := 6
	z := 6
	count := 0
	//fmt.Println(confronto("1234"))
	//creo la strina n(n-1)
	//fmt.Println(confronto("13254"))
	s := ""
	for n != 0 {
		temp := strconv.Itoa(n)
		s = s + temp
		n--
	}
	fmt.Println(s)
	// converto s in numero
	numero, _ := strconv.Atoi(s)
	//fmt.Println(numero)
	for i := 0; i <= numero; i++ {
		h := strconv.Itoa(i)
		//fmt.Println(h)
		if len(h) == z {
			if confronto(h, z) {

				fmt.Println(i, confronto(h, z))
				count++
			}
		}
	}
	fmt.Println("elementi totali: ", count)
}
