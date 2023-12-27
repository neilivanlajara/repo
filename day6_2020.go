//part 1 

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// countPeople := 0
	countMap := 0
	map1 := make(map[rune]bool)

	for scanner.Scan() {
		if scanner.Text() == "" {
			countMap += len(map1)
			map1 = make(map[rune]bool)
		} else {
			// countPeople++
			for _, r := range scanner.Text() {
				map1[r] = true
			}
		}
	}

	fmt.Println(countMap)
}



//part 2 
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	countPeople := 0
	countMap := 0
	map1 := make(map[rune]int)

	for scanner.Scan() {
		if scanner.Text() == "" {
			for _, r := range map1 {
				if r == countPeople {
					countMap++
				}
			}
			countPeople = 0
			map1 = make(map[rune]int)
		} else {
			countPeople++
			for _, r := range scanner.Text() {
				map1[r]++
			}
		}
	}

	fmt.Println(countMap)
}
  
