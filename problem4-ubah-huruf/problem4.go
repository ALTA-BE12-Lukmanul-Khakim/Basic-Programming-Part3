package main

import "fmt"

func UbahHuruf(sentence string) string {
	var res string
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 'A' && sentence[i] <= 'Z' {
			res += string(((sentence[i]-'A')+10)%26 + 'A')
		} else {
			res += (" ")
		}
	}
	return res
}

func main() {
	fmt.Println(UbahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(UbahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(UbahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(UbahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(UbahHuruf("PROGRAMMER"))      // ZBYQBKWWOB
}
