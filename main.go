
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pass("1abc2 pqr3stu8vwx a1b2c3d4e5f treb7uchet"))
}

func pass(str string) int {
	var sum int
	for _, s := range strings.Split(str, " ") {
		var first, last int
		for _, r := range s {
			if r >= '0' && r <= '9' { //line checks if the character r is a digit (i.e., between '0' and '9' inclusive).
				if first == 0 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}
		if first != 0 && last != 0 {
			sum += first*10 + last
		}
	}
	return sum
}