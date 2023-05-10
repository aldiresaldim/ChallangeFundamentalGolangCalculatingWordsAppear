package main

import "fmt"

func main() {
	kalimat := "selamat malam"

	for _, huruf := range kalimat {
		fmt.Println(string(huruf))
	}

	hurufMap := make(map[rune]int)
	for _, huruf := range kalimat {
		hurufMap[huruf]++
	}

	fmt.Println(hurufMap)
}
