package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//1 - 2 -> 33
	//1 - 1 -> 78
	//1 - 3 -> 247
	//1 - 5 -> 68
	//1 - 7 -> 69
	scanner := bufio.NewScanner(os.Stdin)
	trees := 0
	right := 0

	for scanner.Scan() {
		text := scanner.Text()

		if right > len(text)-1 {
			right %= len(text)
		}

		if text[right] == '#' {
			trees++
		}

		right += 7

	}

	fmt.Println(trees)

}
