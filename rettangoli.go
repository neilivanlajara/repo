package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rettangolo struct {
	nome       string
	alt, largh float64
}

func getArea(r Rettangolo) float64 {
	return r.alt * r.largh
}

func maggiore(r1, r2 Rettangolo) (b bool) {

	if getArea(r1) > getArea(r2) {
		return true
	}
	return b
}
func newparseData(s string) (z string, p, o float64) {
	barra1 := strings.IndexRune(s, ' ')
	//var err error
	var x, y int
	barra2 := strings.IndexRune(s[barra1+1:], ' ')

	//fmt.Println(s[barra1+barra2+1+1:])
	z = s[:barra1]

	x, _ = strconv.Atoi(s[barra1+1 : barra1+barra2+1])

	y, _ = strconv.Atoi(s[barra1+1+barra2+1:])

	return z, float64(x), float64(y)
}

func allarga(r1 *Rettangolo, zoom float64) {

	*&r1.alt = r1.alt * zoom
	*&r1.largh = r1.largh * zoom
}

func main() {
	var all []Rettangolo
	all = make([]Rettangolo, 0)

	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("dammi i valori")
	for scanner.Scan() {
		var r Rettangolo

		temp := scanner.Text()
		if temp == "" {
			fmt.Println("test")
			break
		}
		//fmt.Println("questo è temp", temp)
		if strings.Count(temp, " ") == 3 {
			var zoom float64
			var ultimospazio int
			for i, valore := range temp {
				if valore == ' ' {
					ultimospazio = i
				}
			}
			//fmt.Println(s[:ultimospazio])
			r.nome, r.alt, r.largh = newparseData(temp[:ultimospazio])
			//fmt.Println(r)
			//fmt.Println(s[ultimospazio+1:])
			zoomint, _ := strconv.Atoi(temp[ultimospazio+1:])
			//fmt.Println(zoomint)
			zoom = float64(zoomint)
			//fmt.Println("floatzoom: ", zoom)
			allarga(&r, zoom)

		} else {
			r.nome, r.alt, r.largh = newparseData(temp)
		}
		all = append(all, r)

	}
	var max Rettangolo

	for _, r := range all {
		if maggiore(r, max) {
			max = r
		}
	}
	fmt.Println("max", max)
	//fmt.Println(all)
	/*var r Rettangolo
	s := "tre 8 67 2"
	var zoom float64
	var ultimospazio int
	for i, valore := range s {
		if valore == ' ' {
			ultimospazio = i
		}
	}
	//fmt.Println(s[:ultimospazio])
	r.nome, r.alt, r.largh = newparseData(s[:ultimospazio])
	fmt.Println(r)
	//fmt.Println(s[ultimospazio+1:])
	zoomint, _ := strconv.Atoi(s[ultimospazio+1:])
	fmt.Println(zoomint)
	zoom = float64(zoomint)
	fmt.Println("floatzoom: ", zoom)
	allarga(&r, zoom)
	fmt.Println(r)*/
}
