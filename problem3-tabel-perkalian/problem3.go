package main

import (
	"fmt"
)

func CetakTabelPerkalian(number int) string {
	// your code here
	var huruf string
	var k int
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			k = i * j
			huruf = huruf + fmt.Sprint(k, "\t")
		}
		//fmt.Println()
		huruf = huruf + "\n"
	}
	return huruf
}

func main() {
	fmt.Println(CetakTabelPerkalian(9))
}
