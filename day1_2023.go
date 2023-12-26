package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tot := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		temp := scanner.Text()
		temp = strings.ReplaceAll(temp, "one", "o1e")
		temp = strings.ReplaceAll(temp, "two", "t2o")
		temp = strings.ReplaceAll(temp, "three", "t3e")
		temp = strings.ReplaceAll(temp, "four", "f4r")
		temp = strings.ReplaceAll(temp, "five", "f5e")
		temp = strings.ReplaceAll(temp, "six", "s6x")
		temp = strings.ReplaceAll(temp, "seven", "s7n")
		temp = strings.ReplaceAll(temp, "eight", "e8t")
		temp = strings.ReplaceAll(temp, "nine", "n9e")

		left, right := 0, 0

		for _, r := range temp {
			if r >= '0' && r <= '9' {
				left, _ = strconv.Atoi(string(r))
				break
			}
		}

		for i := len(temp) - 1; i >= 0; i-- {
			if temp[i] >= '0' && temp[i] <= '9' {
				right, _ = strconv.Atoi(string(temp[i]))
				break
			}
		}

		tempTot := (left * 10) + right

		fmt.Println(tempTot)
		tot += tempTot

	}

	fmt.Println(tot)
}
