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
	res := ""
	passportsRaw := make([]string, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			res += scanner.Text() + " "
		} else {

			if strings.Count(res, ":") == 8 || strings.Count(res, ":") == 7 && !strings.Contains(res, "cid") {
				passportsRaw = append(passportsRaw, res)
			}

			res = ""
		}
	}

	fmt.Println(passportsRaw)

	modify := func(r rune) rune {
		if r == ' ' {
			return 0
		}
		return r
	}
	for _, r := range passportsRaw {
		count := 0

		temp := strings.Split(r, " ")

		fmt.Println(len(temp))
		for _, atb := range temp {
			q := strings.Map(modify, atb)

			if atb != "" {
				info := q[:3]
				if info == "cid" {
					continue
				} else if info == "byr" {

					year, _ := strconv.Atoi(q[4:])
					if year >= 1920 && year <= 2002 {
						count++
					}
				} else if info == "iyr" {

					year, _ := strconv.Atoi(q[4:])
					if year >= 2010 && year <= 2020 {
						count++
					}
				} else if info == "eyr" {

					year, _ := strconv.Atoi(q[4:])
					if year >= 2020 && year <= 2030 {
						count++
					}
				} else if info == "hgt" {
					height, _ := strconv.Atoi(q[4 : len(q)-2])
					if q[len(q)-1] == 'm' {
						if height >= 150 && height <= 193 {
							count++
						}
					} else {
						if height >= 59 && height <= 76 {
							count++
						}
					}
				} else if info == "hcl" {
					if q[4] == '#' {
						tempCount := 0
						for _, r := range q[5:] {
							if r >= '0' && r <= 'f' {
								tempCount++
							}
						}
						if tempCount == len(q[5:]) {
							count++
						}
					}
				} else if info == "ecl" {
					color := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

					for _, r := range color {
						if r == q[4:] {
							count++
						}
					}
				} else if info == "pid" {

					tempCount := 0
					for _, r := range q[4:] {
						if r >= '0' && r <= '9' {
							tempCount++
						}
					}
					if tempCount == 9 {
						count++
					}

					// fmt.Println(q[4:], count)
					// fmt.Println(q[:3])
				}
			}
			// byr:1962 pid:547578491 eyr:2028 ecl:hzl hgt:63in iyr:2013 hcl:#ap3a2f
		}
		// fmt.Println(count, tot)
		if count == 7 {
			tot++
		}

	}

	fmt.Println(tot, " passaporti validi")

}
