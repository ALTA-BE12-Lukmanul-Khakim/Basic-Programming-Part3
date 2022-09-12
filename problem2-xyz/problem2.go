package main

import "fmt"

func DrawXYZ(n int) string {
	var k int
	var hasil = ""
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			k++
			if k%3 == 0 {
				hasil = hasil + fmt.Sprintf("X ")
			} else if k%2 == 1 {
				hasil = hasil + fmt.Sprintf("Y ")
			} else if k%2 == 0 {
				hasil = hasil + fmt.Sprintf("Z ")
			}
		}
		hasil = hasil + "\n"
	}
	return hasil
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
