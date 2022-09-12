package main

import "fmt"

func PlayWithAsterix(n int) string {
	var tanda string
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			tanda += (" ")
		}
		for k := 1; k <= i; k++ {
			tanda += ("* ")
		}
		tanda = tanda + "\n"
	}
	return tanda
}

func main() {
	fmt.Println(PlayWithAsterix(6))

}
