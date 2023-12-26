package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		temp := 0
		text := scanner.Text()

		min, _ := strconv.Atoi(text[:strings.Index(text, "-")])
		max, _ := strconv.Atoi(text[strings.Index(text, "-")+1 : strings.Index(text, " ")])
		letter := text[strings.Index(text, ":")-1]
		pass := strings.Replace(text[strings.Index(text, ":")+1:], " ", "", -1)
		// fmt.Println(letter, " letter")
		// fmt.Println(pass)
		// fmt.Println(min, max)

		// if strings.Count(pass, string(letter)) >= min && strings.Count(pass, string(letter)) <= max {
		// 	count++
		// }

		if pass[min-1] == letter && pass[max-1] != letter {
			fmt.Println(pass)

			temp++
		}

		if pass[min-1] != letter && pass[max-1] == letter {
			temp++
		}

		if temp == 1 {
			count++
		}

	}

	fmt.Println(count)

}
